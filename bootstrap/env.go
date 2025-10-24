package bootstrap

import (
	"strings"

	"github.com/anhvanhoa/service-core/bootstrap/config"
	"github.com/anhvanhoa/service-core/domain/grpc_client"
)

type dbCache struct {
	Addr        string `mapstructure:"addr"`
	Db          int    `mapstructure:"db"`
	Password    string `mapstructure:"password"`
	MaxIdle     int    `mapstructure:"max_idle"`
	MaxActive   int    `mapstructure:"max_active"`
	IdleTimeout int    `mapstructure:"idle_timeout"`
	Network     string `mapstructure:"network"`
}

type Env struct {
	NodeEnv               string                    `mapstructure:"node_env"`
	SecretService         string                    `mapstructure:"secret_service"`
	UrlDb                 string                    `mapstructure:"url_db"`
	NameService           string                    `mapstructure:"name_service"`
	PortGrpc              int                       `mapstructure:"port_grpc"`
	HostGprc              string                    `mapstructure:"host_grpc"`
	IntervalCheck         string                    `mapstructure:"interval_check"`
	TimeoutCheck          string                    `mapstructure:"timeout_check"`
	UrlMqtt               string                    `mapstructure:"url_mqtt"`
	UsernameMqtt          string                    `mapstructure:"username_mqtt"`
	PasswordMqtt          string                    `mapstructure:"password_mqtt"`
	ClientIdMqtt          string                    `mapstructure:"client_id_mqtt"`
	TlsMqtt               bool                      `mapstructure:"tls_mqtt"`
	PermissionServiceAddr string                    `mapstructure:"permission_service_addr"`
	GrpcClients           []*grpc_client.ConfigGrpc `mapstructure:"grpc_clients"`
	DbCache               *dbCache                  `mapstructure:"db_cache"`
}

func NewEnv(env any) {
	setting := config.DefaultSettingsConfig()
	if setting.IsProduction() {
		setting.SetPath("/config")
		setting.SetFile("device_service.config")
	} else {
		setting.SetFile("dev.config")
	}
	config.NewConfig(setting, env)
}

func (env *Env) IsProduction() bool {
	return strings.ToLower(env.NodeEnv) == "production"
}
