package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yeyus/dmr-entities/pkg/api"
	"google.golang.org/grpc"
)

func printEntity(entity *api.Entity) {
	jsonvalue, _ := json.MarshalIndent(entity, "", "  ")
	fmt.Printf("%s", jsonvalue)
}

func printEntityResponse(entityResponse *api.EntitiesResponse) {
	jsonvalue, _ := json.MarshalIndent(entityResponse, "", "  ")
	fmt.Printf("%s", jsonvalue)
}

func main() {
	host := flag.String("host", "localhost:3000", "host where the client will connect")
	byCallsignCmd := flag.NewFlagSet("GetEntitiesByCallsign", flag.ExitOnError)
	byEntityIDCmd := flag.NewFlagSet("GetEntity", flag.ExitOnError)
	bycEntityID := byCallsignCmd.Uint64("entityId", 1, "record entity id")
	bycCallsign := byCallsignCmd.String("callsign", "EA7JMF", "amateur radio callsign")

	if len(os.Args) < 2 {
		fmt.Println("Error: expected a command (GetEntitiesByCallsign, GetEntity, etc.)")
		os.Exit(1)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(*host, grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("did not connect: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := api.NewEntitiesClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	switch os.Args[1] {
	case "GetEntity":
		byEntityIDCmd.Parse(os.Args[2:])
		fmt.Printf("Command is %s and argument is %d\n", "GetEntity", *bycEntityID)
		response, err := client.GetEntity(ctx, &api.GetEntityRequest{
			Id: *bycEntityID,
		})
		if err != nil {
			log.Panicf("Error: performing %s request to %s: %s", "GetEntity", *host, err)
		}
		printEntity(response)
	case "GetEntitiesByCallsign":
		byCallsignCmd.Parse(os.Args[2:])
		fmt.Printf("Command is %s and argument is %s\n", "GetEntitiesByCallsign", *bycCallsign)
		callsign := make([]string, 1)
		callsign[0] = *bycCallsign
		response, err := client.GetEntitiesByCallsign(ctx, &api.GetEntitiesByCallsignRequest{
			Callsign: callsign,
		})
		if err != nil {
			log.Panicf("Error: performing %s request to %s: %s", "GetEntitiesByCallsign", *host, err)
		}
		printEntityResponse(response)
	default:
		log.Panic(fmt.Printf("Error: entered command %s expected a command: GetEntitiesByCallsign\n", os.Args[1]))
	}

}
