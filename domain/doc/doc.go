package doc

import "context"

type Storer interface {
	Insert(ctx context.Context, coll string, docs ...interface{}) (interface{}, error)
	Update(ctx context.Context, coll string, filter interface{}, update interface{}) (int64, error)
	Delete(ctx context.Context, coll string, filter interface{}) (int64, error)

	Find(ctx context.Context, coll string, filter interface{}, results interface{}) error
	Count(ctx context.Context, coll string, filter interface{}) (int64, error)
}
