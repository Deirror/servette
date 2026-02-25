package doc

import (
	"context"
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoStoreClient struct {
	db *mongo.Database
}

func NewMongoStoreClient(cfg *Config) (*MongoStoreClient, error) {
	clientOpts := options.Client().ApplyURI(cfg.StoreURL)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, err
	}

	return &MongoStoreClient{
		db: client.Database(cfg.Database),
	}, nil
}

func (m *MongoStoreClient) Insert(ctx context.Context, coll string, docs ...interface{}) (interface{}, error) {
	c := m.db.Collection(coll)

	if len(docs) == 1 {
		res, err := c.InsertOne(ctx, docs[0])
		if err != nil {
			return nil, err
		}
		return res.InsertedID, nil
	}

	if len(docs) > 1 {
		res, err := c.InsertMany(ctx, docs)
		if err != nil {
			return nil, err
		}
		return res.InsertedIDs, nil
	}

	return nil, errors.New("no documents provided")
}

func (m *MongoStoreClient) Find(ctx context.Context, coll string, filter interface{}, results interface{}) error {
	c := m.db.Collection(coll)

	val := reflect.ValueOf(results)
	if val.Kind() != reflect.Ptr {
		return errors.New("results argument must be a pointer")
	}
	elem := val.Elem()

	switch elem.Kind() {
	case reflect.Slice:
		cur, err := c.Find(ctx, filter)
		if err != nil {
			return err
		}
		defer cur.Close(ctx)
		return cur.All(ctx, results)

	case reflect.Struct:
		res := c.FindOne(ctx, filter)
		return res.Decode(results)

	default:
		return errors.New("results must be pointer to struct or slice")
	}
}

func (m *MongoStoreClient) Update(ctx context.Context, coll string, filter interface{}, update interface{}) (int64, error) {
	c := m.db.Collection(coll)

	res, err := c.UpdateMany(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

func (m *MongoStoreClient) Delete(ctx context.Context, coll string, filter interface{}) (int64, error) {
	c := m.db.Collection(coll)

	res, err := c.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

func (m *MongoStoreClient) Count(ctx context.Context, coll string, filter interface{}) (int64, error) {
	c := m.db.Collection(coll)
	return c.CountDocuments(ctx, filter, options.Count())
}
