package env

// Storer is the API that any Store implementation
// must satisfy
type Storer interface {
	QueryVarsForService(serviceName string, opts StoreQueryOption) ([]*Var, error)
}

// StoreQueryOption contains additonnal options
// needed to query a store
type StoreQueryOption struct {
	FilterPattern string
	PrefixPath    string
}
