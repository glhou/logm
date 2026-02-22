package core

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type App struct {
	Config *Config
	Db     *pgxpool.Pool
	Logger *zap.Logger
}

func GetApp() *App {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
	logger.Info("Logger setup")
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Config setup")
	logger.Debug("Database url", zap.String("db_url", cfg.DatabaseUrl))
	db, err := newDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Db setup")
	return &App{
		Config: cfg,
		Db:     db,
		Logger: logger,
	}
}
