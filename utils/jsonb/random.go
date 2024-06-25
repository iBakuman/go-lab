package jsonb

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// GenRandomJson recursively creates a JSON structure ensuring a specific path exists.
func GenRandomJson(requiredPath string, depth int) (*Object, error) {
	if depth <= 0 {
		return nil, fmt.Errorf("maximum depth reached")
	}

	// Decompose the path into the next segment and remaining path
	segment, rest, isArray, index := parsePathSegment(requiredPath)

	// Create the root object
	obj := O()

	// If the segment is part of the required path
	if segment != "" {
		var value interface{}
		if isArray {
			// Create array with a placeholder object to fulfill the required path
			array := A()
			nestedObj, err := GenRandomJson(rest, depth-1)
			if err != nil {
				return nil, err
			}

			// Ensure the array has enough elements
			for i := 0; i <= index; i++ {
				if i == index {
					array.Values = append(array.Values, nestedObj)
				} else {
					array.Values = append(array.Values, O()) // Fill with empty objects
				}
			}
			value = array
		} else if rest == "" {
			// If it's the last segment, assign a random value
			value = rand.Intn(100)
		} else {
			// Otherwise, continue with nested object
			nextObj, err := GenRandomJson(rest, depth-1)
			if err != nil {
				return nil, err
			}
			value = nextObj
		}
		obj.Fields = append(obj.Fields, F(segment, value))
	}

	// Optionally add random fields to add complexity
	if rand.Intn(10) > 5 { // Random chance to add extra content
		addRandomFields(obj, depth-1)
	}

	return obj, nil
}

// parsePathSegment splits the next segment from the rest of the path
func parsePathSegment(path string) (segment string, rest string, isArray bool, index int) {
	dotIndex := strings.Index(path, ".")
	bracketIndex := strings.Index(path, "[")

	if dotIndex == -1 && bracketIndex == -1 {
		return path, "", false, -1
	}

	if (bracketIndex != -1 && bracketIndex < dotIndex) || dotIndex == -1 {
		segment = path[:bracketIndex]
		rest = path[bracketIndex+1:]
		isArray = true
		closeBracketIndex := strings.Index(rest, "]")
		index, _ = strconv.Atoi(rest[:closeBracketIndex])
		rest = rest[closeBracketIndex+2:]
	} else {
		segment = path[:dotIndex]
		rest = path[dotIndex+1:]
	}

	return
}

// addRandomFields adds random fields to the given object, respecting the maximum depth.
func addRandomFields(o *Object, depth int) {
	if depth <= 0 {
		return
	}

	count := rand.Intn(3) + 1
	for i := 0; i < count; i++ {
		name := fmt.Sprintf("extraField%d", i)
		if rand.Intn(2) == 0 {
			o.Fields = append(o.Fields, F(name, rand.Intn(100)))
		} else {
			nested := O()
			addRandomFields(nested, depth-1)
			o.Fields = append(o.Fields, F(name, nested))
		}
	}
}
