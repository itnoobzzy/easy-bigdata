package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	cserv1 "easyCasbin/api/casbin_rule/v1"
	rsv1 "easyCasbin/api/role/v1"
	userv1 "easyCasbin/api/user/v1"
	"easyCasbin/internal/conf"
	"easyCasbin/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	usvc *service.UserService,
	rs *service.DomainRoleService,
	cs *service.CasbinRuleService,
	logger log.Logger,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	userv1.RegisterUserServer(srv, usvc)
	rsv1.RegisterDomainRoleServer(srv, rs)
	cserv1.RegisterCasbinRuleServer(srv, cs)
	return srv
}
