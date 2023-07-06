## Работа с JSON файлами  

### Структуризация
***Серилазция*** - процесс конвертации объекта в последовательгость байтов. 
***Десериализация*** - процесс конвертации последовательности байтов в объект.

Идея структуризованного подхода состоит в том, что мы заранее подготавливаем набор структур, с ***ЯВНО ПРОПИСАННЫМИ ПРАВИЛАМИ*** сериализации/десериализации полей объектов.

```
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Struct for representation total slice
// First Level ob JSON object Parsing
type Users struct {
	Users []User `json:"users"`
}

//Internal user representation
//Second level of object JSON parsin
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

//Socail block representation
//Third level of object parsing
type Social struct {
	Vkontakte string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
}

//Функция для распечатывания User
func PrintUser(u *User) {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Social. VK: %s and FB: %s\n", u.Social.Vkontakte, u.Social.Facebook)
}

//1. Рассмотрим процесс десериализации (то есть когда из последовательности в объект)
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
		fmt.Println("================================")
		PrintUser(&u)
	}
}

```

Идея структурированной сериализации/десериализации состоит в том, чтобы общаться с JSON объектами напрямую , через стыкову полей.

Для того, чтобы настроить стыкову полей нужно:
* Определить необходимые уровни объектности JSON (в нашем случае их 3)
* Для каждого уровня объектности подготовить свою структуру, учитывающую набор полей объекта.
```
type Person struct {
    Name string `json:"name"
}
```
* И все! Больше ничего делать не нужно. Остается только считать из файла и поместить в экземпляр!

### Сериализация
* Только один способ - структуризованный.
```
//1. Превратим профессора в последовательность байтов
	byteArr, err := json.MarshalIndent(prof1, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArr))
	err = os.WriteFile("output.json", byteArr, 0666) //-rw-rw-rw-
	if err != nil {
		log.Fatal(err)
	}
```

Сериализация - процесс перегона в байты. Поэтому на это этапе у нас на руках будет ```[]byte``` , который в последствии будет помещен в файл ```output.json```

***ВАЖНО*** : 0664/0666 - права доступа к файлу (в нашем случае это ```rw```)



