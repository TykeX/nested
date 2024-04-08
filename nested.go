package nested

import (
	"fmt"
	"strconv"
	"strings"
)

// Get takes data map[string]any and keys ...string and it searches for value
// nested within data using keys. i.e. data["key1"]["key2"]["key3"].
//
// if a previous key returns array then it tries to convert current key into integer for indexing.
//
// if a previous key returns non-indexable (number, string, boolean) and there are remaining keys to search
// then it will return error and returns closest value it can find
func Get(data map[string]interface{}, keys ...string) (any, error) {
	v := any(data)
	for _, k := range keys {
		switch val := v.(type) {
		case map[string]any:
			temp, ok := val[k]
			if !ok {
				return temp, fmt.Errorf("either key '%s' doesn't exist or is nil", k)
			}
			v = temp
		case []interface{}:
			index, err := strconv.Atoi(k)
			if err != nil || index < 0 || index >= len(val) {
				return v, fmt.Errorf("can not index into array using %v", k)
			}
			v = val[index]
		default:
			return v, fmt.Errorf("previous value is neither map[string]any nor []any")
		}
	}
	return v, nil
}

func MustGet(data map[string]any, keys ...string) any {
	value, err := Get(data, keys...)
	if err != nil {
		panic(err)
	}
	return value
}

// Deprecated: use Get(data, strings.Split("key1.key2.key3", ".")
func Gets(data map[string]any, key string) (any, error) {
	return Get(data, strings.Split(key, ".")...)
}

// Deprecated: use MustGet(data, "key1", "key2")
func GetP(data map[string]any, keys ...string) any {
	return MustGet(data, keys...)
}

// Deprecated: use MustGet(data, strings.Split("key1.key2.key3", ".")
func GetsP(data map[string]any, key string) any {
	value, err := Get(data, strings.Split(key, ".")...)
	if err != nil {
		panic(fmt.Sprintf("failed to get %v", key))
	}
	return value
}

func isObject(d any) bool {
	if _, ok := d.(map[string]any); !ok {
		return false
	} else {
		return true
	}
}

func isArray(d any) bool {
	if _, ok := d.([]any); !ok {
		return false
	}
	return true
}
