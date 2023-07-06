package main

import (
	"fmt"
	"log"
	"net/http"
)

// w - responseWriter (куда писать ответ)
// r - request (откуда брать запрос)
// Обработчик
func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World - I'm new web server</h1>")
}

// Обработчик в зависимости от адреса, по которому пришел запрос
func ReuqestHandler() {
	http.HandleFunc("/", GetHello)               // Если придет запрос по адресу "/" то вызывай GetHello
	log.Fatal(http.ListenAndServe(":8080", nil)) // Запускаем веб сервер в режиме "слушания"
}

func main() {
	ReuqestHandler()
}
