package toecutter

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

type Coin int

const (
	Mintme Coin = iota
)

func (c Coin) String() string {
	return [...]string{"mintme"}[c]
}

type RespHashes struct {
	Message uint64 `json:"message"`
	Status  string `json:"status"`
}

func Hashes(conf Config, coin Coin, sitekey string) (RespHashes, error) {
	var respHashes RespHashes
	siteParam := ""
	if len(strings.TrimSpace(sitekey)) != 0 {
		siteParam = fmt.Sprintf("&site-key=%s", sitekey)
	}
	cmds := []string{conf.Main.Curl,
		"-L",
		fmt.Sprintf("%shashes?currency=%s%s", conf.Main.URL, coin.String(), siteParam),
		"-H",
		fmt.Sprintf("X-API-ID:%s", conf.Main.APIPublic),
		"-H",
		fmt.Sprintf("X-API-KEY:%s", conf.Main.APIPrivate),
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return respHashes, err
	}
	if err := json.Unmarshal(out, &respHashes); err != nil {
		return respHashes, err
	}
	return respHashes, nil
}

func SetHashes(db *leveldb.DB, respHashes RespHashes) error {
	obj, err := json.Marshal(respHashes)
	if err != nil {
		return err
	}
	if err := db.Put([]byte("hashes"), obj, nil); err != nil {
		return err
	}
	return nil
}

func GetHashes(db *leveldb.DB) (RespHashes, error) {
	var respHashes RespHashes
	data, err := db.Get([]byte(fmt.Sprintf("hashes")), nil)
	if err != nil {
		return respHashes, err
	}
	if err := json.Unmarshal(data, &respHashes); err != nil {
		return respHashes, err
	}
	return respHashes, nil
}

type RespReward struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func Reward(conf Config, coin Coin, sitekey string) (RespReward, error) {
	var respReward RespReward
	siteParam := ""
	if len(strings.TrimSpace(sitekey)) != 0 {
		siteParam = fmt.Sprintf("&site-key=%s", sitekey)
	}
	cmds := []string{conf.Main.Curl,
		"-L",
		fmt.Sprintf("%sreward?currency=%s%s", conf.Main.URL, coin.String(), siteParam),
		"-H",
		fmt.Sprintf("X-API-ID:%s", conf.Main.APIPublic),
		"-H",
		fmt.Sprintf("X-API-KEY:%s", conf.Main.APIPrivate),
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return respReward, err
	}
	if err := json.Unmarshal(out, &respReward); err != nil {
		return respReward, err
	}
	return respReward, nil
}

func SetReward(db *leveldb.DB, respReward RespReward) error {
	obj, err := json.Marshal(respReward)
	if err != nil {
		return err
	}
	if err := db.Put([]byte("reward"), obj, nil); err != nil {
		return err
	}
	return nil
}

func GetReward(db *leveldb.DB) (RespReward, error) {
	var respReward RespReward
	data, err := db.Get([]byte(fmt.Sprintf("reward")), nil)
	if err != nil {
		return respReward, err
	}
	if err := json.Unmarshal(data, &respReward); err != nil {
		return respReward, err
	}
	return respReward, nil
}

type Stats struct {
	Hashrate      float32 `json:"hashrate"`
	Hashes        uint64  `json:"hashes"`
	Reward        string  `json:"reward"`
	RewardPending string  `json:"reward-pending"`
}

type RespAccountStats struct {
	Message Stats  `json:"message"`
	Status  string `json:"status"`
}

func AccountStats(conf Config, coin Coin) (RespAccountStats, error) {
	var respAccountStats RespAccountStats
	cmds := []string{conf.Main.Curl,
		"-L",
		fmt.Sprintf("%saccount/stats?currency=%s", conf.Main.URL, coin.String()),
		"-H",
		fmt.Sprintf("X-API-ID:%s", conf.Main.APIPublic),
		"-H",
		fmt.Sprintf("X-API-KEY:%s", conf.Main.APIPrivate),
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return respAccountStats, err
	}
	if err := json.Unmarshal(out, &respAccountStats); err != nil {
		return respAccountStats, err
	}
	return respAccountStats, nil
}

func SetAccountStats(db *leveldb.DB, respAccountStats RespAccountStats) error {
	obj, err := json.Marshal(respAccountStats)
	if err != nil {
		return err
	}
	if err := db.Put([]byte("accountstats"), obj, nil); err != nil {
		return err
	}
	return nil
}

func GetAccountStats(db *leveldb.DB) (RespAccountStats, error) {
	var respAccountStats RespAccountStats
	data, err := db.Get([]byte(fmt.Sprintf("accountstats")), nil)
	if err != nil {
		return respAccountStats, err
	}
	if err := json.Unmarshal(data, &respAccountStats); err != nil {
		return respAccountStats, err
	}
	return respAccountStats, nil
}

type Site struct {
	Name    string `json:"name"`
	SiteKey string `json:"site-key"`
}

type RespAccountSite struct {
	Message []Site `json:"message"`
	Status  string `json:"status"`
}

func AccountSite(conf Config, coin Coin) (RespAccountSite, error) {
	var respAccountSite RespAccountSite
	cmds := []string{conf.Main.Curl,
		"-L",
		fmt.Sprintf("%saccount/sites?currency=%s", conf.Main.URL, coin.String()),
		"-H",
		fmt.Sprintf("X-API-ID:%s", conf.Main.APIPublic),
		"-H",
		fmt.Sprintf("X-API-KEY:%s", conf.Main.APIPrivate),
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return respAccountSite, err
	}
	if err := json.Unmarshal(out, &respAccountSite); err != nil {
		return respAccountSite, err
	}
	return respAccountSite, nil
}

