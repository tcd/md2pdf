package parse

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// AreEqualJSON compares two json-encoded structs.
func AreEqualJSON(b1, b2 []byte) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal(b1, &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal(b2, &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}
