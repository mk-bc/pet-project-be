package auth

import (
	"context"
	"log"
	"time"

	"github.com/mk-bc/pet-project-be/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ClientInterceptor struct {
	client      *client.JobPortalServiceClient
	authMethods map[string]bool
	accessToken string
}

func NewClientInterceptor(
	client *client.JobPortalServiceClient,
	authMethods map[string]bool,
	refreshDuration time.Duration) (*ClientInterceptor, error) {
	interceptor := &ClientInterceptor{
		client:      client,
		authMethods: authMethods,
	}

	err := interceptor.scheduleRefreshToken(refreshDuration)
	if err != nil {
		return nil, err
	}
	return interceptor, nil
}

func (interceptor *ClientInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {
		log.Println("-->Unary Client Interceptor", method)
		if interceptor.authMethods[method] {
			return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (interceptor *ClientInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.accessToken)
}

func (interceptor *ClientInterceptor) scheduleRefreshToken(refreshDuration time.Duration) error {
	err := interceptor.refreshToken()
	if err != nil {
		return err
	}

	go func() {
		wait := refreshDuration
		for {
			time.Sleep(wait)
			err := interceptor.refreshToken()
			if err != nil {
				wait = time.Second
			} else {
				wait = refreshDuration
			}
		}
	}()
	return nil
}

func (interceptor *ClientInterceptor) refreshToken() error {
	accessToken, err := interceptor.client.Login()
	if err != nil {
		return err
	}
	interceptor.accessToken = accessToken
	log.Printf("token refreshed: %v", accessToken)
	return nil
}
