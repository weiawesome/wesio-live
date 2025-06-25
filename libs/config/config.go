package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config 定義應用程序的配置結構
type Config struct {
	// Server 配置
	Server ServerConfig `mapstructure:"server" yaml:"server"`

	// Database 配置
	Database DatabaseConfig `mapstructure:"database" yaml:"database"`

	// Auth 配置
	Auth AuthConfig `mapstructure:"auth" yaml:"auth"`

	// Media 配置
	Media MediaConfig `mapstructure:"media" yaml:"media"`

	// Chat 配置
	Chat ChatConfig `mapstructure:"chat" yaml:"chat"`

	// Room 配置
	Room RoomConfig `mapstructure:"room" yaml:"room"`

	// Logger 配置
	Logger LoggerConfig `mapstructure:"logger" yaml:"logger"`

	// MessageQueue 配置
	MessageQueue MessageQueueConfig `mapstructure:"message_queue" yaml:"message_queue"`
}

// ServerConfig 服務器配置
type ServerConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port int    `mapstructure:"port" yaml:"port"`
	Mode string `mapstructure:"mode" yaml:"mode"` // development, production
}

// DatabaseConfig 數據庫配置
type DatabaseConfig struct {
	Type            string `mapstructure:"type" yaml:"type"` // postgres, mysql, sqlite
	Host            string `mapstructure:"host" yaml:"host"`
	Port            int    `mapstructure:"port" yaml:"port"`
	User            string `mapstructure:"user" yaml:"user"`
	Password        string `mapstructure:"password" yaml:"password"`
	DBName          string `mapstructure:"dbname" yaml:"dbname"`
	SSLMode         string `mapstructure:"sslmode" yaml:"sslmode"`
	MaxOpenConns    int    `mapstructure:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
	ConnMaxLifetime string `mapstructure:"conn_max_lifetime" yaml:"conn_max_lifetime"`
	Timezone        string `mapstructure:"timezone" yaml:"timezone"`
	DSN             string `mapstructure:"dsn" yaml:"dsn"` // 直接指定 DSN，優先級最高
}

// AuthConfig 認證配置
type AuthConfig struct {
	JWTSecret     string `mapstructure:"jwt_secret" yaml:"jwt_secret"`
	TokenExpiry   string `mapstructure:"token_expiry" yaml:"token_expiry"`
	RefreshExpiry string `mapstructure:"refresh_expiry" yaml:"refresh_expiry"`
}

// MediaConfig 媒體配置
type MediaConfig struct {
	StorageType   string `mapstructure:"storage_type" yaml:"storage_type"` // minio, s3, local
	Endpoint      string `mapstructure:"endpoint" yaml:"endpoint"`
	AccessKey     string `mapstructure:"access_key" yaml:"access_key"`
	SecretKey     string `mapstructure:"secret_key" yaml:"secret_key"`
	BucketName    string `mapstructure:"bucket_name" yaml:"bucket_name"`
	UseSSL        bool   `mapstructure:"use_ssl" yaml:"use_ssl"`
	MaxUploadSize int64  `mapstructure:"max_upload_size" yaml:"max_upload_size"` // bytes
}

// ChatConfig 聊天配置
type ChatConfig struct {
	MaxMessageLength int `mapstructure:"max_message_length" yaml:"max_message_length"`
	HistoryLimit     int `mapstructure:"history_limit" yaml:"history_limit"`
}

// RoomConfig 房間配置
type RoomConfig struct {
	MaxParticipants int    `mapstructure:"max_participants" yaml:"max_participants"`
	DefaultTTL      string `mapstructure:"default_ttl" yaml:"default_ttl"`
}

// LoggerConfig 日誌配置
type LoggerConfig struct {
	Level  string `mapstructure:"level" yaml:"level"`   // debug, info, warn, error
	Format string `mapstructure:"format" yaml:"format"` // json, console
	Output string `mapstructure:"output" yaml:"output"` // stdout, file
	File   string `mapstructure:"file" yaml:"file"`     // log file path when output is file
}

// MessageQueueConfig 消息隊列配置
type MessageQueueConfig struct {
	Type     string            `mapstructure:"type" yaml:"type"`         // nats, kafka, redis
	Servers  []string          `mapstructure:"servers" yaml:"servers"`   // 服務器地址列表
	Username string            `mapstructure:"username" yaml:"username"` // 用戶名
	Password string            `mapstructure:"password" yaml:"password"` // 密碼
	TLS      TLSConfig         `mapstructure:"tls" yaml:"tls"`           // TLS 配置
	Options  map[string]string `mapstructure:"options" yaml:"options"`   // 額外選項

	// NATS 特定配置
	NATS NATSConfig `mapstructure:"nats" yaml:"nats"`

	// Kafka 特定配置
	Kafka KafkaConfig `mapstructure:"kafka" yaml:"kafka"`

	// Redis 特定配置
	Redis RedisConfig `mapstructure:"redis" yaml:"redis"`
}

// TLSConfig TLS 配置
type TLSConfig struct {
	Enabled    bool   `mapstructure:"enabled" yaml:"enabled"`
	CertFile   string `mapstructure:"cert_file" yaml:"cert_file"`
	KeyFile    string `mapstructure:"key_file" yaml:"key_file"`
	CAFile     string `mapstructure:"ca_file" yaml:"ca_file"`
	SkipVerify bool   `mapstructure:"skip_verify" yaml:"skip_verify"`
}

// NATSConfig NATS 特定配置
type NATSConfig struct {
	ClusterID     string `mapstructure:"cluster_id" yaml:"cluster_id"`
	ClientID      string `mapstructure:"client_id" yaml:"client_id"`
	MaxReconnects int    `mapstructure:"max_reconnects" yaml:"max_reconnects"`
	ReconnectWait string `mapstructure:"reconnect_wait" yaml:"reconnect_wait"`
	PingInterval  string `mapstructure:"ping_interval" yaml:"ping_interval"`
	MaxPingsOut   int    `mapstructure:"max_pings_out" yaml:"max_pings_out"`
}

// KafkaConfig Kafka 特定配置
type KafkaConfig struct {
	GroupID           string `mapstructure:"group_id" yaml:"group_id"`
	SessionTimeout    string `mapstructure:"session_timeout" yaml:"session_timeout"`
	HeartbeatInterval string `mapstructure:"heartbeat_interval" yaml:"heartbeat_interval"`
	RetryBackoff      string `mapstructure:"retry_backoff" yaml:"retry_backoff"`
	RequiredAcks      int    `mapstructure:"required_acks" yaml:"required_acks"`
	Compression       string `mapstructure:"compression" yaml:"compression"` // none, gzip, snappy, lz4, zstd
}

// RedisConfig Redis 特定配置
type RedisConfig struct {
	DB           int    `mapstructure:"db" yaml:"db"`
	PoolSize     int    `mapstructure:"pool_size" yaml:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns" yaml:"min_idle_conns"`
	DialTimeout  string `mapstructure:"dial_timeout" yaml:"dial_timeout"`
	ReadTimeout  string `mapstructure:"read_timeout" yaml:"read_timeout"`
	WriteTimeout string `mapstructure:"write_timeout" yaml:"write_timeout"`
}

