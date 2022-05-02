package reflection

import (
	"testing"
)

type MyTestStruct struct {
	Field1        string `json:"field1" xml:"field1xml"`
	Field2        string `json:"field2" xml:"-"`
	privateField1 string
}

func (m *MyTestStruct) Init() {
	m.Field1 = "field1"
	m.Field2 = "field2"
}

func (m *MyTestStruct) Method1() string {
	return "method1"
}

func (m *MyTestStruct) Method2(str string) string {
	return str
}

func (m *MyTestStruct) Method3(str ...string) []string {
	return str
}

var myStruct = &MyTestStruct{
	Field1: "field1",
	Field2: "field2",
}

var myReflection = New(myStruct)

func TestHasField(t *testing.T) {
	if !myReflection.HasField("Field1") {
		t.FailNow()
	}
	if !myReflection.HasField("Field2") {
		t.FailNow()
	}
	if myReflection.HasField("Field3") {
		t.FailNow()
	}
}

func TestHasMethod(t *testing.T) {
	if !myReflection.HasMethod("Method1") {
		t.FailNow()
	}
	if !myReflection.HasMethod("Method2") {
		t.FailNow()
	}
	if !myReflection.HasMethod("Method3") {
		t.FailNow()
	}
}

func TestGet(t *testing.T) {
	if value, err := myReflection.Get("Field1"); value != "field1" || err != nil {
		t.FailNow()
	}
	if value, err := myReflection.Get("Field2"); value != "field2" || err != nil {
		t.FailNow()
	}
}

func TestMustGet(t *testing.T) {
	if myReflection.MustGet("Field1") != "field1" {
		t.FailNow()
	}
	if myReflection.MustGet("Field2") != "field2" {
		t.FailNow()
	}
}

func TestSet(t *testing.T) {
	if err := myReflection.Set("Field1", "setField1"); err != nil || myStruct.Field1 != "setField1" {
		t.FailNow()
	}
	if err := myReflection.Set("Field2", "setField2"); err != nil || myStruct.Field2 != "setField2" {
		t.FailNow()
	}
	if err := myReflection.Set("privateField1", "privateField1"); err == nil {
		t.FailNow()
	}
}

func TestMustSet(t *testing.T) {
	myReflection.MustSet("Field1", "mustSetField1")
	if myStruct.Field1 != "mustSetField1" {
		t.FailNow()
	}
	myReflection.MustSet("Field2", "mustSetField2")
	if myStruct.Field2 != "mustSetField2" {
		t.FailNow()
	}
	defer func() {
		if err := recover(); err == nil {
			t.FailNow()
		}
	}()
	myReflection.MustSet("privateField1", "privateField1")
}

func TestGetTag(t *testing.T) {
	if value, err := myReflection.GetTag("Field1", "json"); err != nil || value != "field1" {
		t.FailNow()
	}
	if value, err := myReflection.GetTag("Field2", "json"); err != nil || value != "field2" {
		t.FailNow()
	}
	if value, err := myReflection.GetTag("Field1", "xml"); err != nil || value != "field1xml" {
		t.FailNow()
	}
	if value, err := myReflection.GetTag("Field2", "xml"); err != nil || value != "-" {
		t.FailNow()
	}
}

func TestMustGetTag(t *testing.T) {
	if myReflection.MustGetTag("Field1", "json") != "field1" {
		t.FailNow()
	}
	if myReflection.MustGetTag("Field1", "xml") != "field1xml" {
		t.FailNow()
	}
	if myReflection.MustGetTag("Field2", "json") != "field2" {
		t.FailNow()
	}
	if myReflection.MustGetTag("Field2", "xml") != "-" {
		t.FailNow()
	}
}

func TestCall(t *testing.T) {
	if value, err := myReflection.Call("Method1"); err != nil || value[0].(string) != "method1" {
		t.FailNow()
	}
	if value, err := myReflection.Call("Method2", "method2"); err != nil || value[0].(string) != "method2" {
		t.FailNow()
	}
	if value, err := myReflection.Call("Method3", []string{"method3"}); err != nil || value[0].([]string)[0] != "method3" {
		t.FailNow()
	}
}

func TestMustCall(t *testing.T) {
	if myReflection.MustCall("Method1")[0].(string) != "method1" {
		t.FailNow()
	}
	if myReflection.MustCall("Method2", "method2")[0].(string) != "method2" {
		t.FailNow()
	}
	if myReflection.MustCall("Method3", []string{"method3"})[0].([]string)[0] != "method3" {
		t.FailNow()
	}
}
