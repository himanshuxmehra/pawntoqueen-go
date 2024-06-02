package config

import (
    "log"

    "github.com/jmoiron/sqlx"
    "github.com/lib/pq"
    "github.com/spf13/viper"
)

func Load() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
}

func InitDB() (*sqlx.DB, error) {
    dsn := viper.GetString("database.dsn")
    return sqlx.Connect("postgres", dsn)
}
