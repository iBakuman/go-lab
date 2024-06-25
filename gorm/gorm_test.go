package gorm

import (
	"reflect"
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

func TestReflect(t *testing.T) {
	t.Run("test reflect.TypeOf", func(t *testing.T) {
		var a NameProvider
		a = Person{}
		typ := reflect.TypeOf(a)
		require.Equal(t, "Person", typ.Name())

		a = &Person{}
		typ = reflect.TypeOf(a)
		require.Empty(t, typ.Name())
		typ = reflect.Indirect(reflect.ValueOf(a)).Type()
		require.Equal(t, "Person", typ.Name())

		a = &Animal{}
		typ = reflect.TypeOf(a)
		require.Empty(t, typ.Name())
		typ = reflect.Indirect(reflect.ValueOf(a)).Type()
		require.Equal(t, "Animal", typ.Name())
	})
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

		typ1 := reflect.TypeOf(Person{})
		value1 := reflect.New(typ1).Interface()
		reflect.SliceOf(typ1)
		typ := reflect.TypeOf([]*Person{})
		value := reflect.New(typ).Interface()
		require.NoError(t, conn.Model(value1).Find(value).Error)
		allPerson := *value.(*[]*Person)
		for _, person := range allPerson {
			t.Logf("name: %v, age: %v\n", person.Name, person.Age)
		}
	})
}
