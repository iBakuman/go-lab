package slice

import "testing"

func TestThreeIndex(t *testing.T) {
	kvs := []string{"a", "b", "c"}
	keyVals := kvs[:len(kvs):len(kvs)]
	keyVals = append(keyVals, "d")
	keyVals[0] = "a1"
	kvs[0] = "a2"
	t.Logf("keyVals: %v, len: %d, cap: %d", keyVals, len(keyVals), cap(keyVals))
	t.Logf("kvs: %v, len: %d, cap: %d", kvs, len(kvs), cap(kvs))
}
