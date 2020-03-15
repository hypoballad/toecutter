package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	"github.com/hypoballad/toecutter"
	"github.com/hypoballad/toecutter/toecuttergrpc"
	"github.com/robfig/cron/v3"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc"
)

var db *leveldb.DB

var addr = flag.String("addr", ":50051", "server address")
var dbDir = flag.String("db", "toecutter_db", "level db dir")
var configpath = flag.String("conf", "toecutter.toml", "config")
var modeDebug = flag.Bool("debug", false, "set debug mode")

func Jobs(conf toecutter.Config, db *leveldb.DB, debug bool) {
	if debug {
		log.Println("hashes")
	}
	respHashes, err := toecutter.Hashes(conf, toecutter.Mintme, conf.Main.SiteKey)
	if err != nil {
		log.Println(err)
	} else {
		if debug {
			log.Println(respHashes)
		}
		if err := toecutter.SetHashes(db, respHashes); err != nil {
			log.Println(err)
		}
	}
	time.Sleep(time.Duration(conf.Main.Sleep) * time.Second)
	if debug {
		log.Println("reward")
	}
	respReward, err := toecutter.Reward(conf, toecutter.Mintme, conf.Main.SiteKey)
	if err != nil {
		log.Println(err)
	} else {
		if debug {
			log.Println(respReward)
		}
		if err := toecutter.SetReward(db, respReward); err != nil {
			log.Println(err)
		}
	}
	time.Sleep(time.Duration(conf.Main.Sleep) * time.Second)
	if debug {
		log.Println("AccountStats")
	}
	respAccountStats, err := toecutter.AccountStats(conf, toecutter.Mintme)
	if err != nil {
		log.Println(err)
	} else {
		if debug {
			log.Println(respAccountStats)
		}
		if err := toecutter.SetAccountStats(db, respAccountStats); err != nil {
			log.Println(err)
		}
	}
	time.Sleep(time.Duration(conf.Main.Sleep) * time.Second)
	if debug {
		log.Println("AccountSite")
	}
	respAccountSite, err := toecutter.AccountSite(conf, toecutter.Mintme)
	if err != nil {
		log.Println(err)
	} else {
		if debug {
			log.Println(respAccountSite)
		}
		if err := toecutter.SetAccountSite(db, respAccountSite); err != nil {
			log.Println(err)
		}
	}
	time.Sleep(time.Duration(conf.Main.Sleep) * time.Second)
	if debug {
		log.Println("AccountPayments")
	}
	respAccountPayments, err := toecutter.AccountPayments(conf, toecutter.Mintme)
	if err != nil {
		log.Println(err)
	} else {
		if debug {
			log.Println(respAccountPayments)
		}
		if err := toecutter.SetAccountPayments(db, respAccountPayments); err != nil {
			log.Println(err)
		}
	}
	time.Sleep(time.Duration(conf.Main.Sleep) * time.Second)
	if debug {
		log.Println("SiteStats")
	}
	respSiteStats, err := toecutter.SiteStats(conf)
	if err != nil {
		log.Println(err)
	} else {
		if debug {
			log.Println(respSiteStats)
		}
		if err := toecutter.SetSiteStats(db, respSiteStats); err != nil {
			log.Println(err)
		}
	}
	time.Sleep(time.Duration(conf.Main.Sleep) * time.Second)
	if debug {
		log.Println("Payout1MHash")
	}
	respPayout, err := toecutter.Payout1MHash(conf, toecutter.Mintme)
	if err != nil {
		log.Println(err)
	} else {
		if debug {
			log.Println(respPayout)
		}
		if err := toecutter.SetPayout1MHash(db, respPayout); err != nil {
			log.Println(err)
		}
	}
	if debug {
		log.Println("end")
	}
}

type server struct {
	toecuttergrpc.UnimplementedCoinimpServer
}

func (s server) Hashes(ctx context.Context, in *toecuttergrpc.Empty) (*toecuttergrpc.RespHashes, error) {
	var respHashes *toecuttergrpc.RespHashes
	resp, err := toecutter.GetHashes(db)
	if err != nil {
		return respHashes, err
	}
	respHashes = &toecuttergrpc.RespHashes{Message: resp.Message, Status: resp.Status}
	return respHashes, nil
}

