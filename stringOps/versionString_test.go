package stringOps

import "testing"


func TestSemanticVersion_ToStr_ZeroValue(t *testing.T) {
	v := SemanticVersion{}
	if "0.0.0" != v.ToStr() {
		t.Error("not expected")
	}
	if v.IsValid() {
		t.Error("not expected")
	}
}


func TestSemanticVersion_ToStr_Init(t *testing.T) {
	v := SemanticVersion{1, 1, 1}
	if "1.1.1" != v.ToStr() {
		t.Error("not expected")
	}
	if !v.IsValid() {
		t.Error("not expected")
	}
}


func TestSemanticVersion_EqualTo(t *testing.T) {
	v1 := SemanticVersion{1, 1, 1}
	v2 := SemanticVersion{1, 1, 1}
	v3 := SemanticVersion{}
	if v1 != v2 {
		t.Error("not expected")
	}
	if v1 == v3 {
		t.Error("not expected")
	}
}


func TestSemanticVersion_ToInts(t *testing.T) {
	v := SemanticVersion{1, 1, 1}
	nums := v.ToInts()
	expected := [3]uint{1, 1, 1}
	if expected != nums {
		t.Error("not expected")
	}
}


