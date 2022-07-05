package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	type t struct {
		STR string
		NUM int64
		BOO bool
	}
	var n = t{"пыпыпы", 5, false}
	var m = make(map[string]interface{}, 1)
	m["STR"] = "чагыр-чагыр"
	m["NUM"] = int64(100)
	m["BOO"] = true
	m["NULL"] = "выаываываыва"

	fmt.Println("\n", n)
	setCustomStructureFields(&n, m)
	fmt.Println("\n", n)

}

func setCustomStructureFields(in interface{}, values map[string]interface{}) error {
	fmt.Println(values)
	if in == nil {
		return errors.New("in shouldn't be equal null")
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return errors.New("in should be a struct type")
	}

	for k, v := range values {
		typeField, ok := val.Type().FieldByName(k)

		if !ok {
			continue
		}

		switch typeField.Type.Kind() {
		case reflect.Int64:
			val.FieldByName(k).SetInt(v.(int64))
		case reflect.String:
			val.FieldByName(k).SetString(v.(string))
		case reflect.Bool:
			val.FieldByName(k).SetBool(v.(bool))
		default:
			return errors.New("Too dificult to implement")
		}

	}

	// for k, v := range values {
	// 	//typeField := reflect.ValueOf(in).Elem().FieldByName(k).Type()
	// 	switch reflect.ValueOf(in).Elem().FieldByName(k).Type().(type) {
	// 	case int:
	// 		reflect.ValueOf(in).Elem().FieldByName(k).SetInt(v.(int64))
	// 	case reflect.Float32:
	// 		reflect.ValueOf(in).Elem().FieldByName(k).SetFloat(v.(float64))
	// 	case reflect.String:
	// 		reflect.ValueOf(in).Elem().FieldByName(k).SetString(v.(string))
	// 	case reflect.Bool:
	// 		reflect.ValueOf(in).Elem().FieldByName(k).SetBool(v.(bool))
	// 	default:
	// 		return errors.New("Too dificult to implement")
	// 	}
	// }

	// for i := 0; i < val.NumField(); i++ {
	// 	typeField := val.Type().Field(i)
	// 	fmt.Println(typeField.Name)
	// 	fmt.Println(val.Field(i))
	// 	j, _ := values[typeField.Name].(int64)

	// 	switch typeField.Type.Kind() {
	// 	case reflect.Int:
	// 		val.SetInt(j)
	// 	case reflect.Float32:
	// 		//return v.String()
	// 	case reflect.Float64:
	// 		// return strconv.Itoa(v)
	// 	case reflect.String:
	// 		// return typeField.(string)
	// 	case reflect.Bool:
	// 	case reflect.Chan:
	// 	default:
	// 		return errors.New("Too dificult to implement")
	// 	}

	// 	// v, ok := values[typeField.Name]
	// 	// if ok {
	// 	// 	if typeField.Name == v {
	// 	// 		val.Field(i)
	// 	// 	}
	// 	// }

	// 	// if typeField.Type.Kind() == reflect.Struct {
	// 	// 	log.Printf("nested field: %v", typeField.Name)
	// 	// 	//PrintStruct(val.Field(i).Interface())
	// 	// 	continue
	// 	// }
	// }
	return nil
}
