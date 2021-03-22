package task3_hash1

import "strconv"

//Add
//Delete
//InHash
//PrintValues

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
		str = "[" + key + "]" + v + "\n"
	}
}

func strtoint() {

}
