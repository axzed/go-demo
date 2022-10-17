package main

import (
	"fmt"
	"sync"
)

type SafeSet struct {
	sync.RWMutex
	// we only need the uniqueness of the map key
	// so the value in the map we set it to empty struct
	m map[string]struct{}
}

func NewSet() *SafeSet {
	return &SafeSet{
		m: make(map[string]struct{}),
	}
}

// Add is a method to add elem in set
func (ss *SafeSet) Add(elem string) {
	// before you react you should add lock to the op
	ss.Lock()
	defer ss.Unlock()
	ss.m[elem] = struct{}{}
}

// Del is a method to delete the elem in set
func (ss *SafeSet) Del(elem string) {
	// before you react you should add lock to the op
	ss.Lock()
	defer ss.Unlock()
	delete(ss.m, elem)
}

// IsExist is a method to judge the elem exsist or not
func (ss *SafeSet) IsExist(elem string) bool {
	// before you react you should add lock to the op
	ss.RLock()
	defer ss.RUnlock()
	_, ok := ss.m[elem]
	return ok
}

// PrintElement is a method to show the list of the set
func (ss *SafeSet) GetElemList() []string {
	// before you react you should add lock to the op
	ss.RLock()
	defer ss.RUnlock()
	res := make([]string, 0)
	for k := range ss.m {
		res = append(res, k)
	}
	return res
}

// Merge is a method to merge another set elem in the set now you in use
func (ss *SafeSet) Merge(set *SafeSet) {
	// before you react you should add lock to the op
	ss.Lock()
	defer ss.Unlock()
	keys := set.GetElemList()
	for _, k := range keys {
		ss.Add(k)
	}
}

func setAdd(set *SafeSet, n int) {
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("key_%d", i)
		set.Add(key)
	}
}

func main() {
	s1 := NewSet()
	setAdd(s1, 10)
	s := s1.GetElemList()
	for _, v := range s {
		fmt.Println(v)
	}
}
