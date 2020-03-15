package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/hypoballad/toecutter/toecuttergrpc"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", "localhost:50051", "server address")
var name = flag.String("name", "", "hashes|reward|account/stats|account/site|account/payments|sitestats|payment1mhash")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := toecuttergrpc.NewCoinimpClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	switch *name {
	case "hashes":
		empty := &toecuttergrpc.Empty{}
		respHashes, err := c.Hashes(ctx, empty)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("message: %d, status: %s\n", respHashes.GetMessage(), respHashes.GetStatus())
	case "reward":
		empty := &toecuttergrpc.Empty{}
		respReward, err := c.Reward(ctx, empty)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("message: %s, status: %s\n", respReward.GetMessage(), respReward.GetStatus())
	case "account/stats":
		empty := &toecuttergrpc.Empty{}
		respStats, err := c.AccountStats(ctx, empty)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("hasherate: %f, hashes: %d, reward: %s, rewardpending: %s, status: %s\n", respStats.GetHasherate(), respStats.GetHashes(), respStats.GetReward(), respStats.GetRewardpending(), respStats.GetStatus())
	case "account/site":
		empty := &toecuttergrpc.Empty{}
		respSite, err := c.AccountSite(ctx, empty)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("message: %v, status: %s\n", respSite.GetSites(), respSite.GetStatus())
	case "account/payments":
		empty := &toecuttergrpc.Empty{}
		respPayments, err := c.AccountPayments(ctx, empty)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("payments: %v, status: %s\n", respPayments.GetPayments(), respPayments.GetStatus())
	case "sitestats":
		empty := &toecuttergrpc.Empty{}
		respSiteStats, err := c.SiteStats(ctx, empty)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("name: %s, hashrate: %f, hashes: %d, reward: %s\n", respSiteStats.GetName(), respSiteStats.GetHashrate(), respSiteStats.GetHashes, respSiteStats.GetReward())
	case "payment1mhash":
		empty := &toecuttergrpc.Empty{}
		respPayout, err := c.Payout1MHash(ctx, empty)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("result: %s, status: %s\n", respPayout.GetResult(), respPayout.GetStatus())
	default:
		log.Println("name required")
	}
	os.Exit(0)
}
