package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yeyus/dmr-entities/cmd/server/services"
	"github.com/yeyus/dmr-entities/internal/pkg/db"
	"github.com/yeyus/dmr-entities/pkg/api"
	multiint "github.com/yeyus/go-grpc-interceptor/multiinterceptor"
	"github.com/yeyus/go-grpc-interceptor/xrequestid"
	"google.golang.org/grpc"
)

type Configuration struct {
	DatabaseHost     string `envconfig:"DB_HOST"`
	DatabaseUser     string `envconfig:"DB_USER"`
	DatabasePassword string `envconfig:"DB_PASSWORD"`
	DatabaseName     string `envconfig:"DB_NAME"`
	Debug            bool   `envconfig:"DEBUG"`
}

func main() {
	var c Configuration
	err := envconfig.Process("server", &c)
	if err != nil {
		log.Fatal(err.Error)
	}

	grpcHost := fmt.Sprintf(":%d", 3000)
	log.Printf("Launching grpc server at %s", grpcHost)
	lis, err := net.Listen("tcp", grpcHost)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dbManager := db.GetDBManager(c.DatabaseHost, c.DatabaseUser, c.DatabasePassword, c.DatabaseName, c.Debug)
	s := services.EntitiesServer{
		DB: dbManager,
	}

	// prometheus metrics
	reg := prometheus.NewRegistry()
	grpcMetrics := grpc_prometheus.NewServerMetrics()

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(multiint.NewMultiStreamServerInterceptor(grpc_prometheus.StreamServerInterceptor, xrequestid.StreamServerInterceptor())),
		grpc.UnaryInterceptor(multiint.NewMultiUnaryServerInterceptor(grpc_prometheus.UnaryServerInterceptor, xrequestid.UnaryServerInterceptor())),
	)
	api.RegisterEntitiesServer(grpcServer, &s)
	grpcMetrics.InitializeMetrics(grpcServer)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	go func() {
		serverHost := fmt.Sprintf("0.0.0.0:%d", 8080)
		log.Printf("Launching http metrics server at %s", serverHost)
		if err := http.ListenAndServe(serverHost, nil); err != nil {
			log.Fatalf("Unable to start a http server: %s", err)
		}
	}()

	log.Fatal(grpcServer.Serve(lis))
}
