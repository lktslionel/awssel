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

	t.Run("Export Var using DefaultFormatter", func(t *testing.T) {
		var fmtr Formatter
		//{DefaultFormatter(), "HOME=/home/username"},
		//{BashExportFormatter(), "export HOME=/home/username"},

		examples := []struct {
			ev       Var
			expected string
		}{
			{Var{"PROXY_APP_DIR", "/os/app"}, "'PROXY_APP_DIR=/os/app'"},
			{Var{"API_ID", "0ac41e50-3364-11e8-9d7e"}, "'API_ID=0ac41e50-3364-11e8-9d7e'"},
			// Escape some special characters: [:space:], $
			{Var{"PROXY_USERNAME", "os operator"}, "'PROXY_USERNAME=os operator'"},
			{Var{"PROXY_PASS", `784"f43631c05`}, `'PROXY_PASS=784"f43631c05'`},
			{Var{"KEY_STRING", `cURmOUl4U VVX"$ESCAPE"ZHRsdWs=`}, `'KEY_STRING=cURmOUl4U VVX"$ESCAPE"ZHRsdWs='`},
		}

		t.Run("Using DefaultFormatter", func(t *testing.T) {
			fmtr = DefaultFormatter()

			for _, ex := range examples {
				actual := ex.ev.Export(fmtr)
				assert.Equal(t, actual, ex.expected)
			}

		})

		t.Run("Using BashExportFormatter", func(t *testing.T) {
			t.Skip()
		})
	})
}