func (s server) Reward(ctx context.Context, in *toecuttergrpc.Empty) (*toecuttergrpc.RespReward, error) {
	var respReward *toecuttergrpc.RespReward
	resp, err := toecutter.GetReward(db)
	if err != nil {
		return respReward, err
	}
	respReward = &toecuttergrpc.RespReward{Message: resp.Message, Status: resp.Status}
	return respReward, nil
}

func (s server) AccountStats(ctx context.Context, in *toecuttergrpc.Empty) (*toecuttergrpc.RespStats, error) {
	var respStats *toecuttergrpc.RespStats
	resp, err := toecutter.GetAccountStats(db)
	if err != nil {
		return respStats, err
	}
	respStats = &toecuttergrpc.RespStats{
		Hasherate:     resp.Message.Hashrate,
		Hashes:        resp.Message.Hashes,
		Reward:        resp.Message.Reward,
		Rewardpending: resp.Message.RewardPending,
		Status:        resp.Status,
	}
	return respStats, nil
}

func (s server) AccountSite(ctx context.Context, in *toecuttergrpc.Empty) (*toecuttergrpc.RespSite, error) {
	var respSite *toecuttergrpc.RespSite
	resp, err := toecutter.GetAccountSite(db)
	if err != nil {
		return respSite, err
	}
	sites := []*toecuttergrpc.Site{}
	for _, site := range resp.Message {
		sites = append(sites, &toecuttergrpc.Site{Name: site.Name, Sitekey: site.SiteKey})
	}
	respSite = &toecuttergrpc.RespSite{
		Sites:  sites,
		Status: resp.Status,
	}
	return respSite, nil
}

func (s server) AccountPayments(ctx context.Context, in *toecuttergrpc.Empty) (*toecuttergrpc.RespPayments, error) {
	var respAccountPayments *toecuttergrpc.RespPayments
	resp, err := toecutter.GetAccountPayments(db)
	if err != nil {
		return respAccountPayments, err
	}
	payments := []*toecuttergrpc.Payment{}
	for _, payment := range resp.Message {
		payments = append(payments, &toecuttergrpc.Payment{
			Amount:        payment.Amount,
			Fee:           payment.Fee,
			Status:        payment.Status,
			Walletaddress: payment.WalletAddress,
			Txid:          payment.TxID,
			Timestamp:     payment.Timestamp,
		})
	}
	respAccountPayments = &toecuttergrpc.RespPayments{
		Payments: payments,
		Status:   resp.Status,
	}
	return respAccountPayments, nil
}

func (s server) SiteStats(ctx context.Context, in *toecuttergrpc.Empty) (*toecuttergrpc.RespSiteStats, error) {
	var respSiteStats *toecuttergrpc.RespSiteStats
	resp, err := toecutter.GetSiteStats(db)
	if err != nil {
		return respSiteStats, err
	}
	respSiteStats = &toecuttergrpc.RespSiteStats{
		Name:     resp.Message.Name,
		Hashrate: resp.Message.Hashrate,
		Hashes:   resp.Message.Hashes,
		Reward:   resp.Message.Reward,
	}
	return respSiteStats, nil
}

func (s server) Payout1MHash(ctx context.Context, in *toecuttergrpc.Empty) (*toecuttergrpc.RespPayout1MHash, error) {
	var respPayout1MHash *toecuttergrpc.RespPayout1MHash
	resp, err := toecutter.GetPayout1MHash(db)
	if err != nil {
		return respPayout1MHash, err
	}
	respPayout1MHash = &toecuttergrpc.RespPayout1MHash{
		Result: resp.Message.Result,
		Status: resp.Status,
	}
	return respPayout1MHash, nil
}

func main() {
	flag.Parse()
	var err error
	var conf toecutter.Config
	conf, err = toecutter.DecodeConfigToml(*configpath)
	if err != nil {
		log.Fatalf("failed read config: %v", err)
	}
	var lis net.Listener

	db, err = leveldb.OpenFile(*dbDir, nil)
	defer db.Close()
	lis, err = net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("cron")
	c := cron.New()
	c.AddFunc(conf.Main.Schedule, func() { Jobs(conf, db, *modeDebug) })
	c.Start()
	defer c.Stop()
	s := grpc.NewServer()
	log.Println("server")
	toecuttergrpc.RegisterCoinimpServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