var globalConfig *Config

// LoadConfig 載入配置
// configPath: 配置文件路徑，如果為空則使用默認路徑
// envPrefix: 環境變數前綴，默認為 "WESIO"
func LoadConfig(configPath, envPrefix string) (*Config, error) {
	v := viper.New()

	// 設置默認值
	setDefaults(v)

	// 設置環境變數前綴
	if envPrefix == "" {
		envPrefix = "WESIO"
	}
	v.SetEnvPrefix(envPrefix)

	// 設置環境變數
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// 設置配置文件
	if configPath != "" {
		// 使用指定的配置文件路徑
		v.SetConfigFile(configPath)
	} else {
		// 默認配置文件搜索路徑
		v.SetConfigName("config")
		v.SetConfigType("yaml")

		// 添加配置文件搜索路徑
		v.AddConfigPath(".")
		v.AddConfigPath("./config")

		// 檢查工作目錄
		if wd, err := os.Getwd(); err == nil {
			v.AddConfigPath(wd)
			v.AddConfigPath(filepath.Join(wd, "config"))
		}
	}

	// 讀取配置文件
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到，使用默認值和環境變數
			fmt.Printf("Warning: No config file found, using defaults and environment variables\n")
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	} else {
		fmt.Printf("Using config file: %s\n", v.ConfigFileUsed())
	}

	// 解析配置到結構體
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	// 驗證配置
	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	globalConfig = &config
	return &config, nil
}

