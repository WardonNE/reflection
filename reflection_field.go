package reflection

import (
	"fmt"
	"reflect"
)

type ReflectionField struct {
	structField reflect.StructField
	refType     reflect.Type
	refKind     reflect.Kind
	refValue    reflect.Value
	refTag      reflect.StructTag

	Name  string
	Type  string
	Kind  string
	Value interface{}
}

func newReflectionField(structField reflect.StructField, value reflect.Value) *ReflectionField {
	field := new(ReflectionField)
	field.structField = structField
	field.refType = structField.Type
	field.refKind = structField.Type.Kind()
	field.refValue = value

	field.Name = structField.Name
	field.Type = field.refType.Name()
	field.Kind = field.refKind.String()
	if value.CanInterface() {
		field.Value = value.Interface()
	}
	field.refTag = field.structField.Tag
	return field
}

func (field *ReflectionField) GetReflectType() reflect.Type {
	return field.refType
}

func (field *ReflectionField) GetReflectValue() reflect.Value {
	return field.refValue
}

func (field *ReflectionField) GetReflectKind() reflect.Kind {
	return field.refKind
}

func (field *ReflectionField) IsAnonymous() bool {
	return field.structField.Anonymous
}

func (field *ReflectionField) IsValid() bool {
	return field.refValue.IsValid()
}

func (field *ReflectionField) CanSet() bool {
	return field.refValue.CanSet()
}

func (field *ReflectionField) Set(value interface{}) error {
	if !field.CanSet() {
		return fmt.Errorf("field is not settable")
	}
	field.refValue.Set(reflect.ValueOf(value))
	return nil
}

func (field *ReflectionField) MustSet(value interface{}) {
	if err := field.Set(value); err != nil {
		panic(err)
	}
}

func (field *ReflectionField) GetTag(key string) string {
	return field.refTag.Get(key)
}

func (field *ReflectionField) LookUpTag(key string) (string, bool) {
	return field.refTag.Lookup(key)
}
