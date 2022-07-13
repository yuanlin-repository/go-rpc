package server

import (
	"context"
	"fmt"
	"go-rpc/interceptor"
	"reflect"
)

type Server struct {
	// server config
	opts    *ServerOptions
	service Service

	closing bool
}

type emptyInterface interface{}

type emptyService struct{}

func (s *Server) Start() {
	fmt.Println("start server..")
}

func (s *Server) RegisterService(serviceName string, svr interface{}) error {
	svrType := reflect.TypeOf(svr)
	svrValue := reflect.ValueOf(svr)

	sd := &ServiceDesc{
		ServiceName: serviceName,
		Svr:         svr,
		HandlerType: (*emptyInterface)(nil),
	}

	methods, err := getServiceMethods(svrType, svrValue)

}

func getServiceMethods(svrType reflect.Type, value reflect.Value) ([]*MethodDesc, error) {
	var methods []*MethodDesc

	for i := 0; i < svrType.NumMethod(); i++ {
		method := svrType.Method(i)
		if err := checkMethod(method.Type); err != nil {
			return nil, err
		}
		methodHandler := func(ctx context.Context, svr interface{}, dec func(interface{}) error, ceps []interceptor.ServerInterceptor) (interface{}, error) {

		}
	}
}

// 根据服务器参数创建server
func NewServer(opt ...ServerOption) *Server {
	s := &Server{
		opts: &ServerOptions{},
	}

	// insert config
	for _, o := range opt {
		o(s.opts)
	}

	// 创建 service
	s.service = NewService(s.opts)

	return s
}

func NewService(opts *ServerOptions) Service {
	return &service{
		opts: opts,
	}
}
