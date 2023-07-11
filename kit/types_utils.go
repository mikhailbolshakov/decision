package kit

import (
	"encoding/json"
	"github.com/iancoleman/strcase"
	"math"
	"reflect"
	"strconv"
	"time"
)

func MapsEqual(m1, m2 map[string]interface{}) bool {
	return Equal(m1, m2)
}

func Equal(m1, m2 any) bool {
	return reflect.DeepEqual(m1, m2)
}

func MapToLowerCamelKeys(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		return nil
	}
	r := make(map[string]interface{}, len(m))
	for k, v := range m {
		if vMap, ok := v.(map[string]interface{}); ok && len(vMap) > 0 {
			r[strcase.ToLowerCamel(k)] = MapToLowerCamelKeys(vMap)
		} else {
			r[strcase.ToLowerCamel(k)] = v
		}
	}
	return r
}

func MapInterfacesToBytes(m map[string]interface{}) []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}

func BytesToMapInterfaces(bytes []byte) map[string]interface{} {
	mp := make(map[string]interface{})
	_ = json.Unmarshal(bytes, &mp)
	return mp
}

func StringsToInterfaces(sl []string) []interface{} {
	if sl == nil {
		return nil
	}
	res := make([]interface{}, len(sl))
	for index, value := range sl {
		res[index] = value
	}

	return res
}

func ParseFloat32(s string) *float32 {
	if s == "" {
		return nil
	}
	fl64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return nil
	}
	fl32 := float32(fl64)
	return &fl32
}

func ParseFloat64(s string) *float64 {
	if s == "" {
		return nil
	}
	fl64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}
	return &fl64
}

func Round100(value float64) float64 {
	return math.Round(value*100) / 100
}

func Round10000(value float64) float64 {
	return math.Round(value*10000) / 10000
}

func IntToInt32Ptr(i *int) *int32 {
	if i == nil {
		return nil
	}
	v := int32(*i)
	return &v
}

func IntToInt64Ptr(i *int) *int64 {
	if i == nil {
		return nil
	}
	v := int64(*i)
	return &v
}

func Int32ToIntPtr(i *int32) *int {
	if i == nil {
		return nil
	}
	v := int(*i)
	return &v
}

func Int64ToIntPtr(i *int64) *int {
	if i == nil {
		return nil
	}
	v := int(*i)
	return &v
}

func IntPtr(i int) *int {
	return &i
}

func UInt32Ptr(i uint32) *uint32 {
	return &i
}

func Float32Ptr(i float32) *float32 {
	return &i
}

func Float64Ptr(i float64) *float64 {
	return &i
}

func TimePtr(t time.Time) *time.Time {
	return &t
}

func StringPtr(s string) *string {
	return &s
}

func NowPtr() *time.Time {
	return TimePtr(Now())
}

func BoolPtr(b bool) *bool {
	return &b
}
