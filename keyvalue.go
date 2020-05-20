package otgorm

import (
	"fmt"

	"go.opentelemetry.io/otel/api/kv"
)

//CreateSpanAttribute creates a KeyValue for use as a span attribute
func CreateSpanAttribute(k string, v interface{}) kv.KeyValue {
	switch x := v.(type) {
	case string:
		return kv.String(k, v.(string))
	case bool:
		return kv.Bool(k, v.(bool))
	case int64:
		return kv.Int64(k, v.(int64))
	case int32:
		return kv.Int32(k, v.(int32))
	case int:
		return kv.Int(k, v.(int))
	case float64:
		return kv.Float64(k, v.(float64))
	case float32:
		return kv.Float32(k, v.(float32))
	case uint:
		return kv.Uint(k, v.(uint))
	case uint64:
		return kv.Uint64(k, v.(uint64))
	case uint32:
		return kv.Uint32(k, v.(uint32))
	default:
		return kv.String("attribute.error", fmt.Sprintf("Couldn't convert %s into KeyValue", x))
	}
}
