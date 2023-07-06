package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Struct for representation total slice
// First Level ob JSON object Parsing
type Users struct {
	Users []User `json:"users"`
}

// Internal user representation
// Second level of object JSON parsin
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

// Socail block representation
// Third level of object parsing
type Social struct {
	Vkontakte string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

// Функция для распечатывания User
func PrintUser(u *User) {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Social. VK: %s\n", u.Social.Vkontakte)
	fmt.Printf("Social. FB: %s\n", u.Social.Facebook)
}

// Сериализация
type Professor struct {
	Name       string     `json:"name"`
	ScienceID  int        `json:"science_id"`
	IsWorking  bool       `json:"is_working"`
	University University `json:"university"`
}

type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

// 1. Рассмотрим процесс десериализации (то есть когда из последовательности в объект)
func main() {
	//1. Создадим файл дескриптор
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	fmt.Println("File descriptor successfully created!")

	//2. Теперь десериализуем содержимое jsonFile в экземпляр Go
	// Инициализируем экземпляр Users
	var users Users

	// Вычитываем содержимое jsonFile в ВИДЕ ПОСЛЕДОВАТЕЛЬНОСТИ БАЙТ!
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// Теперь задача - перенести все из byteValue в users - это и есть десериализация!
	json.Unmarshal(byteValue, &users)
	for _, u := range users.Users {
		fmt.Println("==============")
		PrintUser(&u)
	}

	//Сериализация
	fmt.Println("=====Сериализация======")
	prof1 := Professor{
		Name:      "Stefan",
		ScienceID: 1111111,
		IsWorking: true,
		University: University{
			Name: "MGU",
			City: "Moscow",
		},
	}

	//Превратим профессора в последовательность байтов
	byteArr, err := json.MarshalIndent(prof1, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArr))
	err = os.WriteFile("output.json", byteArr, 0666) //-rw-rw-rw-
	if err != nil {
		log.Fatal(err)
	}

}
