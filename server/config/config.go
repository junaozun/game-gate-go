package config

type GatewayConfig struct {
	Logic  LogicConfig  `yaml:"logic"`
	Client ClientConfig `yaml:"client"`
	Etcd   Etcd         `yaml:"etcd"`
}

type Etcd struct {
	ServerAddress string `yaml:"server_address"`
	TTL           int    `yaml:"ttl"`
	Key           string `yaml:"key"`
}

type LogicConfig struct {
	Ip              string `yaml:"ip"`
	Thread          int    `yaml:"thread"`
	MaxReadMsgSize  int    `yaml:"max_read_msg_size"`
	MaxWriteMsgSize int    `yaml:"max_write_msg_size"`
}

type ClientConfig struct {
	IP                 string `yaml:"ip"`
	Thread             int    `yaml:"thread"`
	MaxReadMsgSize     int    `yaml:"max_read_msg_size"`
	MaxWriteMsgSize    int    `yaml:"max_write_msg_size"`
	WriteCompressLimit int    `yaml:"write_compress_limit"`
	ReadTimeout        int    `yaml:"read_timeout"`
}