// GetConfig 獲取全局配置實例
func GetConfig() *Config {
	if globalConfig == nil {
		panic("config not loaded, please call LoadConfig first")
	}
	return globalConfig
}

// setDefaults 設置默認配置值
func setDefaults(v *viper.Viper) {
	for key, value := range DefaultValues {
		v.SetDefault(key, value)
	}
}

// validateConfig 驗證配置的有效性
func validateConfig(config *Config) error {
	// 驗證服務器配置
	if config.Server.Port <= 0 || config.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", config.Server.Port)
	}

	// 驗證數據庫配置
	if config.Database.Port <= 0 || config.Database.Port > 65535 {
		return fmt.Errorf("invalid database port: %d", config.Database.Port)
	}

	// 驗證日誌級別
	validLogLevels := map[string]bool{
		"debug": true, "info": true, "warn": true, "error": true,
	}
	if !validLogLevels[config.Logger.Level] {
		return fmt.Errorf("invalid log level: %s", config.Logger.Level)
	}

	// 驗證日誌格式
	validLogFormats := map[string]bool{
		"json": true, "console": true,
	}
	if !validLogFormats[config.Logger.Format] {
		return fmt.Errorf("invalid log format: %s", config.Logger.Format)
	}

	return nil
}

// GetDatabaseURL 獲取數據庫連接 URL
func (c *Config) GetDatabaseURL() string {
	// 如果有直接指定 DSN，優先使用
	if c.Database.DSN != "" {
		return c.Database.DSN
	}

	// 根據數據庫類型生成對應的 URL
	switch c.Database.Type {
	case "postgres", "postgresql":
		url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			c.Database.User,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
			c.Database.DBName,
			c.Database.SSLMode,
		)
		if c.Database.Timezone != "" {
			url += "&timezone=" + c.Database.Timezone
		}
		return url

	case "mysql":
		url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.Database.User,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
			c.Database.DBName,
		)
		if c.Database.Timezone != "" {
			url += "&time_zone=" + c.Database.Timezone
		}
		return url

	case "sqlite", "sqlite3":
		return c.Database.DBName // SQLite 使用文件路徑

	default:
		// 默認使用 PostgreSQL 格式
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			c.Database.User,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
			c.Database.DBName,
			c.Database.SSLMode,
		)
	}
}

// GetServerAddress 獲取服務器地址
func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

// IsProduction 判斷是否為生產環境
func (c *Config) IsProduction() bool {
	return c.Server.Mode == "production"
}

// IsDevelopment 判斷是否為開發環境
func (c *Config) IsDevelopment() bool {
	return c.Server.Mode == "development"
}

// GetMessageQueueURL 獲取消息隊列連接 URL
func (c *Config) GetMessageQueueURL() []string {
	if len(c.MessageQueue.Servers) > 0 {
		return c.MessageQueue.Servers
	}

	// 根據類型返回默認連接
	switch c.MessageQueue.Type {
	case "nats":
		return []string{"nats://localhost:4222"}
	case "kafka":
		return []string{"localhost:9092"}
	case "redis":
		return []string{"redis://localhost:6379"}
	default:
		return []string{"nats://localhost:4222"}
	}
}

// IsNATS 判斷是否使用 NATS
func (c *Config) IsNATS() bool {
	return c.MessageQueue.Type == "nats"
}

// IsKafka 判斷是否使用 Kafka
func (c *Config) IsKafka() bool {
	return c.MessageQueue.Type == "kafka"
}

// IsRedis 判斷是否使用 Redis 作為消息隊列
func (c *Config) IsRedis() bool {
	return c.MessageQueue.Type == "redis"
}

// GetRedisURL 獲取 Redis 連接 URL (當作為消息隊列時)
func (c *Config) GetRedisURL() string {
	if len(c.MessageQueue.Servers) > 0 {
		return c.MessageQueue.Servers[0]
	}

	// 構建 Redis URL
	url := "redis://"
	if c.MessageQueue.Username != "" {
		url += c.MessageQueue.Username
		if c.MessageQueue.Password != "" {
			url += ":" + c.MessageQueue.Password
		}
		url += "@"
	}
	url += "localhost:6379"
	if c.MessageQueue.Redis.DB != 0 {
		url += fmt.Sprintf("/%d", c.MessageQueue.Redis.DB)
	}

	return url
}
