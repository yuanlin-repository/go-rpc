package server

import "time"

type ServerOptions struct {
	// listening address, e.g. : (ip://127.0.0.1:8080)
	address string
	// network type, e.g. : tcp、udp
	network string
	// protocol type, default: proto
	protocol string
	// timeout
	timeout time.Duration
	// serialization type, default: proto
	serializationType string
}

type ServerOption func(*ServerOptions)

func WithAddress(address string) ServerOption {
	return func(o *ServerOptions) {
		o.address = address
	}
}

func WithNetWork(network string) ServerOption {
	return func(o *ServerOptions) {
		o.network = network
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(o *ServerOptions) {
		o.timeout = timeout
	}
}

func WithProtocol(protocol string) ServerOption {
	return func(o *ServerOptions) {
		o.protocol = protocol
	}
}

func WithSerialization(serializationType string) ServerOption {
	return func(o *ServerOptions) {
		o.serializationType = serializationType
	}
}
