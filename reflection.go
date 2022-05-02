package reflection

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Reflection struct {
	refType      reflect.Type
	refValue     reflect.Value
	refKind      reflect.Kind
	refTypeElem  reflect.Type
	refValueElem reflect.Value
	isStructPtr  bool
	isStruct     bool

	Type          string
	Name          string
	Namespace     string
	NameWithSpace string
	PkgPath       string
	Fields        map[string]*ReflectionField
	FieldNames    []string
	Methods       map[string]*ReflectionMethod
	MethodNames   []string
}

func New(object interface{}) *Reflection {
	r := new(Reflection)
	r.refType = reflect.TypeOf(object)
	r.refValue = reflect.ValueOf(object)
	r.refTypeElem = r.refType
	r.refValueElem = r.refValue
	r.refKind = r.refType.Kind()
	r.isStruct = r.refKind == reflect.Struct
	r.isStructPtr = r.refKind == reflect.Ptr && r.refType.Elem().Kind() == reflect.Struct
	if !r.isStruct && !r.isStructPtr {
		panic(fmt.Errorf("invalid instance kind `%s`, struct or struct ptr is needed", r.refKind.String()))
	}
	if r.isStructPtr {
		r.refTypeElem = r.refType.Elem()
		r.refValueElem = r.refValue.Elem()
	}
	r.PkgPath = r.refTypeElem.PkgPath()
	r.Namespace = strings.Join(strings.Split(r.PkgPath, "/")[1:], ".")
	r.Type = r.refType.String()
	r.Name = r.refTypeElem.Name()
	r.NameWithSpace = fmt.Sprintf("%s.%s", r.Namespace, r.Name)
	r.loadFields()
	r.loadMethods()
	return r
}

func (r *Reflection) loadFields() {
	r.Fields = make(map[string]*ReflectionField, 0)
	r.FieldNames = make([]string, 0)
	if r.refValueElem.IsValid() {
		count := r.refTypeElem.NumField()
		for i := 0; i < count; i++ {
			structField := r.refTypeElem.Field(i)
			fieldValue := r.refValueElem.Field(i)
			field := newReflectionField(structField, fieldValue)
			r.FieldNames = append(r.FieldNames, field.Name)
			r.Fields[field.Name] = field
		}
	}
}

func (r *Reflection) loadMethods() {
	r.Methods = make(map[string]*ReflectionMethod, 0)
	r.MethodNames = make([]string, 0)
	count := r.refTypeElem.NumMethod()
	for i := 0; i < count; i++ {
		structMethod := r.refTypeElem.Method(i)
		methodValue := r.refValueElem.Method(i)
		method := newReflectionMethod(structMethod, methodValue)
		r.MethodNames = append(r.MethodNames, method.Name)
		r.Methods[method.Name] = method
	}
	count = r.refType.NumMethod()
	for i := 0; i < count; i++ {
		structMethod := r.refType.Method(i)
		methodValue := r.refValue.Method(i)
		method := newReflectionMethod(structMethod, methodValue)
		r.MethodNames = append(r.MethodNames, method.Name)
		r.Methods[method.Name] = method
	}
}

func (r *Reflection) String() string {
	str, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(str)
}

func (r *Reflection) HasField(name string) bool {
	_, ok := r.Fields[name]
	return ok
}

func (r *Reflection) HasMethod(name string) bool {
	_, ok := r.Methods[name]
	return ok
}

func (r *Reflection) Get(name string) (interface{}, error) {
	field, ok := r.Fields[name]
	if ok {
		return field.Value, nil
	}
	return nil, fmt.Errorf("field `%s` not exists", name)
}

func (r *Reflection) MustGet(name string) interface{} {
	value, err := r.Get(name)
	if err != nil {
		panic(err)
	}
	return value
}

func (r *Reflection) Set(name string, value interface{}) error {
	if !r.HasField(name) {
		return fmt.Errorf("field `%s` not exists", name)
	}
	field := r.Fields[name]
	return field.Set(value)
}

func (r *Reflection) MustSet(name string, value interface{}) {
	if err := r.Set(name, value); err != nil {
		panic(err)
	}
}

func (r *Reflection) GetTag(field string, tagName string) (string, error) {
	if !r.HasField(field) {
		return "", fmt.Errorf("field `%s` not exists", field)
	}
	if value, exists := r.Fields[field].LookUpTag(tagName); !exists {
		return "", fmt.Errorf("tag `%s` not exists", tagName)
	} else {
		return value, nil
	}
}

func (r *Reflection) MustGetTag(field string, tagName string) string {
	if value, err := r.GetTag(field, tagName); err != nil {
		panic(err)
	} else {
		return value
	}
}

func (r *Reflection) Call(name string, args ...interface{}) ([]interface{}, error) {
	method, ok := r.Methods[name]
	if ok {
		return method.Call(args...)
	}
	return nil, fmt.Errorf("method `%s` not exists", name)
}

func (r *Reflection) MustCall(name string, args ...interface{}) []interface{} {
	result, err := r.Call(name, args...)
	if err != nil {
		panic(err)
	}
	return result
}
