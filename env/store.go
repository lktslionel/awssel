package env

// Storer is the API that any Store implementation
// must satisfy
type Storer interface {
	QueryVarsForService(serviceName string, opts ...StoreQueryOptions) ([]*Var, error)
}

// StoreQueryOptions contains additonnal options
// needed to query a store
type StoreQueryOptions struct {
	FilterPattern string
	PrefixPath    string
}
