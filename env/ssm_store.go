package env

import (
	"context"
	"path"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var (
	// AwsselDefaultEnvvarsLimit is the max number of values
	// that it is retrieved from the SSM store
	AwsselDefaultEnvvarsLimit int64 = 100
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
	Region   string
	insecure bool
	Endpoint string
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

		cfg.Endpoint = aws.String(opts.Endpoint)
		cfg.Region = aws.String(opts.Region)
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
		envvars       []*Var
		keyPath       string
		filterPattern string
	)

	// Build the service key path from the given options
	if len(opts) > 0 {
		keyPath = path.Join(opts[0].PrefixPath, name)
		filterPattern = opts[0].FilterPattern
	} else {
		keyPath = name
	}

	// Check wether we got a pattern given as query option
	//isfilterPatternGiven := len(filterPattern) > 0

	params := ssm.GetParametersByPathInput{
		Path:       aws.String(keyPath),
		MaxResults: aws.Int64(AwsselDefaultEnvvarsLimit),
	}

	ctx := context.Background()

	requestPager := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			req, _ := s.conn.GetParametersByPath(&params)
			req.SetContext(ctx)
			return req, nil
		},
	}

	response, err := 

	if err != nil {
		return envvars, err
	}

	for _, param := range response.Parameters {
		envvar := VarFromSSMParameter(param)

		isMatched, _ := regexp.MatchString(filterPattern, envvar.Name)
		if isMatched {
			envvars = append(envvars, envvar)
		}
	}

	return envvars, nil
}
