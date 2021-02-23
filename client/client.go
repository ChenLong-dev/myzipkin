/**
* @Author: cl
* @Date: 2021/2/23 14:20
 */
package main

import (
	hpb "client/proto/hello"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	openzipkin "github.com/openzipkin/zipkin-go-opentracing"
	"google.golang.org/grpc"
	"time"
)

func InitTracer(hostPort, serverName string) (opentracing.Tracer, error) {
	zipkinUrl := "http://localhost:9411/api/v1/spans"
	collector, err := openzipkin.NewHTTPCollector(zipkinUrl)
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %v\n", err)
		return nil, err
	}
	tracer, err := openzipkin.NewTracer(
		openzipkin.NewRecorder(collector, true, hostPort, serverName),
		openzipkin.ClientServerSameSpan(true),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %v\n", err)
		return nil, err
	}

	opentracing.SetGlobalTracer(tracer)

	return tracer, nil
}

func client() {
	addr := "10.107.63.170:1111"
	tracer, err := InitTracer(addr, "cl:client")
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

	req := &hpb.HelloReq{
		Key: "req",
		Val: addr,
	}
	client := hpb.NewHello1Client(conn)
	resp, err := client.GetHello1(context.Background(), req)
	if err != nil {
		fmt.Printf("failed to GetHello: %v\n", err)
		return
	}
	fmt.Printf("get GetHello response:%v\n", resp)
}

func main() {
	for {
		client()
		<-time.After(time.Second * 10)
	}
}
