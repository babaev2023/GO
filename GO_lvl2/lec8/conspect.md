Работа с моделью 

### Откат миграции
Для выполнения отката ```migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" down```