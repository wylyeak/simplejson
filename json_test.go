package simplejson

import (
	"testing"
	"strconv"
)

var jsonString = `{
	"keyString":"stringValue",
	"keyInt": 123,
	"keyFloat64": 1.23,
	"keyFloat32": 1.23,
	"keyBool": true,
	"keyJSONObject": {
						"keyString":"stringValue",
						"keyInt": 123
					},
	"keyArray":[
		"string1",
		"string2"
	]
}`

var jsonArrayAllTypes = `{
	"keyString":["stringValue"],
	"keyInt": [123],
	"keyFloat64": [1.23],
	"keyFloat32": [1.23],
	"keyBool": [true],
	"keyJSONObject": [{
						"keyString":"stringValue",
						"keyInt": 123
					}],
	"keyArray":[[
		"string1",
		"string2"
	]]
}`

var jsonArrayAsRoot = `[
	{
		"elemKeyString": "stringValue",
		"elemKeyInt": 0
	},
	123,
	"stringElement"
]
`
var jsonArrayAsRoot2 = `[
	0,
	1,
	2
]
`

func TestJsonArrayWithInts(t *testing.T) {
	jsonArray, err := NewJSONArrayFromString(jsonArrayAsRoot2);
	if err != nil {
		t.Error("Parsing failed with error: " + err.Error())
		t.FailNow()
	}
	if jsonArray.Int(0) != 0 || jsonArray.Int(1) != 1 || jsonArray.Int(2) != 2 {
		t.Error("Parsing failed with error: " + err.Error())
		t.FailNow()
	}
}

func TestParseInt(t *testing.T) {
	var someFloat64 float64
	var someInt int
	var someInterface interface{}
	someInt = 123
	someFloat64 = 123
	someInterface = 123
	if parseInt(someInt) != 123 || parseInt(someFloat64) != 123 || parseInt(someInterface) != 123 {
		t.Error("parseInt could not parse interface")
		t.FailNow()
	}
}

func TestOpt(t *testing.T) {
	jsonObject, err := NewJSONObjectFromString(jsonString);
	if err != nil {
		t.Error("Parsing failed with error: " + err.Error())
		t.FailNow()
	}

	stringValue, ok := jsonObject.OptString("keyString")
	if (stringValue != "stringValue") {
		t.Error("keyString was " + stringValue + " instead of \"stringValue\" ")
	}

	stringValue, ok = jsonObject.OptString("keyString1")
	if (ok || stringValue != "") {
		t.Errorf("ok=%v stringValue=%v", ok, stringValue)
	}

	intValue, ok := jsonObject.OptInt("keyInt")
	if ( !ok || intValue != 123) {
		t.Errorf("keyInt was %d instead of 123 ", intValue)
	}

	intValue, ok = jsonObject.OptInt("keyInt1")
	if ( ok || intValue != 0) {
		t.Errorf("ok=%v intValue=%v", ok, intValue)
	}

	float32Value, ok := jsonObject.OptFloat32("keyFloat32")
	if (!ok || float32Value != 1.23) {
		t.Errorf("keyFloat32 was %d instead of 1.23 ", float32Value)
	}

	float32Value, ok = jsonObject.OptFloat32("keyFloat321")
	if (ok || float32Value != 0) {
		t.Errorf("ok=%v float32Value=%v", ok, float32Value)
	}

	float64Value, ok := jsonObject.OptFloat64("keyFloat64")
	if (!ok || float64Value != 1.23) {
		t.Errorf("keyFloat64 was %d instead of 1.23 ", float64Value)
	}

	float64Value, ok = jsonObject.OptFloat64("keyFloat641")
	if (ok || float64Value != 0) {
		t.Errorf("ok=%v float64Value=%v", ok, float64Value)
	}

	boolValue, ok := jsonObject.OptBool("keyBool")
	if (!ok || boolValue != true) {
		t.Errorf("keyBool was %b instead of true ", boolValue)
	}

	boolValue, ok = jsonObject.OptBool("keyBool1")
	if ( ok || boolValue != false) {
		t.Errorf("ok=%v boolValue=%v", ok, boolValue)
	}

	jsonValue, ok := jsonObject.OptJSONObject("keyJSONObject")
	if (!ok || jsonValue.String("keyString") != "stringValue") {
		t.Error("keyJSONObject didn't include the string strin")
	}

	jsonValue, ok = jsonObject.OptJSONObject("keyJSONObject1")
	if (ok || jsonValue != nil) {
		t.Errorf("ok=%v jsonValue=%v", ok, jsonValue)
	}
}

