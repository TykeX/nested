package nested

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	raw := []byte(`{
    "n":null,
    "k":[false, 3, 4, true, "string"],
    "d": {
      "e": [ { "name": "mango" }, { "name": "banana" } ],
      "f": {"j": false}
    },
    "m": {
      "0": "zero",
      "1": {
        "2": "two",
        "3": [1,2,43]
      }
    }
  }`)

	var stuff map[string]any
	err := json.Unmarshal(raw, &stuff)
	if err != nil {
		panic("can not parse test json")
	}

	value, err := Get(stuff, "d", "f", "a")
	if err == nil {
		fmt.Printf("err should be non-nil and value should be nil; err is %v value is %v", err, value)
		t.Fail()
	}

	value, err = Get(stuff, "n")
	if err != nil {
		fmt.Printf("err should be nil and value should be nil; err is %v and value is %v", err, value)
		t.Fail()
	}

	if value != nil {
		fmt.Printf("err should be nil and value should be nil; err is %v and value is %v", err, value)
		t.Fail()
	}

	value, err = Get(stuff, "k", "1")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if value.(float64) != 3 {
		t.Fail()
	}

	value, err = Get(stuff, "k", "d", "a")
	if err == nil {
		t.Fail()
	}

	value, err = Get(stuff, "d", "e", "0", "name")
	if err != nil {
		t.Fail()
	}

	if value.(string) != "mango" {
		t.Fail()
	}

	value, err1 := Get(stuff, "d", "e", "0", "name", "cow")
	if err1 == nil {
		t.Fail()
	}

	value, err = Get(stuff, "d", "e", "100", "name")
	if err == nil {
		t.Fail()
	}
	value, err = Get(stuff, "m", "1", "2")
	if err != nil {
		t.Fail()
	}

	if value.(string) != "two" {
		t.Fail()
	}

	//--

	value, err = Get(stuff, "m", "1", "3", "2")
	if err != nil {
		t.Fail()
	}

	if value.(float64) != 43 {
		t.Fail()
	}
}

// must panic
func TestGetsP(t *testing.T) {
	raw := []byte(`{ "name": {"first_name": "pawan"} }`)

	var stuff map[string]any
	err := json.Unmarshal(raw, &stuff)
	if err != nil {
		panic("can not parse json")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	value := GetsP(stuff, "name.last_name")
	fmt.Printf("%v", value)
}

// must panic
func TestGetP(t *testing.T) {
	raw := []byte(`{ "name":"pawan" }`)

	var stuff map[string]any
	err := json.Unmarshal(raw, &stuff)
	if err != nil {
		panic("can not parse json")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	value := GetP(stuff, "name", "last_name")
	fmt.Printf("%v", value)
}

// Edge cases

func TestEdgeCase(t *testing.T) {
	raw := []byte(`{ "first_name": "pawan", "last_name": null}`)
	var stuff map[string]any
	err := json.Unmarshal(raw, &stuff)
	if err != nil {
		panic("can not parse json")
	}

	value, err := Get(stuff, "first_name", "k")
	if err == nil {
		t.Log("err must not be nil")
		t.Fail()
	}

	value, err = Get(stuff, "last_name")
	if err != nil {
		t.Log("err must be nil")
		t.Fail()
	}
	if value != nil {
		t.Log("value must be nil")
		t.Fail()
	}

	value, err = Get(stuff, "non_existent_key")
	if err == nil {
		t.Log("err must not be nil")
		t.Fail()
	}
}
