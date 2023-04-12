package nested

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/tykex/ckret"
)

func TestGet(t *testing.T) {
	raw := []byte(`{
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

	value, err := Get(stuff, "k", "1")
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

func TestWithCkret(t *testing.T) {
	os.Setenv("ENVIRONMENT", "local")
	ckret.Init(&aws.Config{Region: aws.String("ap-south-1")})

	value, err := Get(ckret.GetCkret(), "kyc-comet", "INDIVIDUAL_CKYC", "PROVIDERS", "0")

	if err != nil {
		fmt.Printf("%v", err)
		t.FailNow()
	}

	if value.(string) != "DECENTRO_CKYC" {
		fmt.Printf("%v", value)
		t.FailNow()
	}
}

func TestGets(t *testing.T) {
	os.Setenv("ENVIRONMENT", "local")
	ckret.Init(&aws.Config{Region: aws.String("ap-south-1")})

	value, err := Gets(ckret.GetCkret(), "kyc-comet.INDIVIDUAL_CKYC.PROVIDERS.0")
	if err != nil {
		fmt.Printf("%v", err)
		t.FailNow()
	}

	if value.(string) != "DECENTRO_CKYC" {
		fmt.Printf("%v", value)
		t.FailNow()
	}

}
