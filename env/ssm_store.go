package env

import (
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// SSMStore is an implementation of a store.
// It is used to query SSM Parameter store
type SSMStore struct {
	conn *ssm.SSM
}

// SSMStoreOptions gather information need to
// create a session with the AWS SSM service
type SSMStoreOptions struct {
	roleARN  string
	region   *string
	insecure *bool
	endpoint *string
}

// NewSSMStore returns a client to query AWS SSM Parameter Store
// It use a SSMStoreOptions as a parameter to pass info needed
// by the store
func NewSSMStore(o ...SSMStoreOptions) *SSMStore {
	var (
		opts SSMStoreOptions
		cfg  *aws.Config
	)

	// Set AWS config using given SSM Store options
	cfg = aws.NewConfig()
	if len(o) > 0 {
		opts = o[0]

		cfg.Endpoint = opts.endpoint
		cfg.Region = opts.region
	}

	// Create the session Objecft
	sess := session.Must(session.NewSession())

	return &SSMStore{
		conn: ssm.New(sess, cfg),
	}
}

// QueryVarsForService is used to query SSM Parameter Store
// It returns all env vars related to a given service.
// or and error if something went wrong
//
// See env.StoreQueryOption for more information about available options
func (s *SSMStore) QueryVarsForService(name string, opts ...StoreQueryOptions) ([]*Var, error) {

	var (
		envars  []*Var
		keyPath string
	)

	// Build the service key path from the given options
	if len(opts) > 0 {
		keyPath = path.Join(opts[0].PrefixPath, name)
	} else {
		keyPath = name
	}

	response, err := s.conn.GetParametersByPath(&ssm.GetParametersByPathInput{
		Path: aws.String(keyPath),
	})

	if err != nil {
		return envars, err
	}

	for _, param := range response.Parameters {
		envars = append(envars, VarFromSSMParameter(param))
	}

	return envars, nil
}
