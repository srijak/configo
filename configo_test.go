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
	o, _ := f.Get("int_val").AsInt()
	assertEqual(t, o, 345)
}
func TestInt_NonExistantKey_NoDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("int_val_not_there").AsInt()
	assertNotNil(t, err)
	assertEqual(t, o, 0)
}
func TestInt_NonExistantKey_HasDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("int_val_not_there").Default("23").AsInt()
	assertNil(t, err)
	assertEqual(t, o, 23)
}

func TestBool_Simple(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, _ := f.Get("bool_val").AsBool()
	assertEqual(t, o, true)
}
func TestBool_NonExistantKey_NoDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("bool_val_not_there").AsBool()
	assertNotNil(t, err)
	assertEqual(t, o, false)
}
func TestBool_NonExistantKey_HasDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("bool_val_not_there").Default("True").AsBool()
	assertNil(t, err)
	assertEqual(t, o, true)
}

func TestString_Simple(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, _ := f.Get("string_val").AsString()
	assertEqual(t, o, "any string. only single line though.")
}
func TestString_NonExistantKey_NoDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("string_val_not_there").AsString()
	assertNotNil(t, err)
	assertEqual(t, o, "")
}
func TestString_NonExistantKey_HasDefault(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	o, err := f.Get("string_val_not_there").Default("hello").AsString()
	assertNil(t, err)
	assertEqual(t, o, "hello")
}

type TestConfigStruct struct {
	// Properties to be populated using 
	// the config file should be exportable
	Astring string
	Somenum int
	Abool   bool
}

func TestHydrate(t *testing.T) {
	tc := TestConfigStruct{}
	NewConfigo("./test_data/hydrate.conf").Hydrate(&tc)

	assertEqual(t, tc.Astring, "abra")
	assertEqual(t, tc.Somenum, 23)
	assertEqual(t, tc.Abool, true)
}
