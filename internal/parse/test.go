package parse

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// AreStructsEqual compares two structs by encoding them into byte slices,
// decoding the bytes into empty interfaces,
// and comparing the interfaces with reflect.DeepEqual.
func AreStructsEqual(i1, i2 interface{}) (areEqual bool, err error) {
	b1, err := json.Marshal(i1)
	if err != nil {
		return
	}
	b2, err := json.Marshal(i2)
	if err != nil {
		return
	}

	var o1 interface{}
	var o2 interface{}
	err = json.Unmarshal(b1, &o1)
	if err != nil {
		err = fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
		return
	}
	err = json.Unmarshal(b2, &o2)
	if err != nil {
		err = fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
		return
	}

	return reflect.DeepEqual(o1, o2), nil
}

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
