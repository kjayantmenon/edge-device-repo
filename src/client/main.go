package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/kjayantmenon/assetregistry/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:58086", "Asset Service Server address")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewAssetRegistryClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetAllAssets(ctx, &pb.NoAssetParams{})
	if err != nil {
		log.Fatalf("could not get Asset Listreet: %v", err)
	}
	log.Printf("Assets: %s", r.Assets)
}
