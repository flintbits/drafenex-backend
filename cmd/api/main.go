package main

import (
	"log"

	"github.com/flintbits/drafenex-backend/internal/config"
	"github.com/flintbits/drafenex-backend/internal/container"
	"github.com/flintbits/drafenex-backend/internal/database"
	"github.com/flintbits/drafenex-backend/internal/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
		return
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close() //runs when function returns

	serviceContainer := container.New(pool, cfg)

	router := http.SetupRouter(cfg, serviceContainer)

	err = router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}
