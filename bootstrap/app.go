package bootstrap

import (
	"device-service/infrastructure/repo"

	"github.com/anhvanhoa/service-core/bootstrap/db"
	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/service-core/domain/mq"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	Env    *Env
	DB     *pg.DB
	Log    *log.LogGRPCImpl
	Repo   repo.Repositories
	Helper utils.Helper
	MQ     mq.MQI
}

func App() *Application {
	env := Env{}
	NewEnv(&env)
	logConfig := log.NewConfig()
	log := log.InitLogGRPC(logConfig, zapcore.DebugLevel, env.IsProduction())
	db := db.NewPostgresDB(db.ConfigDB{
		URL:  env.UrlDb,
		Mode: env.NodeEnv,
	})

	configMqtt := mq.MQConfig{
		URL:                  env.UrlMqtt,
		Username:             env.UsernameMqtt,
		Password:             env.PasswordMqtt,
		ClientID:             env.ClientIdMqtt,
		CleanSession:         true,
		KeepAlive:            60,
		PingTimeout:          10,
		AutoReconnect:        true,
		ConnectTimeout:       30,
		MaxReconnectInterval: 10,
		ConnectRetryInterval: 1,
	}

	err := configMqtt.Validate()
	if err != nil {
		log.Fatal("Failed to validate MQTT config: " + err.Error())
	}

	mqtt := mq.NewMQ(configMqtt)
	token := mqtt.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatal("Failed to connect to MQTT: " + token.Error().Error())
	}

	helper := utils.NewHelper()
	repo := repo.InitRepositories(db, helper)

	return &Application{
		Env:    &env,
		DB:     db,
		Log:    log,
		Repo:   repo,
		Helper: helper,
		MQ:     mqtt,
	}
}
