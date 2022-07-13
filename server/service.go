package server

import (
	"context"
	"fmt"
	"go-rpc/interceptor"
)

type Service interface {
	Register(string, Handler)
	Serve(*ServerOptions)
	Close()
	Name() string
}

type service struct {
	// server
	svr interface{}
	// Each service is managed in one context
	ctx context.Context
	// context controller
	cancel context.CancelFunc
	// service name
	serviceName string
	// serviceMethodName -> methodHandler
	handlers map[string]Handler
	opts     *ServerOptions

	// whether the service is closing
	closing bool
}

func (s *service) Name() string {
	return s.serviceName
}

type ServiceDesc struct {
	Svr         interface{}
	ServiceName string
	Methods     []*MethodDesc
	HandlerType interface{}
}

type MethodDesc struct {
	MethodName string
	Handler    Handler
}

// Handler is the handler of a method
type Handler func(context.Context, interface{}, func(interface{}) error, []interceptor.ServerInterceptor) (interface{}, error)

// register method handler for service
func (s *service) Register(handlerName string, handler Handler) {
	if s.handlers == nil {
		s.handlers = make(map[string]Handler)
	}
	s.handlers[handlerName] = handler
}

func (s *service) Serve(options *ServerOptions) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Close() {
	s.closing = true
	if s.cancel != nil {
		s.cancel()
	}
	fmt.Println("service closing ...")
}
