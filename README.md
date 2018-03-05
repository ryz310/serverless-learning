# Learning Serverless Framework with go-lang

## Setup

Install Serverless Framework via npm.

```
$ npm install -g serverless
```

Set-up your Provider Credentials.

```
$ serverless config credentials --provider aws --key {AWS Access Key ID} --secret {AWS Secret Access Key}
```

Install golang via brew if you don't install yet.

```
$ brew install go
```

Install `dep` what dependency management tool for Go.

```
$ brew install dep
```

And make sure you're in your ${GOPATH}/src directory.

## Deploy a Service

```
$ make
$ serverless deploy -v
```

## Remove the Service

```
$ serverless remove
```
