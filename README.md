###### OLST / TOOLING
# AWSSEL (AWS Ssm Env Loader)


A tool that load all env vars stored in AWS SSM Parameter Store, for a given service.

See [Our Moviations doc](docs/awssel/motivations.md) for more information about why we build this tool.
See [Our Architecture doc](docs/awssel/how-it-works.md) to understand how things work under the hood.

## Install

To install the tool from your shell, run:

```
go get -u github.com/Stores-Discount/awssel
```

## Usage

```
awssel <command> [options]
```

### Command

**Awssel** has a only 2 commands : 



Command | Description | Options
---------|----------|---------
 help | Show command help text | -
 load | Load env vars for a service | See [Load options](docs/awssel/cmd/load.md)


## Examples

Get all environment variables for `proxy-web` service

```bash
awssel load --service-name proxy-web
```



## Tests


### QA


## Guidelines

See the [guidelines doc].

## Contribute

See the [contributing doc].

## FAQ

See the [faq doc].

## Maintainers

* Lionel T. [@lktslionel](https://twitter.com/lktslionel)

## License
 
[MIT license]


[Changelog]: docs/CHANGELOG.md
[contributing doc]: docs/CONTRIBUTE.md
[guidelines doc]: docs/GUIDELINES.md
[faq doc]: docs/FAQ.md
[MIT license]: LICENSE