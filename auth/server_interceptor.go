package auth

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ServerInterceptor struct {
	jwtManager      *JWTManager
	accessibleRoles map[string][]string
}

func NewServerInterceptor(jwtManager *JWTManager, accessibleRoles map[string][]string) *ServerInterceptor {
	return &ServerInterceptor{
		jwtManager:      jwtManager,
		accessibleRoles: accessibleRoles,
	}
}

func (interceptor *ServerInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("-->Unary interceptor: ", info.FullMethod)
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (interceptor *ServerInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	// if no role is required to access the method
	if !ok {
		// log.Println("No role required, public access:", accessibleRoles, interceptor.accessibleRoles)
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Metadata not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "Authorization token not provied")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Access token is invalid: %v", err)
	}

	for _, role := range accessibleRoles {
		if role == claims.Role {
			log.Printf("Role exists: %v", role)
			return nil
		}
	}
	return status.Errorf(codes.PermissionDenied, "Permission not granted to access this rpc")
}
