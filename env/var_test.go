package env

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/stretchr/testify/assert"
)

func TestVar(t *testing.T) {

	t.Run("New Envar from SSM Parameter Struct", func(t *testing.T) {
		param := &ssm.Parameter{
			Name:  aws.String("/os/prod/proxy/HOME"),
			Type:  aws.String(ssm.ParameterTypeString),
			Value: aws.String("/home/username"),
		}

		expected := &Var{
			Name:  "HOME",
			Value: "/home/username",
		}

		actual := NewVarFromSSMParameter(param)

		assert.Equal(t, expected, actual)
	})

}
