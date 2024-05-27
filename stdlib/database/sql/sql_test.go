package sql

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	open, err := sql.Open("mysql", "user:password@/dbname")
	require.NoError(t, err)
	open.SetConnMaxIdleTime(30 * time.Minute)
}
