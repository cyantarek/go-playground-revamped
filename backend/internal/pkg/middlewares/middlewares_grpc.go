package middlewares

import (
	"backend/config"
	"backend/pkg/jwtpacker"
	"errors"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (mw Middleware) JwtInterceptorGRPC(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	if config.Cfg.AuthSkipper[info.FullMethod] {
		return handler(ctx, req)
	}

	m, _ := metadata.FromIncomingContext(ctx)

	if len(m.Get("authorization")) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized: token not found")
	}

	token := m.Get("authorization")[0]

	if token == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized: token not found")
	}

	// extract token from header
	tokenParts := strings.Split(token, " ")

	if len(tokenParts) != 2 {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized: token not found")
	}

	ownerID, err := mw.rawJWT(tokenParts[1])
	//ownerID, err := mw.firebase(tokenParts[1])
	if err != nil {
		return "", status.Error(codes.Unauthenticated, err.Error())
	}

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("ownerID", ownerID))

	// then call the handler
	return handler(ctx, req)
}

func (mw Middleware) rawJWT(token string) (string, error) {
	if valid, claims := jwtpacker.ValidateToken(token); valid {
		if email, ok := claims["email"]; ok && email != "" {
			return email.(string), nil
		} else {
			return "", errors.New("unauthorized: email on token not found")
		}
	} else {
		return "", errors.New("unauthorized: token expired. Please use refresh token")
	}
}
