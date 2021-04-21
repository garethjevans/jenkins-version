[![Go Report Card](https://goreportcard.com/badge/github.com/jenkins-infra/jenkins-version)](https://goreportcard.com/report/github.com/jenkins-infra/jenkins-version)
[![Downloads](https://img.shields.io/github/downloads/jenkins-infra/jenkins-version/total.svg)]()

# jenkins-version

a small CLI that can be used to determine the latest jenkins verision from maven metadata

## To Install

```
brew tap jenkins-infra/tap
brew install jv
```

This can be used a docker container with the following:

```
docker run -it jenkins-infra/jv:main
```

## Usage

To get the latest weekly release:

```
jv get [--username <username> --password <password>]
```

To get the latest LTS release:

```
jv get --version-identifier lts [--username <username> --password <password>]
```

To get the latest LTS for a particular release train:

```
jv get --version-identifier 2.249 [--username <username> --password <password>]
```

## Documentation

More indepth documentation can be found [here](./docs/jv.md)

## Development

To build the application:

```
make build
```

To test:

```
make test
```
