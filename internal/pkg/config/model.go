package config

type Config struct {
	LogConfig struct {
		Name  string `mapstructure:"name"`
		Env   string `mapstructure:"env"`
		Level string `mapstructure:"level"`
	} `mapstructure:"log_config"`

	GinConfig struct {
		Port      string `mapstructure:"port"`
		DebugMode bool   `mapstructure:"debug_mode"`
	} `mapstructure:"gin_config"`

	MySQLConfig struct {
		Enable         bool   `mapstructure:"enable"`
		Master         string `mapstructure:"master" env:"MYSQL_MASTER_HOST"`
		Slave          string `mapstructure:"slave" env:"MYSQL_SLAVE_HOST"`
		Username       string `mapstructure:"username"`
		Password       string `mapstructure:"password"`
		Database       string `mapstructure:"database"`
		MaxIdle        int    `mapstructure:"max_idle"`
		MaxOpen        int    `mapstructure:"max_open"`
		ConnMaxLifeSec int    `mapstructure:"conn_max_life_sec"`
		LogMode        bool   `mapstructure:"log_mode"`
	} `mapstructure:"mysql_config"`

	RedisConfig struct {
		Enable         bool   `mapstructure:"enable"`
		Address        string `mapstructure:"address" env:"REDIS_HOST"`
		Password       string `mapstructure:"password"`
		MaxIdle        int    `mapstructure:"max_idle"`
		MaxActive      int    `mapstructure:"max_active"`
		IdleTimeout    int    `mapstructure:"idle_timeout"`
		ConnectTimeout int    `mapstructure:"connect_timeout"`
		ReadTimeout    int    `mapstructure:"read_timeout"`
		WriteTimeout   int    `mapstructure:"write_timeout"`
	} `mapstructure:"redis_config"`

	MongoConfig struct {
		Enable      bool   `mapstructure:"enable"`
		Host        string `mapstructure:"host"`
		Password    string `mapstructure:"passwd"`
		User        string `mapstructure:"user"`
		DB          string `mapstructure:"db"`
		MaxIdleTime int    `mapstructure:"max_idle_time"`
		MaxOpenConn uint64 `mapstructure:"max_open_conn"`
	} `mapstructure:"mongo_config"`

	CacheConfig struct {
		CacheSize int `mapstructure:"cache_size"` // 單位bytes
	} `mapstructure:"cache_config"`

	RateConfig struct {
		Rate           int  `mapstructure:"rate"`             // 请求速率
		Burst          int  `mapstructure:"burst"`            // 请求突发数
		StoreSize      int  `mapstructure:"store_size"`       // 为内存存储定义大小
		UseMemoryStore bool `mapstructure:"use_memory_store"` // 使用内存存储还是Redis存储
	} `mapstructure:"rate_config"`

	WsConfig struct {
		Port                string `mapstructure:"port"`                   // 连接握手超时时间
		MaxConnNum          int    `mapstructure:"max_conn_num"`           // 长连接允许最大链接数
		HandshakeTimeoutSec int    `mapstructure:"handshake_timeout_sec"`  // 连接握手超时时间
		MessageMaxMsgLength int    `mapstructure:"message_max_msg_length"` // 允许消息最大长度
		WriteBufferSize     int    `mapstructure:"write_buffer_size"`      // websocket write buffer, default: 4096, 4kb.
	} `mapstructure:"ws_config"`

	KafkaConfig struct {
		Brokers []string `mapstructure:"brokers"`
		GroupID string   `mapstructure:"group_id"`
		Topic   string   `mapstructure:"topic"`
	} `mapstructure:"kafka_config"`

	QueueConfig struct {
		Mode string `mapstructure:"mode"` // 使用哪種模式 Ex: go channel / redis / kafka
	} `mapstructure:"queue_config"`

	PromConfig struct { // 監控服務設定
		Enable      bool `mapstructure:"enable"`
		EnableDB    bool `mapstructure:"enable_db"`
		EnableRedis bool `mapstructure:"enable_redis"`
	} `mapstructure:"prom_config"`

	ListenerConfig struct { // 監聽器
		Enable  bool `mapstructure:"enable"`
		Msg     int  `mapstructure:"msg"`
		MsgSave int  `mapstructure:"msg_save"`
	} `mapstructure:"listener_config"`
}
