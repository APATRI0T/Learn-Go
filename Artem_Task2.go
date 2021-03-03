package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	var HashTable = map[string]string{"md5hash": "value"}

	Value := "password"
	Res := AddHash(HashTable, GetMD5Hash(Value), Value)

	fmt.Println(Res)
	printMap(HashTable)
	fmt.Println("====================")
	fmt.Println("Get hash:5f4dcc3b5aa765d61d8327deb882cf99, value:", GetValueByHash(HashTable, "5f4dcc3b5aa765d61d8327deb882cf99"))
	fmt.Println("Get hash:SomeUnknownHash, value:", GetValueByHash(HashTable, "SomeUnknownHash"))
}

func GetValueByHash(hashtable map[string]string, key string) string {
	value := hashtable[key]
	return value
}

func AddHash(hashtable map[string]string, key string, value string) string {
	_, ok := hashtable[key]
	if ok {
		return "Already in hash"
	} else {
		hashtable[key] = value
		return "Hash added successfully"
	}
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func printMap(m map[string]string) {
	var maxLenKey int
	for k, _ := range m {
		if len(k) > maxLenKey {
			maxLenKey = len(k)
		}
	}

	for k, v := range m {
		fmt.Println(k + ": " + strings.Repeat(" ", maxLenKey-len(k)) + v)
	}
}
