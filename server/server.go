/**
* @Author: cl
* @Date: 2021/2/23 14:14
 */
package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	openzipkin "github.com/openzipkin/zipkin-go-opentracing"
	"google.golang.org/grpc"
	"math"
	"net"
	hpb "server/proto/hello"
	"time"
)

func main() {
	go InitServer1()
	go InitServer2()

	select {}
}

func InitTracer(hostPort, serverName string) (opentracing.Tracer, error) {
	zipkinUrl := "http://localhost:9411/api/v1/spans"
	collector, err := openzipkin.NewHTTPCollector(zipkinUrl)
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %v", err)
		return nil, err
	}
	tracer, err := openzipkin.NewTracer(
		openzipkin.NewRecorder(collector, true, hostPort, serverName),
		openzipkin.ClientServerSameSpan(true),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %v", err)
		return nil, err
	}

	opentracing.SetGlobalTracer(tracer)

	return tracer, nil
}

func InitServer1() {
	addr := "0.0.0.0:1111"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	tracer, err := InitTracer(addr, "cl:server_1111")
	if err != nil {
		fmt.Printf("failed to InitTracer: %v\n", err)
		return
	}
	fmt.Printf("server listening on port: %s\n", addr)
	var dopts []grpc.ServerOption
	dopts = append(dopts, grpc.MaxRecvMsgSize(math.MaxInt64))
	dopts = append(dopts, grpc.MaxConcurrentStreams(math.MaxInt32))
	dopts = append(dopts, grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())))
	grpcDServer := grpc.NewServer(dopts...)
	hpb.RegisterHello1Server(grpcDServer, &HelloServer1{})
	err = grpcDServer.Serve(lis)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("server listening on port 1111")
	}
}

func InitServer2() {
	addr := "0.0.0.0:2222"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	tracer, err := InitTracer(addr, "cl:server_2222")
	if err != nil {
		fmt.Printf("failed to InitTracer: %v\n", err)
		return
	}
	fmt.Printf("server listening on port: %s\n", addr)
	var dopts []grpc.ServerOption
	dopts = append(dopts, grpc.MaxRecvMsgSize(math.MaxInt64))
	dopts = append(dopts, grpc.MaxConcurrentStreams(math.MaxInt32))
	dopts = append(dopts, grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())))
	grpcDServer := grpc.NewServer(dopts...)
	hpb.RegisterHello2Server(grpcDServer, &HelloServer2{})
	err = grpcDServer.Serve(lis)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("server listening on port 2222")
	}
}

type HelloServer1 struct {
}

func (rpc *HelloServer1) GetHello1(ctx context.Context, req *hpb.HelloReq) (resp *hpb.HelloResp, err error) {
	<-time.After(time.Second * 2)
	return ProcessGetHello1(ctx, req)
}

func ProcessGetHello1(ctx context.Context, req *hpb.HelloReq) (resp *hpb.HelloResp, err error) {
	addr := "10.107.63.170:2222"
	tracer, err := InitTracer(addr, "cl:client_1111")
	if err != nil {
		fmt.Printf("failed to InitTracer: %v\n", err)
		return
	}
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())))
	if err != nil {
		fmt.Printf("failed to Dial: %v\n", err)
		return
	}
	defer conn.Close()

	client := hpb.NewHello2Client(conn)
	resp, err = client.GetHello2(context.Background(), req)
	if err != nil {
		fmt.Printf("failed to GetHello: %v\n", err)
		return
	}
	fmt.Printf("get GetHello response:%v\n", resp)
	return
}

type HelloServer2 struct {
}

func (rpc *HelloServer2) GetHello2(ctx context.Context, req *hpb.HelloReq) (resp *hpb.HelloResp, err error) {
	<-time.After(time.Second * 2)
	return ProcessGetHello2(ctx, req)
}

func ProcessGetHello2(ctx context.Context, req *hpb.HelloReq) (resp *hpb.HelloResp, err error) {
	resp = &hpb.HelloResp{
		Key: "ProcessGetHello2" + "----" + req.Key,
		Val: "success",
	}
	return resp, nil
}
