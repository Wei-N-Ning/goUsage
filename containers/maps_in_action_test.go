// https://blog.golang.org/go-maps-in-action

package containers

import (
	"testing"
	"sync"
)


func TestDeclareMap_ExpectNil(t *testing.T) {
	var m map[string]int
	if nil != m {
		t.Error()
	}
}


func TestDeclareMap_ExpectZeroValueRead(t *testing.T) {
	var m map[string]int
	if 0 != m["asd"] {
		t.Error()
	}
	//m["asd"] = 1  // panic: assignment to entry in nil map
}


func TestInitializeEmptyMap_Make(t *testing.T) {
	m := make(map[string]int)
	m["asd"] = 1
	if 1 != m["asd"] {
		t.Error()
	}
}


func TestReadPattern_SingleValue(t *testing.T) {
	m := map[string]int {
		"asd": 1,
	}
	v := m["asd"]
	n := m["bsd"]
	if 1 != v && 0 != n {
		t.Error()
	}
}


func TestDeletePattern_SingleValue(t *testing.T) {
	m := map[string]int {
		"asd": 1,
	}
	delete(m, "asd")
	delete(m, "bsd")
	if 0 != m["asd"] {
		t.Error()
	}
}


func TestReadPattern_WithValidation(t *testing.T) {
	m := map[string]int {
		"asd": 1,
	}
	v, ok := m["asd"]
	if !ok && v != 1 {
		t.Error()
	}
	if _, hasKey := m["asd"]; !hasKey {
		t.Error()
	}
}


func TestIteration_KeyValuePairs(t *testing.T) {
	m := map[string]int {
		"asd": 1,
		"bsd": 2,
	}
	sum := 0
	for _, v := range m {
		sum += v
	}
	if 3 != sum {
		t.Error()
	}
}


func TestUseMapAsSet(t *testing.T) {
	unique := map[int]bool{}
	uniqueSum := 0
	for _, v := range []int{1, 2, 1, 1, 4, 5, 1} {
		if !unique[v] {
			unique[v] = true
			uniqueSum += v
		}
	}
	if 12 != uniqueSum {
		t.Error(unique)
	}
}


func TestMapOfSlices_AppendToNil(t *testing.T) {
	m := make(map[string][]int)
	for _, v := range []int{0, 11, 4, 133, 61, 45} {
		if v >= 50 {
			m["large"] = append(m["large"], v)
		}
	}
	if 1 != len(m) && 2 != len(m["large"]) {
		t.Error(m)
	}
}


func TestNestedMap(t *testing.T) {
	m := map[string]map[string]int {
		"ak104": {
			"damage": 30,
			"velocity": 800},
		"xm8": {
			"damage": 45,
			"velocity": 650,
		},
	}
	highestDamage := 0
	highestDamageWeapon := ""
	for weapon, stats := range m {
		if stats["damage"] > highestDamage {
			highestDamage = stats["damage"]
			highestDamageWeapon = weapon
		}
	}
	if "xm8" != highestDamageWeapon {
		t.Error()
	}
}


func TestUseStructAsKey(t *testing.T) {
	type Weapon struct {
		damage   int
		velocity int
	}
	m := map[string]Weapon {
		"ak104": {30, 800},
		"xm8": {45, 650},
	}
	highestDamage := 0
	highestDamageWeapon := ""
	for name, weapon := range m {
		if weapon.damage > highestDamage {
			highestDamageWeapon = name
		}
	}
	if "xm8" != highestDamageWeapon {
		t.Error()
	}
}


func TestMapMutexRWAccess(t *testing.T) {
	type ThreadSafe struct{
		sync.RWMutex
		m map[string]int
	}
	threadSafe := ThreadSafe{m: make(map[string]int)}

	// read
	threadSafe.RLock()
	var _ = threadSafe.m["something"]
	threadSafe.RUnlock()

	// write
	threadSafe.Lock()
	threadSafe.m["something"] = 1
	threadSafe.Unlock()
}




