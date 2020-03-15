package toecutter

import (
	"reflect"
	"testing"
)

func TestHashes(t *testing.T) {
	config, err := DecodeConfigToml("toecutter.toml")
	if err != nil {
		t.Error(err)
	}
	type args struct {
		conf    Config
		coin    Coin
		sitekey string
	}
	tests := []struct {
		name    string
		args    args
		want    RespHashes
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Hashes", args: args{conf: config, coin: Mintme, sitekey: config.Main.SiteKey}, want: RespHashes{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hashes(tt.args.conf, tt.args.coin, tt.args.sitekey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hashes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hashes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReward(t *testing.T) {
	config, err := DecodeConfigToml("toecutter.toml")
	if err != nil {
		t.Error(err)
	}
	type args struct {
		conf    Config
		coin    Coin
		sitekey string
	}
	tests := []struct {
		name    string
		args    args
		want    RespReward
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Reward", args: args{conf: config, coin: Mintme, sitekey: config.Main.SiteKey}, want: RespReward{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reward(tt.args.conf, tt.args.coin, tt.args.sitekey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reward() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountStats(t *testing.T) {
	config, err := DecodeConfigToml("toecutter.toml")
	if err != nil {
		t.Error(err)
	}
	type args struct {
		conf Config
		coin Coin
	}
	tests := []struct {
		name    string
		args    args
		want    RespAccountStats
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "AccountsStats", args: args{conf: config, coin: Mintme}, want: RespAccountStats{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AccountStats(tt.args.conf, tt.args.coin)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountSite(t *testing.T) {
	config, err := DecodeConfigToml("toecutter.toml")
	if err != nil {
		t.Error(err)
	}
	type args struct {
		conf Config
		coin Coin
	}
	tests := []struct {
		name    string
		args    args
		want    RespAccountSite
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "AccountsSite", args: args{conf: config, coin: Mintme}, want: RespAccountSite{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AccountSite(tt.args.conf, tt.args.coin)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountSite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountSite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountPayments(t *testing.T) {
	config, err := DecodeConfigToml("toecutter.toml")
	if err != nil {
		t.Error(err)
	}
	type args struct {
		conf Config
		coin Coin
	}
	tests := []struct {
		name    string
		args    args
		want    RespAccountPayments
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "AccountsPayments", args: args{conf: config, coin: Mintme}, want: RespAccountPayments{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AccountPayments(tt.args.conf, tt.args.coin)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountPayments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountPayments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSiteStats(t *testing.T) {
	config, err := DecodeConfigToml("toecutter.toml")
	if err != nil {
		t.Error(err)
	}
	type args struct {
		conf Config
	}
	tests := []struct {
		name    string
		args    args
		want    RespSiteStats
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "SiteStats", args: args{conf: config}, want: RespSiteStats{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SiteStats(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("SiteStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SiteStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPayout1MHash(t *testing.T) {
	config, err := DecodeConfigToml("toecutter.toml")
	if err != nil {
		t.Error(err)
	}
	type args struct {
		conf Config
		coin Coin
	}
	tests := []struct {
		name    string
		args    args
		want    RespPayout
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "AccountsStats", args: args{conf: config, coin: Mintme}, want: RespPayout{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Payout1MHash(tt.args.conf, tt.args.coin)
			if (err != nil) != tt.wantErr {
				t.Errorf("Payout1MHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Payout1MHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