func SetAccountSite(db *leveldb.DB, respAccountSite RespAccountSite) error {
	obj, err := json.Marshal(respAccountSite)
	if err != nil {
		return err
	}
	if err := db.Put([]byte("accountsite"), obj, nil); err != nil {
		return err
	}
	return nil
}

func GetAccountSite(db *leveldb.DB) (RespAccountSite, error) {
	var respAccountSite RespAccountSite
	data, err := db.Get([]byte(fmt.Sprintf("accountsite")), nil)
	if err != nil {
		return respAccountSite, err
	}
	if err := json.Unmarshal(data, &respAccountSite); err != nil {
		return respAccountSite, err
	}
	return respAccountSite, nil
}

type Payments struct {
	Amount        string `json:"amount"`
	Fee           string `json:"fee"`
	Status        string `json:"status"`
	WalletAddress string `json:"wallet_address"`
	TxID          string `json:"tx_id"`
	Timestamp     uint64 `json:"timestamp"`
}

type RespAccountPayments struct {
	Message []Payments `json:"message"`
	Status  string     `json:"status"`
}

func AccountPayments(conf Config, coin Coin) (RespAccountPayments, error) {
	var respAccountPayments RespAccountPayments
	cmds := []string{conf.Main.Curl,
		"-L",
		fmt.Sprintf("%saccount/payments?currency=%s", conf.Main.URL, coin.String()),
		"-H",
		fmt.Sprintf("X-API-ID:%s", conf.Main.APIPublic),
		"-H",
		fmt.Sprintf("X-API-KEY:%s", conf.Main.APIPrivate),
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return respAccountPayments, err
	}
	if err := json.Unmarshal(out, &respAccountPayments); err != nil {
		return respAccountPayments, err
	}
	return respAccountPayments, nil
}

func SetAccountPayments(db *leveldb.DB, respAccountPayments RespAccountPayments) error {
	obj, err := json.Marshal(respAccountPayments)
	if err != nil {
		return err
	}
	if err := db.Put([]byte("accountpayments"), obj, nil); err != nil {
		return err
	}
	return nil
}

func GetAccountPayments(db *leveldb.DB) (RespAccountPayments, error) {
	var respAccountPayments RespAccountPayments
	data, err := db.Get([]byte(fmt.Sprintf("accountpayments")), nil)
	if err != nil {
		return respAccountPayments, err
	}
	if err := json.Unmarshal(data, &respAccountPayments); err != nil {
		return respAccountPayments, err
	}
	return respAccountPayments, nil
}

type SiteData struct {
	Name     string  `json:"name"`
	Hashrate float32 `json:"hashrate"`
	Hashes   uint64  `json:"hashes"`
	Reward   string  `json:"reward"`
}

type RespSiteStats struct {
	Message SiteData `json:"message"`
	Status  string   `json:"status"`
}

func SiteStats(conf Config) (RespSiteStats, error) {
	var respSiteStats RespSiteStats
	cmds := []string{conf.Main.Curl,
		"-L",
		fmt.Sprintf("%ssite?site-key=%s", conf.Main.URL, conf.Main.SiteKey),
		"-H",
		fmt.Sprintf("X-API-ID:%s", conf.Main.APIPublic),
		"-H",
		fmt.Sprintf("X-API-KEY:%s", conf.Main.APIPrivate),
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return respSiteStats, err
	}
	if err := json.Unmarshal(out, &respSiteStats); err != nil {
		return respSiteStats, err
	}
	return respSiteStats, nil
}

func SetSiteStats(db *leveldb.DB, respSiteStats RespSiteStats) error {
	obj, err := json.Marshal(respSiteStats)
	if err != nil {
		return err
	}
	if err := db.Put([]byte("sitestats"), obj, nil); err != nil {
		return err
	}
	return nil
}

func GetSiteStats(db *leveldb.DB) (RespSiteStats, error) {
	var respSiteStats RespSiteStats
	data, err := db.Get([]byte(fmt.Sprintf("sitestats")), nil)
	if err != nil {
		return respSiteStats, err
	}
	if err := json.Unmarshal(data, &respSiteStats); err != nil {
		return respSiteStats, err
	}
	return respSiteStats, nil
}

type Payout struct {
	Result string `json:"result"`
}

type RespPayout struct {
	Message Payout `json:"message"`
	Status  string `json:"status"`
}

func Payout1MHash(conf Config, coin Coin) (RespPayout, error) {
	var respPayout RespPayout
	cmds := []string{conf.Main.Curl,
		"-L",
		fmt.Sprintf("%spayout/1mhash?currency=%s", conf.Main.URL, coin.String()),
		"-H",
		fmt.Sprintf("X-API-ID:%s", conf.Main.APIPublic),
		"-H",
		fmt.Sprintf("X-API-KEY:%s", conf.Main.APIPrivate),
	}
	cmd := exec.Command(cmds[0], cmds[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return respPayout, err
	}
	if err := json.Unmarshal(out, &respPayout); err != nil {
		return respPayout, err
	}
	return respPayout, nil
}

func SetPayout1MHash(db *leveldb.DB, respPayout RespPayout) error {
	obj, err := json.Marshal(respPayout)
	if err != nil {
		return err
	}
	if err := db.Put([]byte("payout1mhash"), obj, nil); err != nil {
		return err
	}
	return nil
}

func GetPayout1MHash(db *leveldb.DB) (RespPayout, error) {
	var respPayout RespPayout
	data, err := db.Get([]byte(fmt.Sprintf("payout1mhash")), nil)
	if err != nil {
		return respPayout, err
	}
	if err := json.Unmarshal(data, &respPayout); err != nil {
		return respPayout, err
	}
	return respPayout, nil
}
