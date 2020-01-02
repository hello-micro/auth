# Auth Service

This is the Auth service

[![Build Status](https://cloud.drone.io/api/badges/hello-micro/auth/status.svg)](https://cloud.drone.io/hello-micro/auth)
[![LINK](https://img.shields.io/badge/link-Github-%23FF4D5B.svg?style=flat-square)](https://github.com/hello-micro/auth) 
[![Go Report Card](https://goreportcard.com/badge/github.com/hello-micro/auth)](https://goreportcard.com/report/github.com/hello-micro/auth)
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg?style=flat-square)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

Generated with

```
micro new github.com/hello-micro/auth --namespace=hello.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: hello.micro.srv.auth
- Type: srv
- Alias: auth

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./auth-srv
// Or
./auth-srv --broker=nats --registry=etcd --transport=nats 
```

Build a docker image
```
make docker
```