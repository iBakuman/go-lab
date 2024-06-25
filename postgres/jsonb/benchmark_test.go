package jsonb

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"go-lab/internal/testhelpers"
	"go-lab/utils/jsonb"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	Data json.RawMessage `gorm:"type:jsonb"`
}

const (
	requireJsonPath = "identity.email"
)

func genRandomData(t testing.TB, size int) []*Model {
	var models []*Model
	for i := 0; i < size; i++ {
		randomJson, err := jsonb.GenRandomJson(requireJsonPath, 10)
		require.NoError(t, err)
		data, err := randomJson.Marshal(context.Background())
		require.NoError(t, err)
		models = append(models, &Model{
			Data: data,
		})
	}
	return models
}

func insertRandomData(t testing.TB, conn *gorm.DB, size int) []*Model {
	var preSize int64
	require.NoError(t, conn.Model(&Model{}).Count(&preSize).Error)
	dataset := genRandomData(t, size)
	require.NoError(t, conn.Create(dataset).Error)
	var postSize int64
	require.NoError(t, conn.Model(&Model{}).Count(&postSize).Error)
	require.Equal(t, preSize+int64(size), postSize)
	return dataset
}

func BenchmarkJsonB(b *testing.B) {
	conn, cleanup := testhelpers.SetupPostgresWithGorm(b)
	b.Cleanup(func() {
		require.NoError(b, cleanup())
	})
	require.NoError(b, conn.AutoMigrate(&Model{}))
	dataset := insertRandomData(b, conn, 10000)
	emails := make([]string, 0, 1000)
	for _, m := range dataset {
		email := gjson.GetBytes(m.Data, requireJsonPath).String()
		require.NotEmpty(b, email)
		emails = append(emails, email)
	}
	selectRandomEmails := func() string {
		return emails[b.N%len(emails)]
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		email := selectRandomEmails()
		var result []*Model
		require.NoError(b, conn.Model(&Model{}).Where("data->>identity.email = ?", email).Find(&result).Error)
		require.NotNil(b, result)
		require.Greater(b, len(result), 0)
		for _, r := range result {
			require.NotEmpty(b, r.Data)
			actual := gjson.GetBytes(r.Data, requireJsonPath).String()
			require.Equal(b, email, actual)
		}
	}
}
