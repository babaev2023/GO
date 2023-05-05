package main

import (
	"log"
	"testing"
)

//Принято называть:
// * <script_name>_test.go
// * <package_name>_test.go

// Для того, чтобы тестировать функции (методы, стркутуры, интерфейсы и т.д.)
// На каждый юнит создаем по своей тестирующей функции (Test)
// Приянто каждую такую функию начинать со слова Test....
func TestAdd(t *testing.T) {
	//1-ый test-case
	if res := Add(10, 20); res != 30 {
		log.Fatalf("Add(10, 20) fail. expected %d, got %d\n", 30, res)
		// log.Fatal провоцирует завершение работы кода
	}

	if res := Add(30, 30); res != 60 {
		log.Fatalf("Add(30, 30) fail. expected %d, got %d\n", 60, res)
	}
}

func TestSub(t *testing.T) {
	if res := Sub(10, 20); res != 200 {
		log.Fatalf("Sub(10, 20) fail. expected %d, got %d\n", 200, res)
		// log.Fatal провоцирует завершение работы кода
	}

	if res := Sub(30, 30); res != 900 {
		log.Fatalf("Add(30, 30) fail. expected %d, got %d\n", 900, res)
	}
}

func TestMult(t *testing.T) {}

//Для запуска тестов используем команду go test
// Детально : go test -v

//Coverage (покрытие) - показывает сколько % кода покрыто модульными тестами
// go test -cover -v
// 75~80% coverage - этого бывает более чем достаточно!

// Все что начинается с слова Test - будет запущено командой go test
// В Go приянто, что создается 1 модуль с тестами на весь пакет (вне зависимости от количества модулей в нем)
// Обязательно посмотрите, как происходит связка с CI для Go (TravisCI/CircleCI)
