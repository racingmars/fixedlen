package fixedlen

import "reflect"
import "bytes"
import "strconv"

func Encode(in interface{}) []byte {
	var b bytes.Buffer

	t := reflect.TypeOf(in)
	if t.Kind() != reflect.Struct {
		return []byte("not struct")
	}
	v := reflect.ValueOf(in)

	cnt := t.NumField()
	for i := 0; i < cnt; i++ {
		fldt := t.Field(i)
		fldv := v.Field(i)
		len, _ := strconv.Atoi(fldt.Tag.Get("fw"))
		// TODO: what if there isn't a length tag?
		b.WriteString(asString(fldt.Type, fldv, len))
	}

	return b.Bytes()
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
		s = strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		s = strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		s = strconv.FormatFloat(v.Float(), 'g', -1, 32)
	case reflect.Float64:
		s = strconv.FormatFloat(v.Float(), 'g', -1, 64)
	case reflect.String:
		s = v.String()
	default:
		s = "UNKNOWN"
	}

	return padString(s, len)
}

func padString(s string, l int) string {
	if i := l - len(s); i > 0 {
		pad := ""
		for j := 0; j < i; j++ {
			pad += " "
		}
		return s + pad
	} else {
		return s
	}
}
