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
	L2ooContractAddress    string        `yaml:"l2oo_contract_address"`
	MsmContractAddress     string        `yaml:"msm_contract_address"`
	L1ChainID              uint64        `yaml:"l1_chain_id"`
}

type NodeConfig struct {
	DBCfg          DB            `yaml:"db_cfg"`
	KeyPath        string        `yaml:"key_path"`
	RollupRpc      string        `yaml:"rollup_rpc"`
	WsAddr         string        `yaml:"ws_addr"`
	P2PPort        string        `yaml:"p2p_port"`
	BootstrapPeers string        `yaml:"bootstrap_peers"`
	ExternalIP     string        `yaml:"external_ip"`
	SignTimeout    time.Duration `yaml:"sign_timeout"`
	OutputInterval time.Duration `yaml:"output_interval"`
}

type ManagerConfig struct {
	L1EthRpc       string        `yaml:"l1_eth_rpc"`
	RollupRpc      string        `yaml:"rollup_rpc"`
	SdkRpc         string        `yaml:"sdk_rpc"`
	WsAddr         string        `yaml:"ws_addr"`
	DBCfg          DB            `yaml:"db_cfg"`
	SignTimeout    time.Duration `yaml:"sign_timeout"`
	NodeMembers    string        `yaml:"node_members"`
	PollInterval   time.Duration `yaml:"poll_interval"`
	NetworkTimeout time.Duration `yaml:"network_timeout"`
}

type DB struct {
	DbHost     string `yaml:"db_host"`
	DbPort     int    `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
}

func DefaultConfiguration() *Config {

	return &Config{
		WaitPeersFullConnected: true,
		Node: NodeConfig{
			DBCfg: DB{
				DbHost: "127.0.0.1",
				DbPort: 5432,
				DbName: "finality_node",
				DbUser: "postgres",
			},
			P2PPort: "8000",
		},
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
