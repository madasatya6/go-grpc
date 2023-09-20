package handler

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)


func authMethods() map[string]bool {
	return map[string]bool{
		"/transaction.TransactionService/ApproveTransactions": true,
	}
}

func publicMethods() map[string]bool {
	return map[string]bool{
		"/user.UserService/GreetUser": true,
	}
}

func AuthUser(ctx context.Context) (err error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metadata")
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	log.Printf("Metadata %v+\n", md)

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	if accessToken != "123456" {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return nil
}

func (h *handlerService) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("UnaryServerInterceptor PRE", info.FullMethod)

	authMethods := authMethods()
	if err := AuthUser(ctx); err != nil && authMethods[info.FullMethod] {
		return resp, err
	}

	m, err := handler(ctx, req)

	log.Println("UnaryServerInterceptor POST", info.FullMethod)

	return m, err
}

func (h *handlerService) StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("StreamServerInterceptor PRE", info.FullMethod)

	authMethods := authMethods()
	if err := AuthUser(ss.Context()); err != nil && authMethods[info.FullMethod] {
		return err
	}

	err := handler(srv, newWrappedStream(ss))

	log.Println("StreamServerInterceptor POST")

	return err
} 

type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Received %T - %v\n", m, m)

	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Sent %T - %v\n", m, m)

	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}
