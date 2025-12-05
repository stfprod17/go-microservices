package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "github.com/stfprod17/go-microservices/sso/internal/config/services/auth"
	"google.golang.org/grpc"
)

// App представляет gRPC сервер
type App struct {
	log     *slog.Logger
	GRPCSrv *grpc.Server
	port    int
	authSvc *authgrpc.Auth
}

// New создаёт новый gRPC сервер
func New(log *slog.Logger, authService *authgrpc.Auth, port int) *App {
	return &App{
		log:     log,
		GRPCSrv: grpc.NewServer(),
		port:    port,
		authSvc: authService,
	}
}

// MustRun запускает сервер и паникует при ошибке
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run запускает сервер
func (a *App) Run() error {
	const op = "grpcapp.Run"

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info("gRPC server is running", slog.String("addr", listener.Addr().String()))

	// Здесь регистрируй свои gRPC сервисы
	// Например:
	// authpb.RegisterAuthServiceServer(a.GRPCSrv, a.authSvc)

	if err := a.GRPCSrv.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop останавливает сервер
func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", a.port))

	a.GRPCSrv.GracefulStop()
}
