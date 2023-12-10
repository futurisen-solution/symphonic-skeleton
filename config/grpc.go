package config

import (
	"github.com/fwidjaya20/symphonic/facades"
)

func init() {
	config := facades.Config()

	config.Add("grpc", map[string]any{
		"self": map[string]any{
			"host": config.Get("GRPC_SELF_HOST", "localhost"),
			"port": config.Get("GRPC_SELF_PORT", "9000"),
		},
	})
}
