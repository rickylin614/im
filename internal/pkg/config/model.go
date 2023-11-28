package config

type Config struct {
	LogConfig   LogConfig   `mapstructure:"log_config"`
	GinConfig   GinConfig   `mapstructure:"gin_config"`
	MySQLConfig MySQLConfig `mapstructure:"mysql_config"`
	RedisConfig RedisConfig `mapstructure:"redis_config"`
	MongoConfig MongoConfig `mapstructure:"mongo_config"`
	CacheConfig CacheConfig `mapstructure:"cache_config"`
	RateConfig  RateConfig  `mapstructure:"rate_config"`
}

type LogConfig struct {
	Name  string `mapstructure:"name"`
	Env   string `mapstructure:"env"`
	Level string `mapstructure:"level"`
}

type GinConfig struct {
	Port      string `mapstructure:"port"`
	DebugMode bool   `mapstructure:"debug_mode"`
}

type MySQLConfig struct {
	Master         string `mapstructure:"master" env:"MYSQL_MASTER_HOST"`
	Slave          string `mapstructure:"slave" env:"MYSQL_SLAVE_HOST"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	Database       string `mapstructure:"database"`
	MaxIdle        int    `mapstructure:"max_idle"`
	MaxOpen        int    `mapstructure:"max_open"`
	ConnMaxLifeSec int    `mapstructure:"conn_max_life_sec"`
	LogMode        bool   `mapstructure:"log_mode"`
}

type RedisConfig struct {
	Address        string `mapstructure:"address" env:"REDIS_HOST"`
	Password       string `mapstructure:"password"`
	MaxIdle        int    `mapstructure:"max_idle"`
	MaxActive      int    `mapstructure:"max_active"`
	IdleTimeout    int    `mapstructure:"idle_timeout"`
	ConnectTimeout int    `mapstructure:"connect_timeout"`
	ReadTimeout    int    `mapstructure:"read_timeout"`
	WriteTimeout   int    `mapstructure:"write_timeout"`
}

type MongoConfig struct {
	Host        string `mapstructure:"host"`
	Password    string `mapstructure:"passwd"`
	User        string `mapstructure:"user"`
	DB          string `mapstructure:"db"`
	MaxIdleTime int    `mapstructure:"max_idle_time"`
	MaxOpenConn uint64 `mapstructure:"max_open_conn"`
}

type CacheConfig struct {
	CacheSize int `mapstructure:"cache_size"` // 單位bytes
}

type RateConfig struct {
	Rate           int  `mapstructure:"rate"`             // 请求速率
	Burst          int  `mapstructure:"burst"`            // 请求突发数
	StoreSize      int  `mapstructure:"store_size"`       // 为内存存储定义大小
	UseMemoryStore bool `mapstructure:"use_memory_store"` // 使用内存存储还是Redis存储
}
