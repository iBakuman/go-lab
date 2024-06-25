package utils

import "encoding/json"

func MustBeautifyJson(jsonStr string) string {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		panic(err)
	}
	prettyJson, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(prettyJson)
}
