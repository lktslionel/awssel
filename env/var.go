package env

import (
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Var represent a value object for an Env Variable
type Var struct {
	Name  string
	Value string
}

// NewVarFromSSMParameter create a new env var object from
// an SSM Paramater Struct
func NewVarFromSSMParameter(p *ssm.Parameter) *Var {

	// Remove prefix from parameter name
	key := aws.StringValue(p.Name)
	name := path.Base(key)

	return &Var{
		Name:  name,
		Value: aws.StringValue(p.Value),
	}
}

// String return the string representation of an evn var
func (v *Var) String() string { return "" }

// Export return a string representation of an env var
// according to the given formatter spec
func (v *Var) Export(f Formatter) string { return "" }
