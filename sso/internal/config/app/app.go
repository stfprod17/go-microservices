package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/stfprod17/go-microservices/sso/internal/config/app/grpc"
	"github.com/stfprod17/go-microservices/sso/internal/config/services/auth"
	"github.com/stfprod17/go-microservices/sso/internal/storage/sqlite"
)

type App struct {
	log      *slog.Logger
	GrpcApp  *grpcapp.App
	Storage  *sqlite.Storage // если нужно хранить
	TokenTTL time.Duration
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApplication := grpcapp.New(log, authService, grpcPort)

	return &App{
		log:      log,
		GrpcApp:  grpcApplication,
		Storage:  storage,
		TokenTTL: tokenTTL,
	}
}

// Метод для запуска gRPC сервера
func (a *App) MustRun() {
	a.GrpcApp.MustRun()
}
