package kinds

import "reflect"

//
//
//
//
//
func isClassedNull(o any) bool {
	rv := reflect.ValueOf(o)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

//
func isEmpty(o any) bool {
	rv := reflect.ValueOf(o)
	switch rv.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return rv.Len() == 0
	default:
		return false
	}
}
