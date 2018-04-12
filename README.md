###### OLST / TOOLING
# AWSSEL (AWS Ssm Env Loader)


A tool that load all env vars stored in AWS SSM Parameter Store, for a given service.

See [Our Moviations doc](docs/awssel/motivations.md) for more information about why we build this tool and understand how things work under the hood.



## Contents

* [Install]
* [Usage]
* [Examples]
* [TODO]
* [Tests]
* [Guidelines]
* [Contribute]
* [FAQ]
* [References]
* [Maintainers]
* [License]

<br>

## Install

To install the tool from your shell, run:

```
go get -u github.com/lktslionel/awssel
```

## Usage

```
awssel <command> [options]
```

### Commands

**Awssel** has a only 2 commands : 



Command | Description | Options
---------|----------|---------
 help | Show command help text | -
 load | Load env vars for a service | See [Load options](docs/awssel/cmd/load.md)


## Examples

Get all environment variables for `proxy-web` service, at path `/os` in the region `eu-west-1`.
We have a localstack instance running on `http://localhost:4583`.
If you use this tool on AWS SSM, there is no need to provide the `endpoint` option.

```bash
awssel load --service-name proxy --prefix-path /os --aws-region eu-west-1 --endpoint http://localhost:4583

# Ouput >
# 
# PROXY_USER=web-agent
# PROXY_PASS=784f43631c05
```

### Export

Now if you want to export the output and use it in your script to load those environment variables, add the option `--export` : 

```bash
awssel load --export --service-name proxy --prefix-path /os --aws-region eu-west-1 --endpoint http://localhost:4583 

# Ouput >
# 
# export 'PROXY_USER=web-agent'
# export 'PROXY_PASS=784f43631c05'
```

### Filter

You can also apply filter environment varaible names using a regex.  

Without any filter, running the following command with result in : 

```bash
awssel load --service-name sd-web --prefix-path /os/prod/support/IT/core  --aws-region eu-west-1 --filter-pattern=".*URL"  --endpoint http://localhost:4583 --export

# Ouput >
# 
# OS_SDWEB_HTTP_URL=http://www.my-web-service.com
# OS_SDWEB_HTTPS_URL=https://www.my-web-service.com
# OS_SDWEB_DOMAIN=my-web-service.com
```

Now, let says we want to get only the HTTPS URL of our `sd-web` service; we can use the option `--filter-pattern` to do so:


```bash
awssel load --filter-pattern='.*HTTPS.*_URL'--service-name sd-web --prefix-path /os/prod/support/IT/core  --aws-region eu-west-1   --endpoint http://localhost:4583 

# Ouput >
# 
# OS_SDWEB_HTTPS_URL=https://www.my-web-service.com
```

> I used a slighly complicated regex to show you that the tool support any regex, as long as it is valid.
> See [References] to get an idea of the regex syntax supported by `awssel`.

## TODO

* [ ] List errors about AWS Credentials (FAQ)
* [ ] Add a section explaining how to run this tool in an AWS

## Tests

**awssel** is battle tested. We use the [localstack](https://github.com/localstack/localstack) project to mock AWS services.
In fact, before launch our tests, we prepare our test environment by starting an AWS SSM Mock; running at `http://localhost:4583`.

Preparing the test environment, involves executing the following test fixtures: 

- Start localstack AWS SSM Service
  ```bash
  rake code:test:prepare  # Start testing env
  ```
- Add test entries into the parameter store. Look into the file [tasks/populate.rb](tasks/populate.rb) for more details.
  ```bash
  rake code:test:seed  # Populate SSM with test values
  ```

At this point, the test environment is ready and you can run the following command to launch our tests:

```bash
rake code:test:run # Run all tests
```

##### CAUTION
> You choose to focus on testing the business logic behind **awssel** and not the CLI UI.
> Feel free to contribute; see [Contribute](#contribute) section to know how to contribute.


### QA


## Guidelines

See the [guidelines doc].

## Contribute

See the [contributing doc].

## FAQ

See the [faq doc].

## References

* [Golang regex syntax](https://github.com/google/re2/wiki/Syntax)

## Maintainers

* Lionel T. [@lktslionel](https://twitter.com/lktslionel)

## License
 
[MIT license]


[Changelog]: docs/CHANGELOG.md
[contributing doc]: docs/CONTRIBUTE.md
[guidelines doc]: docs/GUIDELINES.md
[faq doc]: docs/FAQ.md
[MIT license]: LICENSE
[Install]: #Install
[Usage]: #Usage
[Examples]: #Examples
[TODO]: #TODO
[Tests]: #Tests
[Guidelines]: #Guidelines
[Contribute]: #Contribute
[FAQ]: #FAQ
[References]: #References
[Maintainers]: #Maintainers
[License]: #License