package configo

import "testing"

func assertEqual(t *testing.T, o, e interface{}) {
	if o != e {
		t.Errorf("expected %d, got %d", e, o)
	}
}

func assertNil(t *testing.T, o interface{}) {
	if o != nil {
		t.Error("expected nil, got", o)
	}
}

func assertNotNil(t *testing.T, o interface{}) {
	if o == nil {
		t.Error("expected NOT nil, got", o)
	}
}

func TestInt_Simple(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, _ := f.Get("int_val").asInt()
	assertEqual(t, o, 345)
}
func TestInt_NonExistantKey_NoDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("int_val_not_there").asInt()
	assertNotNil(t, err)
	assertEqual(t, o, 0)
}
func TestInt_NonExistantKey_HasDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("int_val_not_there").Default("23").asInt()
	assertNil(t, err)
	assertEqual(t, o, 23)
}

func TestBool_Simple(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, _ := f.Get("bool_val").asBool()
	assertEqual(t, o, true)
}
func TestBool_NonExistantKey_NoDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("bool_val_not_there").asBool()
	assertNotNil(t, err)
	assertEqual(t, o, false)
}
func TestBool_NonExistantKey_HasDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("bool_val_not_there").Default("True").asBool()
	assertNil(t, err)
	assertEqual(t, o, true)
}

func TestString_Simple(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, _ := f.Get("string_val").asString()
	assertEqual(t, o, "any string. only single line though.")
}
func TestString_NonExistantKey_NoDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("string_val_not_there").asString()
	assertNotNil(t, err)
	assertEqual(t, o, "")
}
func TestString_NonExistantKey_HasDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("string_val_not_there").Default("hello").asString()
	assertNil(t, err)
	assertEqual(t, o, "hello")
}
