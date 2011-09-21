package configo

import "testing"

func assertEqual(t *testing.T, e, o interface{}) {
	if o != e {
		t.Errorf("expected %d, got %d", e, o)
	}
}

func TestInt_Simple(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	assertEqual(t, f.getInt("int_val", 0), 345)
}

func TestInt_NonExistentKey_UsesValue(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	assertEqual(t, f.getInt("int_val_not_there", 0), 0)
}

func TestBool_Simple(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	assertEqual(t, f.getBool("bool_val", false), true)
}

func TestString_Simple(t *testing.T) {
	f := NewConfigo("./test_data/test.conf")
	f.Load()
	assertEqual(t, f.getString("string_val", ""), "any string. only single line though.")
}
