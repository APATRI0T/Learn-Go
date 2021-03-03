package main

import (
	"fmt"
)

/*
Мапы
	инициализация
	функция добавление значения
	функция чтение
	проверка, есть там такая запись

	Стуктуры пользовательские
*/

/*
	Структура
		Пользователь
			id
			username
			age
			email

	Мапа
		user
		id

*/
type tUser struct {
	id       int
	name     string
	FullName string
	email    string
	age      int
}

func main() {

	var users = map[string]tUser{
		"serg":   {1, "serg", "Худояров Сергей Владимирович", "shudoyarov@gmail.com", 31},
		"pushka": {2, "pushka", "Пушкин Александр Сергеевич", "pupupu@gmail.com", 222},
		"lerm":   {3, "lerm", "Лермонтов Михаил Юрьевич", "lerma@gmail.com", 206},
	}

	fmt.Println("UserID=", GetUserId(users, "lermg"))
	fmt.Println("User:", users["lerm"])
	fmt.Println("====================")

	newUser := tUser{4, "gogol", "Гоголь как его там, по батюшке.. забыл", "gogol@mail.ru", 90}
	AddUser(users, newUser)
	fmt.Println("====================\n", users)

	var (
		tmpStruct tUser
		newMap    map[string]tUser
	)
	newMap = make(map[string]tUser)
	newMap2 := map[string]tUser{}
	tmpStruct = tUser{5, "evtushenko", "Евтушенко ляляля", "eve@mail.ru", 88}
	AddUser(newMap, tmpStruct)
	AddUser(newMap2, tmpStruct)
	fmt.Println("====================\n", newMap)
	fmt.Println("====================\n", newMap2)
}

func GetUserId(m map[string]tUser, UserName string) int {
	var bUserExist bool

	user, bUserExist := m[UserName]
	if bUserExist {
		return user.id
	} else {
		return -1
	}

}
func AddUser(m map[string]tUser, user tUser) int {
	// Узнать, есть ли пользователь с таким именем
	// Если есть, то вывести ошибку
	// Если нет, то добавить, и вернуть его ID

	_, UserExist := m[user.name]
	if UserExist {
		fmt.Println("user with username ", user.name, " already exist :( \n Can't add user.")
		return -1
	} else {
		m[user.name] = user
		return user.id
	}
}

/*
Создать мапы стркутур
Мапу через make
а структуру литералом


*/
