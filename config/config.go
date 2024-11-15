package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Server    ServerConfig
	Pages     PagesConfig
	Log       LogConfig
	CORS      CORSConfig
	Auth      AuthConfig
	Blacklist BlacklistConfig
	Whitelist WhitelistConfig
	RateLimit RateLimitConfig
}

type ServerConfig struct {
	Port      int    `toml:"port"`
	Host      string `toml:"host"`
	SizeLimit int    `toml:"sizeLimit"`
	EnableH2C bool   `toml:"enableH2C"`
}

type PagesConfig struct {
	Enabled   bool   `toml:"enabled"`
	StaticDir string `toml:"staticDir"`
}

type LogConfig struct {
	LogFilePath string `toml:"logFilePath"`
	MaxLogSize  int    `toml:"maxLogSize"`
}

type CORSConfig struct {
	Enabled bool `toml:"enabled"`
}

type AuthConfig struct {
	Enabled    bool   `toml:"enabled"`
	AuthMethod string `toml:"authMethod"`
	AuthToken  string `toml:"authToken"`
}

type BlacklistConfig struct {
	Enabled       bool   `toml:"enabled"`
	BlacklistFile string `toml:"blacklistFile"`
}

type WhitelistConfig struct {
	Enabled       bool   `toml:"enabled"`
	WhitelistFile string `toml:"whitelistFile"`
}

type RateLimitConfig struct {
	Enabled       bool `toml:"enabled"`
	RatePerMinute int  `toml:"ratePerMinute"`
	Burst         int  `toml:"burst"`
}

// LoadConfig 从 TOML 配置文件加载配置
func LoadConfig(filePath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
