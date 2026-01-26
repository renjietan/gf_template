package interceptor

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func ConvertToJSONSerializable(data interface{}) interface{} {
	if data == nil {
		return nil
	}

	v := reflect.ValueOf(data)
	return convertValue(v)
}

// 递归转换反射值
func convertValue(v reflect.Value) interface{} {
	// 处理指针，获取其指向的值
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	// 处理nil值
	if !v.IsValid() {
		return nil
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		// 处理切片和数组
		if v.Len() == 0 {
			return []interface{}{}
		}

		result := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			result[i] = convertValue(v.Index(i))
		}
		return result

	case reflect.Map:
		// 处理映射 - 键必须是字符串类型（JSON要求）
		if v.Len() == 0 {
			return map[string]interface{}{}
		}

		result := make(map[string]interface{})
		iter := v.MapRange()
		for iter.Next() {
			// 将键转换为字符串
			key := fmt.Sprintf("%v", iter.Key().Interface())
			result[key] = convertValue(iter.Value())
		}
		return result

	case reflect.Struct:
		// 处理结构体
		result := make(map[string]interface{})
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			fieldName := field.Name

			// 处理JSON tag
			if jsonTag := field.Tag.Get("json"); jsonTag != "" {
				// 提取JSON tag的名称（忽略omitempty等选项）
				if parts := strings.Split(jsonTag, ","); len(parts) > 0 && parts[0] != "" {
					fieldName = parts[0]
				}
				// 如果tag是"-"，则跳过该字段
				if fieldName == "-" {
					continue
				}
			}

			result[fieldName] = convertValue(v.Field(i))
		}
		return result

	case reflect.Interface:
		// 处理接口类型
		if v.IsNil() {
			return nil
		}
		return convertValue(v.Elem())

	case reflect.Bool:
		return v.Bool()

	case reflect.String:
		return v.String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()

	case reflect.Float32, reflect.Float64:
		return v.Float()

	case reflect.Complex64, reflect.Complex128:
		// 复数类型转换为字符串表示
		return fmt.Sprintf("%v", v.Complex())

	default:
		// 其他类型（如Chan, Func, UnsafePointer等）转换为字符串表示
		return fmt.Sprintf("%v", v.Interface())
	}
}

// PrintAsJSON 将数据以美观的JSON格式打印
func PrintAsJSON(data interface{}) (res string, err string) {
	jsonData := ConvertToJSONSerializable(data)

	jsonBytes, error := json.MarshalIndent(jsonData, "", "  ")
	if error != nil {
		return "", ""
	}
	return string(jsonBytes), ""
}
