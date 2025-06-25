package config

// DefaultValues 集中管理所有預設值
var DefaultValues = map[string]interface{}{
	// Server 預設值
	"server.host": "0.0.0.0",
	"server.port": 8080,
	"server.mode": "development",

	// Database 預設值
	"database.type":              "postgres",
	"database.host":              "localhost",
	"database.port":              5432,
	"database.sslmode":           "disable",
	"database.max_open_conns":    25,
	"database.max_idle_conns":    10,
	"database.conn_max_lifetime": "5m",
	"database.timezone":          "UTC",

	// Auth 預設值
	"auth.token_expiry":   "24h",
	"auth.refresh_expiry": "168h", // 7 days

	// Media 預設值
	"media.storage_type":    "local",
	"media.max_upload_size": 100 * 1024 * 1024, // 100MB
	"media.use_ssl":         false,

	// Chat 預設值
	"chat.max_message_length": 1000,
	"chat.history_limit":      100,

	// Room 預設值
	"room.max_participants": 50,
	"room.default_ttl":      "24h",

	// Logger 預設值
	"logger.level":  "info",
	"logger.format": "json",
	"logger.output": "stdout",

	// MessageQueue 預設值
	"message_queue.type":    "nats",
	"message_queue.servers": []string{"nats://localhost:4222"},

	// NATS 預設值
	"message_queue.nats.max_reconnects": 5,
	"message_queue.nats.reconnect_wait": "2s",
	"message_queue.nats.ping_interval":  "20s",
	"message_queue.nats.max_pings_out":  2,

	// Kafka 預設值
	"message_queue.kafka.session_timeout":    "10s",
	"message_queue.kafka.heartbeat_interval": "3s",
	"message_queue.kafka.retry_backoff":      "2s",
	"message_queue.kafka.required_acks":      1,
	"message_queue.kafka.compression":        "none",

	// Redis 預設值
	"message_queue.redis.db":             0,
	"message_queue.redis.pool_size":      10,
	"message_queue.redis.min_idle_conns": 5,
	"message_queue.redis.dial_timeout":   "5s",
	"message_queue.redis.read_timeout":   "3s",
	"message_queue.redis.write_timeout":  "3s",

	// TLS 預設值
	"message_queue.tls.enabled":     false,
	"message_queue.tls.skip_verify": false,
}
