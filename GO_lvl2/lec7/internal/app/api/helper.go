package api

import (
	"net/http"

	"github.com/babaev2023/GO/GO_lvl2/lec7/storage"
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Пытаемся откунфигурировать наш API инстанс (а конкретнее - поле logger)
func (a *API) configreLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// Пытаемся отконфигурировать маршрутизатор (а конкретнее поле router API)
func (a *API) configreRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is rest api!"))
	})
}

// Пытаемся отконфигурировать наше хранилище (storage API)
func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)
	//Пытаемся установить соединениение, если невозможно - возвращаем ошибку!
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
