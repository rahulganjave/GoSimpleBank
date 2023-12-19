package main

import (
	"context"
	"gosimplebank/api"
	db "gosimplebank/db/sqlc"
	"gosimplebank/pb"
	"gosimplebank/util"
	"log"
	"net"

	"gosimplebank/gapi"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		//log.Fatal().Err(err).Msg("cannot load config")
		log.Fatal("cannot load config:", err)
	}

	// if config.Environment == "development" {
	// 	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// }

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(connPool) //DB
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store) /*, taskDistributor*/
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	// gprcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store) //Gin Server
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
