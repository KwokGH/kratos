package server

import (
	"context"
	"fmt"
	"github.com/KwokGH/kratos/api/v1/user"
	"github.com/KwokGH/kratos/internal/conf"
	"github.com/KwokGH/kratos/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server,
	userService *service.UserService,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			selector.Server(
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte("123"), nil
				}, jwt.WithSigningMethod(jwt2.SigningMethodHS256), jwt.WithClaims(func() jwt2.Claims {
					return &jwt2.MapClaims{}
				}))).
				Match(NewWhiteListMatcher()).Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
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

	user.RegisterUserHTTPServer(srv, userService)

	return srv
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.v1.user.User/Login"] = struct{}{}
	whiteList["/api.v1.user.User/GetUserDetail"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		fmt.Println(operation)
		if _, ok := whiteList[operation]; ok {
			if t, ok := transport.FromServerContext(ctx); ok {
				fmt.Println(t.RequestHeader().Get("Authorization"))
				if len(t.RequestHeader().Get("Authorization")) > 0 {
					return true
				}
			}
			return false
		}
		return true
	}
}

func MyMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if _, ok := transport.FromServerContext(ctx); ok {
				return nil, nil
			}
			return nil, nil
		}
	}
}
