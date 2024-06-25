package gorm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go-lab/internal/testhelpers"
	"gorm.io/gorm/logger"
)

type NameProvider interface {
	GetName() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) GetName() string {
	return p.Name
}

type Animal struct {
	Name string
	Sex  int
}

func (a *Animal) GetName() string {
	return a.Name
}

func TestCreate(t *testing.T) {
	conn, f := testhelpers.SetupPostgresWithGorm(t)
	conn.Logger = conn.Logger.LogMode(logger.Info)
	t.Cleanup(func() {
		require.NoError(t, f())
	})
	require.NoError(t, conn.AutoMigrate(&Person{}, &Animal{}))
	t.Run("test create", func(t *testing.T) {
		var a NameProvider
		a = &Person{Name: "John", Age: 30}
		require.NoError(t, conn.Create(a).Error)
		var actual Person
		require.NoError(t, conn.First(&actual, "name = ?", "John").Error)
		require.Equal(t, a.GetName(), actual.GetName())
		require.Equal(t, a.(*Person).Age, actual.Age)

		var b NameProvider
		b = &Animal{Name: "Dog", Sex: 1}
		require.NoError(t, conn.Create(b).Error)
		var actual2 Animal
		require.NoError(t, conn.First(&actual2, "name = ?", "Dog").Error)
		require.Equal(t, b.GetName(), actual2.GetName())
		require.Equal(t, b.(*Animal).Sex, actual2.Sex)
	})

	t.Run("test find", func(t *testing.T) {
		var a NameProvider
		a = &Person{Name: "John", Age: 30}
		require.NoError(t, conn.Create(a).Error)
		var actual NameProvider
		require.Panics(t, func() {
			conn.First(&actual, "name = ?", "John")
		})
		actual = &Person{}
		require.NoError(t, conn.First(actual, "name = ?", "John").Error)
		actual = Person{}
		require.Panics(t, func() {
			conn.First(&actual, "name = ?", "John")
		})
	})
}
