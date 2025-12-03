package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/jayraj/myapp/internal/env"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db: dbconfig{
			dns: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecomdb sslmode=disable"),
		},
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//Database
	conn, err := pgx.Connect(ctx, cfg.db.dns)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("connected to database", "dns", cfg.db.dns)
	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server has failed to start", "error", err)
		os.Exit(1)
	}

}
