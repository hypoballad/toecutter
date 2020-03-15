package toecutter

import "github.com/BurntSushi/toml"

type Config struct {
	Main MainConfig `toml:"main"`
}

// MainConfig ...
type MainConfig struct {
	Curl       string `toml:"curl"`
	URL        string `toml:"url"`
	SiteKey    string `toml:"site-key"`
	APIPublic  string `toml:"api-public"`
	APIPrivate string `toml:"api-private"`
	Schedule   string `toml:"schedule"`
	Sleep      int    `toml:"sleep"`
}

// DecodeConfigToml ...
func DecodeConfigToml(tomlfile string) (Config, error) {
	var config Config
	_, err := toml.DecodeFile(tomlfile, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