func TestJSONObjectAllTypesExceptArray(t *testing.T) {
	jsonObject, err := NewJSONObjectFromString(jsonString);
	if err != nil {
		t.Error("Parsing failed with error: " + err.Error())
		t.FailNow()
	}

	stringValue := jsonObject.String("keyString")
	if (stringValue != "stringValue") {
		t.Error("keyString was " + stringValue + " instead of \"stringValue\" ")
	}

	intValue := jsonObject.Int("keyInt")
	if (intValue != 123) {
		t.Errorf("keyInt was %d instead of 123 ", intValue)
	}

	float32Value := jsonObject.Float32("keyFloat32")
	if (float32Value != 1.23) {
		t.Errorf("keyFloat32 was %d instead of 1.23 ", float32Value)
	}

	float64Value := jsonObject.Float64("keyFloat64")
	if (float64Value != 1.23) {
		t.Errorf("keyFloat64 was %d instead of 1.23 ", float64Value)
	}

	boolValue := jsonObject.Bool("keyBool")
	if (boolValue != true) {
		t.Errorf("keyBool was %b instead of true ", boolValue)
	}

	jsonValue := jsonObject.JSONObject("keyJSONObject")
	if (jsonValue.String("keyString") != "stringValue") {
		t.Error("keyJSONObject didn't include the string strin")
	}
}

func TestJSONArrayAllTypes(t *testing.T) {
	jsonObject, err := NewJSONObjectFromString(jsonArrayAllTypes);
	if err != nil {
		t.Error("Parsing failed with error: " + err.Error())
		t.FailNow()
	}

	stringValue := jsonObject.JSONArray("keyString").String(0)
	if (stringValue != "stringValue") {
		t.Error("keyString was " + stringValue + " instead of \"stringValue\" ")
	}

	intValue := jsonObject.JSONArray("keyInt").Int(0)
	if (intValue != 123) {
		t.Errorf("keyInt was %d instead of 123 ", intValue)
	}

	float32Value := jsonObject.JSONArray("keyFloat32").Float32(0)
	if (float32Value != 1.23) {
		t.Errorf("keyFloat32 was %d instead of 1.23 ", float32Value)
	}

	float64Value := jsonObject.JSONArray("keyFloat64").Float64(0)
	if (float64Value != 1.23) {
		t.Errorf("keyFloat64 was %d instead of 1.23 ", float64Value)
	}

	boolValue := jsonObject.JSONArray("keyBool").Bool(0)
	if (boolValue != true) {
		t.Errorf("keyBool was %b instead of true ", boolValue)
	}

	jsonValue := jsonObject.JSONArray("keyJSONObject").JSONObject(0)
	if (jsonValue.String("keyString") != "stringValue") {
		t.Error("keyJSONObject didn't include the string strin")
	}
	keys := jsonObject.Keys()
	if len(keys) != 7 {
		t.Errorf("keys:%v", keys)
	}
}

func TestParsingStringToJSONArray(t *testing.T) {
	jsonArray, err := NewJSONArrayFromString(jsonArrayAsRoot);
	if err != nil {
		t.Error("Parsing failed with error: " + err.Error())
		t.FailNow()
	}

	object0 := jsonArray.JSONObject(0)
	if stringValue := object0.String("elemKeyString"); stringValue != "stringValue" {
		t.Error("object[0]::elemKeyString was " + stringValue + " instead of \"stringValue\" ")
	}
	if intValue := object0.Int("elemKeyInt"); intValue != 0 {
		t.Error("object[0]::elemKeyInt was " + strconv.Itoa(intValue) + " instead of 0 ")
	}

	object1 := jsonArray.Int(1)
	if object1 != 123 {
		t.Error("object1 was %d instead of 123 ", object1)
	}

	object2 := jsonArray.String(2)
	if object2 != "stringElement" {
		t.Error("object1 was " + object2 + " instead of \"stringElement\"")
	}
}

func TestSet(t *testing.T) {
	jsonobject := NewJSONObject();
	jsonobject.Set("object1", "stringValue")
	if value := jsonobject.String("object1"); value != "stringValue" {
		t.Error("object1 was " + value + " instead of \"stringValue\"")
	}

	object2 := make([]float32, 5, 5)
	object2[3] = 19.88
	jsonobject.Set("object2", object2)
	valueArray := jsonobject.JSONArray("object2")
	if current := valueArray.Int(0); current != 0 {
		t.Error("object2[0] was %d instead of 0", current)
	} else if current := valueArray.Int(1); current != 0 {
		t.Error("object2[1] was %d instead of 0", current)
	} else if current := valueArray.Int(2); current != 0 {
		t.Error("object2[2] was %d instead of 0", current)
	} else if current := valueArray.Float32(3); current != 19.88 {
		t.Error("object2[3] was %d instead of 1988", current)
	} else if current := valueArray.Int(4); current != 0 {
		t.Error("object2[4] was %d instead of 0", current)
	}
}
