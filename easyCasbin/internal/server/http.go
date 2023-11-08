package server

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwtv4 "github.com/golang-jwt/jwt/v4"

	cv1 "easyCasbin/api/casbin_rule/v1"
	db "easyCasbin/api/db"
	rv1 "easyCasbin/api/role/v1"
	uv1 "easyCasbin/api/user/v1"
	"easyCasbin/internal/conf"
	"easyCasbin/internal/service"
)

func NewWhiteListMatcher(w string) selector.MatchFunc {
	whiteList := make(map[string]bool)
	wl := strings.Split(w, ",")
	for _, u := range wl {
		whiteList[strings.TrimSpace(u)] = true
	}
	//whiteList["/api.user.v1.User/Login"] = true
	//whiteList["/api.user.v1.User/register"] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	us *service.UserService,
	rs *service.DomainRoleService,
	cs *service.CasbinRuleService,
	ds *service.DbIniterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			validate.Validator(),
			selector.Server(
				jwt.Server(
					func(token *jwtv4.Token) (interface{}, error) {
						return []byte(c.Jwt.SigningKey), nil
					},
					//jwt.WithClaims(func() jwtv4.Claims {
					//	return &mjwt.CustomClaims{}
					//}),
				),
			).Match(NewWhiteListMatcher(c.Jwt.WhiteList)).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// swagger api
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)

	uv1.RegisterUserHTTPServer(srv, us)
	db.RegisterInitDBHTTPServer(srv, ds)
	rv1.RegisterDomainRoleHTTPServer(srv, rs)
	cv1.RegisterCasbinRuleHTTPServer(srv, cs)
	return srv
}
