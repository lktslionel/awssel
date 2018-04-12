###### DOCS / AWSSEL / COMMAND
## Load

This command load every environment variables defined for a given service.

### Usage

```bash
awssel load [options]
```

### Options

Name &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;     | Description | Required | Example
---------|----------|---------|----
 `--service-name` | Service name | Yes | -
 `--prefix-path` | Prefix path from which env vars will be retieved<br>All env vars for a given service will be located under the path: `<prefix-path>/<service-name>` | Yes | **prefix-path** = `/os/prod/support/dsi/core`<br><br>**service-name** = `directory-service`
 `--aws-region` | AWS Region Name | Yes | -
 `--filter-pattern` | The pattern that every env var name must match | No | To get all var names prefixed starting with `OS_`,<br>use this filter pattern `--filter-pattern=OS_*`
 `--output-format` | Output env vars to : `default, json, yaml` format | No | 
 `--exportable` | Prefix output string with `export` statement.<br>It only works with `output-format` set `default` | No | Output `export NAME=VALUE`
 `--aws-role` | AWS Role ARN. <br>Note that the role must have a right set of permission to actually request anything for AWS parameter store | No |`os-fake-role`
 `--aws-profile` | AWS Profile Name | No | -
