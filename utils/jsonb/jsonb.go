package jsonb

import (
	"context"
	"encoding/json"
)

type Component interface {
	Marshal(ctx context.Context) ([]byte, error)
}

type Object struct {
	Fields []*Field
}

type Field struct {
	Name  string
	Value interface{}
}

type Array struct {
	Values []interface{}
}

func F(name string, value interface{}) *Field {
	return &Field{
		Name:  name,
		Value: value,
	}
}

func O(fs ...*Field) *Object {
	return &Object{
		Fields: fs,
	}
}

func A(vs ...interface{}) *Array {
	return &Array{
		Values: vs,
	}
}

func (o *Object) Marshal(ctx context.Context) ([]byte, error) {
	data := make(map[string]interface{})
	for _, field := range o.Fields {
		switch v := field.Value.(type) {
		case *Object:
			marshaled, err := v.Marshal(ctx)
			if err != nil {
				return nil, err
			}
			var nestedData map[string]interface{}
			err = json.Unmarshal(marshaled, &nestedData)
			if err != nil {
				return nil, err
			}
			data[field.Name] = nestedData
		case *Array:
			marshaled, err := v.Marshal(ctx)
			if err != nil {
				return nil, err
			}
			var nestedData []interface{}
			err = json.Unmarshal(marshaled, &nestedData)
			if err != nil {
				return nil, err
			}
			data[field.Name] = nestedData
		default:
			data[field.Name] = field.Value
		}
	}
	return json.Marshal(data)
}

func (a *Array) Marshal(ctx context.Context) ([]byte, error) {
	data := make([]interface{}, len(a.Values))
	for i, value := range a.Values {
		switch v := value.(type) {
		case *Object:
			marshaled, err := v.Marshal(ctx)
			if err != nil {
				return nil, err
			}
			var nestedData map[string]interface{}
			err = json.Unmarshal(marshaled, &nestedData)
			if err != nil {
				return nil, err
			}
			data[i] = nestedData
		case *Array:
			marshaled, err := v.Marshal(ctx)
			if err != nil {
				return nil, err
			}
			var nestedData []interface{}
			err = json.Unmarshal(marshaled, &nestedData)
			if err != nil {
				return nil, err
			}
			data[i] = nestedData
		default:
			data[i] = value
		}
	}
	return json.Marshal(data)
}
