package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("dev.config.yml")
	viper.ReadInConfig()
	dbURL := viper.GetString("url_db")
	dbOptions, err := pg.ParseURL(dbURL)
	if err != nil {
		log.Fatal("Failed to parse database URL:", err)
	}

	var seedFiles []string

	// Connect to database
	db := pg.Connect(dbOptions)

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Please provide an argument: up or down")
	}
	if args[0] == "up" {
		seedFiles = []string{
			"migrations/seed/000001_seed_device_types.up.sql",
			"migrations/seed/000002_seed_iot_devices.up.sql",
			"migrations/seed/000003_seed_iot_device_history.up.sql",
			"migrations/seed/000004_seed_sensor_data.up.sql",
		}
	}
	if args[0] == "down" {
		seedFiles = []string{
			"migrations/seed/000001_seed_device_types.down.sql",
			"migrations/seed/000002_seed_iot_devices.down.sql",
			"migrations/seed/000003_seed_iot_device_history.down.sql",
			"migrations/seed/000004_seed_sensor_data.down.sql",
		}
	}

	for _, file := range seedFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("Failed to read file:", err)
		}
		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatal("Failed to execute file:", err)
		}
	}
	fmt.Println("Seed data inserted successfully.")
	defer db.Close()
}
