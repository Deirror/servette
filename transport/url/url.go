package urlx

type Provider interface {
	GetURL() string
	WithQuery(arg, val string) string
}
