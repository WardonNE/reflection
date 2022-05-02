package reflection

import (
	"fmt"
	"reflect"
)

type ReflectionMethod struct {
	structMethod reflect.Method
	refType      reflect.Type
	refValue     reflect.Value
	refKind      reflect.Kind

	Name string
	Kind string
}

func newReflectionMethod(structMethod reflect.Method, value reflect.Value) *ReflectionMethod {
	method := new(ReflectionMethod)
	method.structMethod = structMethod
	method.refType = structMethod.Type
	method.refValue = value
	method.refKind = structMethod.Type.Kind()

	method.Name = structMethod.Name
	method.Kind = method.refKind.String()

	return method
}

func (method *ReflectionMethod) IsExported() bool {
	return method.structMethod.IsExported()
}

func (method *ReflectionMethod) Call(args ...interface{}) ([]interface{}, error) {
	if !method.IsExported() {
		return nil, fmt.Errorf("method is not exported")
	}
	in := make([]reflect.Value, 0)
	for _, arg := range args {
		in = append(in, reflect.ValueOf(arg))
	}
	var outs []reflect.Value
	if method.refType.IsVariadic() {
		outs = method.refValue.CallSlice(in)
	} else {
		outs = method.refValue.Call(in)
	}
	returns := make([]interface{}, 0)
	for _, out := range outs {
		returns = append(returns, out.Interface())
	}
	return returns, nil
}

func (method *ReflectionMethod) MustCall(args ...interface{}) []interface{} {
	if returns, err := method.Call(args...); err != nil {
		panic(err)
	} else {
		return returns
	}
}
