package main

import (
	"context"
	"log"
	"net"
	"strconv"

	pb "github.com/kjayantmenon/assetregistry/proto"
	"google.golang.org/grpc"
)

var addr = "0.0.0.0:58086"

type AssetRegistryServer struct {
	pb.AssetRegistryServer
}

func main() {
	log.Printf("Starting the Asset Registry Service...")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Encountered err: %v\n", err)
	}

	log.Printf("Asset Registry Service listening on %s\n", lis.Addr())

	ars := grpc.NewServer()
	pb.RegisterAssetRegistryServer(ars, &AssetRegistryServer{})
	if err = ars.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

	log.Printf("!!!Asset Registry Service started!!! ")
}

func (ars *AssetRegistryServer) GetAllAssets(ctx context.Context, in *pb.NoAssetParams) (*pb.AssetList, error) {
	al := &pb.AssetList{}
	// a := &pb.Asset{}
	// a.AssetId = "1"

	var assets []*pb.Asset
	for i := 0; i <= 5; i++ {
		a := createAsset(strconv.Itoa(i))
		assets = append(assets, a)
	}
	// assets = append(assets, a)
	// a = &pb.Asset{}
	// a.AssetId = "2"
	// assets = append(assets, a)

	al.Assets = assets
	//_ = append(al.Assets, a)
	return al, nil
}

func createAsset(id string) *pb.Asset {
	a := &pb.Asset{}
	a.AssetId = id
	a.Name = "asset-" + id
	a.Description = "Test Description for asset-" + id
	a.Protocol = "opcua"
	a.PrimaryEndpoint = "101.10.10." + id
	a.SecondaryEndpoint = "101.10.20." + id
	return a
}
