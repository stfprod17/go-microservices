package suite

import (
	"context"
	"net"
	"strconv"
	"testing"

	ssov1 "github.com/stfprod17/protos/gen/go/sso"
	"github.com/stfprod17/sso/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T
	cfg        *config.Config
	AuthClient ssov1.AuthClient
}

// Config возвращает конфигурацию Suite
func (s *Suite) Config() *config.Config {
	return s.cfg
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadByPath("../config/local.yaml")

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.DialContext(
		context.Background(),
		grpcAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T:          t,
		cfg:        cfg,
		AuthClient: ssov1.NewAuthClient(cc),
	}
}

func grpcAddress(cfg *config.Config) string {
	grpcHost := "localhost" // или возьми из cfg, если есть
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}
