package main

import (
	"log"

	"github.com/babaevs2023/GO/GO_lvl2/lec11/models"
	"github.com/babaevs2023/GO/GO_lvl2/lec11/routers"
	"github.com/babaevs2023/GO/GO_lvl2/lec11/storage"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var err error

func main() {
	//go get -u girhub.com/jinzhu/gorm/<dialects>

	storage.DB, err = gorm.Open("postgres", "host=localhost dbname=restapi port=5432 user=postgres password=1 sslmode=disable")
	if err != nil {
		log.Println("error while accessing database:", err)
	}
	defer storage.DB.Close()                  // не забудем закрыть соединение
	storage.DB.AutoMigrate(&models.Article{}) // в этот момент орм сама сгенерит все запросы, миграции и их применит

	r := routers.SetupRouter()

	// r - gin маршрутизатор
	r.Run()
}
