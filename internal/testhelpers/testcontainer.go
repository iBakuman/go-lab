package testhelpers

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	DBName     = "test"
	DBUser     = "test"
	DBPassword = "test"
)

var (
	PostgresImage = "postgres:16-alpine"
)

func SetupPostgresContainer(t *testing.T) (func(), string) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        PostgresImage,
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       DBName,
			"POSTGRES_USER":     DBUser,
			"POSTGRES_PASSWORD": DBPassword,
		},
		WaitingFor: wait.
			ForLog("database system is ready to accept connections").
			WithStartupTimeout(10 * time.Second),
	}
	pgContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	host, err := pgContainer.Host(ctx)
	require.NoError(t, err)
	port, err := pgContainer.MappedPort(ctx, "5432")
	require.NoError(t, err)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DBUser, DBPassword, host, port.Port(), DBName)
	teardown := func() {
		require.NoError(t, pgContainer.Terminate(ctx))
	}
	return teardown, dsn
}
