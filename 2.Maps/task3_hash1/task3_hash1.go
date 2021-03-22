package task3_hash1

import (
	"fmt"
	"strconv"
)

//Add
//Delete
//InHash
//PrintValues

func Run() {
	m := make(map[string]int64)
	//
	FillMap(m)
	// print / getlist

	fmt.Println("> Get list:")
	fmt.Println(GetList(m))

	fmt.Println("> Check in hash:")
	fmt.Println("Visa in hash: ", InHash(m, "Visa"))

	fmt.Println("> Del Visa:")
	Del(m, "Visa")

	fmt.Println("> Get list2:")
	fmt.Println(GetList(m))
}

func FillMap(m map[string]int64) map[string]int64 {
	Add(m, "American Express", 378282246310005)
	Add(m, "Master Card", 5105105105105100)
	Add(m, "Visa", 4222222222222)
	Add(m, "JCB", 3530111333300000)
	return m
}

func Add(m map[string]int64, key string, v int64) {
	m[key] = v
}

func Del(m map[string]int64, key string) {
	delete(m, key)
}

func InHash(m map[string]int64, key string) bool {
	_, ok := m[key]
	return ok
}

func GetList(m map[string]int64) string {
	var str string
	for key, v := range m {
		str += "[" + key + "] -> " + strconv.Itoa(int(v)) + "\n"
	}

	return str
}
