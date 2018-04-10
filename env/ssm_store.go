package env

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// SSMStore is an implementation of a store.
// It is used to query SSM Parameter store
type SSMStore struct {
	conn *ssm.SSM
	sess *session.Session
}

// SSMStoreOptions gather information need to
// create a session with the AWS SSM service
type SSMStoreOptions struct {
	roleARN  string
	region   string
	profile  string
	endpoint *string
}

// NewSSMStore returns a client to query AWS SSM Parameter Store
// It use a SSMStoreOptions as a parameter to pass info needed
// by the store
func NewSSMStore(o ...SSMStoreOptions) *SSMStore {
	var (
		sessOpts session.Options
		opts     SSMStoreOptions
		cfg      *aws.Config
	)

	// Set config attrs
	cfg = aws.NewConfig()
	if len(o) > 0 {
		opts = o[0]

		cfg.Endpoint = opts.endpoint
		cfg.Region = aws.String(opts.region)

		// Create session options filled with value provided
		// by the SSMStoreOptions
		sessOpts = session.Options{
			Config:  *cfg,
			Profile: opts.profile,
		}
	}

	// Create the session Objecft
	sess := session.Must(session.NewSessionWithOptions(sessOpts))

	if opts.roleARN != "" {
		// Retrieves the token from sts for that role
		creds := stscreds.NewCredentials(sess, opts.roleARN)
		cfg = cfg.WithCredentials(creds)
	}

	return &SSMStore{
		sess: sess,
		conn: ssm.New(sess),
	}
}

// QueryVarsForService is used to query SSM Parameter Store
// It returns all env vars related to a given service.
// or and error if something went wrong
//
// See env.StoreQueryOption for more information about available options
func (s *SSMStore) QueryVarsForService(name string, opts ...StoreQueryOption) ([]*Var, error) {

	var envars []*Var
	response, err := s.conn.GetParametersByPath(&ssm.GetParametersByPathInput{})

	if err != nil {
		return envars, err
	}

	for _, param := range response.Parameters {
		envars = append(envars, VarFromSSMParameter(param))
	}

	return envars, nil
}
