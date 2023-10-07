package goist

import "reflect"

// SliceIntDiff2Right returns the elements in `a` that aren't in `b`.
func SliceIntDiff2Right(a, b []int) (c []int) {

	m := make(map[any]any)

	for _, v := range b {
		m[v] = 1
	}

	for _, v := range a {
		if _, ok := m[v]; !ok {
			c = append(c, v)
		}
	}

	return
}

func ToInterfaceSlice(sl interface{}) (isl []interface{}) {
	switch reflect.TypeOf(sl).Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return
	}

	s := reflect.ValueOf(sl)
	for i := 0; i < s.Len(); i++ {
		isl = append(isl, s.Index(i).Interface())
	}

	return
}