package main

import (
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/etcdv3"
	_ "github.com/micro/go-plugins/transport/nats"
)
