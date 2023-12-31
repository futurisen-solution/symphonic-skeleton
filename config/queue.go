package config

import "github.com/fwidjaya20/symphonic/facades"

func init() {
	config := facades.Config()

	config.Add("queue", map[string]any{
		"connections": map[string]any{
			"kafka": map[string]any{
				"host": facades.Config().Get("KAFKA_HOST"),
				"port": facades.Config().Get("KAFKA_PORT"),
			},
			"rabbitmq": map[string]any{
				"protocol": facades.Config().Get("RABBITMQ_PROTOCOL"),
				"username": facades.Config().Get("RABBITMQ_USERNAME"),
				"password": facades.Config().Get("RABBITMQ_PASSWORD"),
				"host":     facades.Config().Get("RABBITMQ_HOST"),
				"port":     facades.Config().Get("RABBITMQ_PORT"),
			},
			"redis": map[string]any{
				"host":     facades.Config().Get("database.connections.redis.host"),
				"port":     facades.Config().Get("database.connections.redis.port"),
				"database": facades.Config().Get("database.connections.redis.database"),
				"password": facades.Config().Get("database.connections.redis.password"),
			},
		},
		"default": "kafka",
	})
}
