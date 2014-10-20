package fixedlen

import "reflect"
import "fmt"
import "bytes"

func Marshal(in interface{}) string {
	t := reflect.TypeOf(in)
	if t.Kind() == reflect.Struct {
		fmt.Println("It's a struct.")
		printInfo(in, t)
	} else {
		fmt.Println("It's not a struct.")
	}
	return t.Name()
}

func printInfo(in interface{}, t reflect.Type) {
	var b bytes.Buffer
	cnt := t.NumField()
	fmt.Println("Field count:", cnt)
	for i := 0; i < cnt; i++ {
		fld := t.Field(i)
		fmt.Println("Field name:", fld.Name, "Tag:", fld.Tag.Get("fw"))
		b.WriteString(asString(t.Field(i).Type, reflect.ValueOf(in).Field(i), 10))
	}
	fmt.Println(b.String())
}

func asString(t reflect.Type, v reflect.Value, len int) string {
	var s string

	switch t.Kind() {
	case reflect.Bool:
		if v.Bool() {
			s = "T"
		} else {
			s = "F"
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		s = "NUMBER"
	default:
		s = "UNKNOWN"
	}
	return s
}
