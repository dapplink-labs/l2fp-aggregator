package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Config struct {
	Node                   NodeConfig    `yaml:"node"`
	Manager                ManagerConfig `yaml:"manager"`
	WaitPeersFullConnected bool          `yaml:"wait_peers_full_connected"`
	StartingHeight         int64         `yaml:"starting_height"`
	BlockStep              uint64        `yaml:"block_step"`
	BabylonRpc             string        `yaml:"babylon_rpc"`
}

type NodeConfig struct {
	LevelDbFolder    string        `yaml:"level_db_folder"`
	KeyPath          string        `yaml:"key_path"`
	WsAddr           string        `yaml:"ws_addr"`
	SignTimeout      time.Duration `yaml:"sign_timeout"`
	WaitScanInterval time.Duration `yaml:"wait_scan_interval"`
}

type ManagerConfig struct {
	LevelDbFolder string        `yaml:"level_db_folder"`
	SdkRpc        string        `yaml:"sdk_rpc"`
	WsAddr        string        `yaml:"ws_addr"`
	SignTimeout   time.Duration `yaml:"sign_timeout"`
	NodeMembers   string        `yaml:"node_members"`
}

func DefaultConfiguration() *Config {
	return &Config{
		WaitPeersFullConnected: true,
	}
}

func NewConfig(path string) (*Config, error) {
	var config = new(Config)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return DefaultConfiguration(), nil
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
