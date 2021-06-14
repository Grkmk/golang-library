package slices

// import "reflect"

func FindString(sl *[]string, i *string) (*string, bool) {
	for _, c := range *sl {
		if c == *i {
			return &c, true
		}
	}

	return nil, false
}

// returns -1 if unable to find element
func IndexOfString(sl *[]string, s *string) (int, bool) {
	var i int
	for _, el := range *sl {
		if el == *s {
			return i, true
		}
		i++
	}

	return -1, false
}

// func takeSliceArg(arg interface{}) ([]interface{}, bool) {
// 	slice, ok := takeArg(arg, reflect.Slice)
// 	if !ok {
// 		return nil, false
// 	}

// 	c := slice.Len()
// 	out := make([]interface{}, c)

// 	for i := 0; i < c; i++ {
// 		out[i] = slice.Index(i).Interface()
// 	}

// 	return out, true
// }

// func takeArg(arg interface{}, kind reflect.Kind) (reflect.Value, bool) {
// 	val := reflect.ValueOf(arg)

// 	return val, val.Kind() == kind
// }
