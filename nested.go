package nested

import (
	"errors"
	"strconv"
	"strings"
)

// Get takes data map[string]any and keys ...string and it searches for value
// nested within data using keys. for example data["key1"]["key2"]["key3"].
//
// if any previous key returns array then it tries to convert current key into integer for indexing.
//
// if any previous key returns non-indexable (number, string, boolean) and there are remaining keys to search
// then it will return error and returns closest value it can find
func Get(data map[string]any, keys ...string) (any, error) {
	value := any(data)
	for _, key := range keys {
		if isObject(value) {
			value = value.(map[string]any)[key]
		} else if isArray(value) {
			index, err := strconv.ParseInt(key, 10, 64)
			if err != nil {
				return value, errors.New("got array but key can not be converted to int")
			}
			if index >= int64(len(value.([]any))) {
				return value, errors.New("index out of bound")
			}
			value = value.([]any)[index].(any)
		} else {
			// there is a key that it need to get but the prev value is neither map nor array
			return value, errors.New("previous value is not (array or map)")
		}
	}
	return value, nil
}

// Gets() calls Get() after spliting key on dot (.)
func Gets(data map[string]any, key string) (any, error) {
	return Get(data, strings.Split(key, ".")...)
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
