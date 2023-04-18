package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"` // - означает, что это поле не будет сериализовано
}

func main() {
	man := Person{
		Name:        "Ivan",
		Email:       "Ivan@google.com",
		DateOfBirth: time.Now(),
	}
	jsMan, err := json.Marshal(man)
	if err != nil {
		log.Fatalln("unable marshal to json")
	}
	fmt.Printf("Man %v", string(jsMan)) // Man {"Имя":"Ivan","Почта":"Ivan@google.com"}
}
