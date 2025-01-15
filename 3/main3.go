package main

import "fmt"

func main() {
	new := NewMap()
	new.Add("qwe", 15)
	new.Add("qwerty", 153321)
	fmt.Println("map after add:", new)
	new.Delete("qwe")
	fmt.Println("map after delete:", new)
	copyMap := new.Copy()
	fmt.Println("copyMap:", copyMap)
	exist := new.Exist("qwerty")
	fmt.Println("exist:", exist)
	v, ok := copyMap.Get("qwerty")
	fmt.Printf("v=%v, ok=%v\n", v, ok)

}

type StringIntMap struct {
	strIntMap map[string]int
}

func NewMap() *StringIntMap {
	return &StringIntMap{
		strIntMap: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.strIntMap[key] = value
}

func (m *StringIntMap) Delete(key string) {
	delete(m.strIntMap, key)
}

func (m *StringIntMap) Copy() *StringIntMap {
	newMap := NewMap()
	for key, value := range m.strIntMap {
		newMap.strIntMap[key] = value
	}
	return newMap
}

func (m *StringIntMap) Exist(key string) bool {
	_, exists := m.strIntMap[key]
	return exists
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, ok := m.strIntMap[key]
	return value, ok
}
