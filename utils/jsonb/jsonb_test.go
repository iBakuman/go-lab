package jsonb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonBuilder(t *testing.T) {
	cases := []struct {
		name     string
		input    Component
		expected string
	}{
		{
			name:     "empty object",
			input:    O(),
			expected: "{}",
		},
		{
			name: "object with fields",
			input: O(
				F("name", "john"),
				F("age", 30),
			),
			expected: `{"age":30,"name":"john"}`,
		},
		{
			name: "nested object",
			input: O(
				F("name", "john"),
				F("address", O(
					F("city", "New York"),
					F("zip", 10001),
				)),
			),
			expected: `{"address":{"city":"New York","zip":10001},"name":"john"}`,
		},
		{
			name: "nested array",
			input: O(
				F("name", "john"),
				F("phones", A("123-456-7890", "098-765-4321")),
			),
			expected: `{"name":"john","phones":["123-456-7890","098-765-4321"]}`,
		},
		{
			name: "nested array and object",
			input: O(
				F("name", "john"),
				F("phones", A(
					O(F("type", "home"), F("number", "123-456-7890")),
					O(F("type", "work"), F("number", "098-765-4321")),
				)),
			),
			expected: `{"name":"john","phones":[{"number":"123-456-7890","type":"home"},{"number":"098-765-4321","type":"work"}]}`,
		},
		{
			name:     "empty array",
			input:    A(),
			expected: "[]",
		},
		{
			name:     "array with values",
			input:    A("john", 30),
			expected: `["john",30]`,
		},
		{
			name: "array with nested objects",
			input: A(
				O(F("name", "john"), F("age", 30)),
				O(F("name", "jane"), F("age", 25)),
			),
			expected: `[{"age":30,"name":"john"},{"age":25,"name":"jane"}]`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, err := c.input.Marshal(context.Background())
			assert.NoError(t, err)
			assert.JSONEq(t, c.expected, string(data))
		})
	}
}
