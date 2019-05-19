package simplejson

import "encoding/json"

type JSONObject struct {
	innerMap map[string]interface{}
}

// NewJSONObjectFromString returns a new JSONObject, parsed from a string or an error if unsuccessful
func NewJSONObjectFromString(jsonobject string) (*JSONObject, error) {
	var resultingMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonobject), &resultingMap)
	if err != nil {
		return nil, err
	}
	return &JSONObject{resultingMap}, nil
}

// NewJSONObject creates and returns a new JSONObject
func NewJSONObject() *JSONObject {
	return &JSONObject{make(map[string]interface{})}
}

func newJSONObjectWithMap(fromMap map[string]interface{}) *JSONObject {
	return &JSONObject{fromMap}
}

// JSONArray returns JSONArray from specific key
func (this *JSONObject) JSONArray(key string) *JSONArray {
	return &JSONArray{interfaceToInterfaceArray(this.innerMap[key])}
}

// OptJSONArray like JSONArray
func (this *JSONObject) OptJSONArray(key string) (*JSONArray, bool) {
	obj, ok := this.innerMap[key]
	if !ok {
		return nil, ok
	} else {
		return &JSONArray{interfaceToInterfaceArray(obj)}, true
	}
}

// JSONObject returns JSONObject from specific key
func (this *JSONObject) JSONObject(key string) *JSONObject {
	return &JSONObject{this.innerMap[key].(map[string]interface{})}
}

func (this *JSONObject) OptJSONObject(key string) (*JSONObject, bool) {
	obj, ok := this.innerMap[key]
	if !ok {
		return nil, ok
	} else {
		return &JSONObject{obj.(map[string]interface{})}, true
	}
}


// String returns string from specific key
func (this *JSONObject) String(key string) string {
	return this.innerMap[key].(string)
}

func (this *JSONObject) OptString(key string) (string, bool) {
	obj, ok := this.innerMap[key]
	if !ok {
		return "", ok
	} else {
		return obj.(string), true
	}
}

// Bool returns bool from specific key
func (this *JSONObject) Bool(key string) bool {
	return this.innerMap[key].(bool)
}

func (this *JSONObject) OptBool(key string) (bool, bool) {
	obj, ok := this.innerMap[key]
	if !ok {
		return false, ok
	} else {
		return obj.(bool), true
	}
}

// Int returns int from specific key
func (this *JSONObject) Int(key string) int {
	return parseInt(this.innerMap[key])
}

func (this *JSONObject) Int64(key string) int64 {
	return parseInt64(this.innerMap[key])
}

func (this *JSONObject) OptInt(key string) (int, bool) {
	obj, ok := this.innerMap[key]
	if !ok || obj == nil {
		return 0, ok
	} else {
		return parseInt(obj), true
	}

}

func (this *JSONObject) OptInt64(key string) (int64, bool) {
	obj, ok := this.innerMap[key]
	if !ok || obj == nil {
		return 0, ok
	} else {
		return parseInt64(obj), true
	}

}


// Float32 returns float32 from specific key
func (this *JSONObject) Float32(key string) float32 {
	return float32(this.innerMap[key].(float64))
}

func (this *JSONObject) OptFloat32(key string) (float32, bool) {
	obj, ok := this.innerMap[key]
	if !ok {
		return float32(0), ok
	} else {
		return float32(obj.(float64)), true
	}
}

// Float64 returns float64 from specific key
func (this *JSONObject) Float64(key string) float64 {
	return this.innerMap[key].(float64)
}

func (this *JSONObject) OptFloat64(key string) (float64, bool) {
	obj, ok := this.innerMap[key]
	if !ok {
		return float64(0), ok
	} else {
		return obj.(float64), true
	}
}

// Set sets the value of Key
func (this *JSONObject) Set(key string, value interface{}) bool {
	unmarshalled, ok := interfaceToJsonCompatible(value)
	if !ok {
		return ok
	}
	this.innerMap[key] = unmarshalled
	return true
}

// String return json-representation as string
func (this *JSONObject) AsString() (string, error) {
	jsonString, err := json.Marshal(this.innerMap)
	return string(jsonString), err
}

func (this *JSONObject) Keys() ([]string) {
	var result []string
	for key := range this.innerMap {
		result = append(result, key)
	}
	return result
}

func (this *JSONObject) ToMap() map[string]interface{} {
	return this.innerMap
}

func (this *JSONObject) MustUnmarshal(v interface{}) {
	str, err := this.AsString()
	if err != nil {
		panic(err)
	} else {
		err := json.Unmarshal([]byte(str), v)
		if err != nil {
			panic(err)
		}
	}

}

func (this *JSONObject) Unmarshal(v interface{}) error {
	str, err := this.AsString()
	if err != nil {
		return err
	} else {
		return json.Unmarshal([]byte(str), v)
	}

}