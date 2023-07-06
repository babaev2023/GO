## Стандартный веб-сервер

### Инициализация go mod
```
go mod init github.com/babaev2023/GO/GO_lvl2/lec6
```

### Cтандартные шаблоны
**Полезная ссылка**: https://github.com/golang-standards/project-layout (здесь можно найти советы по стрктурированию/пакетированию/рефакторингу любого Go-приложения)

### Создадим входную точку в приложение
Стандартный шаблон входной точки :
```
cmd/<app_name>/main.go
```
Мы создадим :
```
cmd/api/main.go
```

### Инициализация ядра сервера
Стандартным шаблоном диктуется следующая схема
```
internal/app/<app_name>/<app_name>.go
```
У нас будет ```internal/app/api/api.go```

### Важный пункт про конфигурацию
**Правило**: в Go принято, что:
* конфигурация всегда хранится в сторонних файлах (.toml, .env) 
* в Go проектах ВСЕГДА присутствуют дефолтные настройки (исключение - бд стараются не дефолтить)

### Конфигурирование API сервера
Базово, для конфигурация нужен лишь порт.
```
intrenal/app/api/config.go
```

### Создадим конфиги 
```
configs/<app_name>.toml или configs/.env
```

```
//api.toml
bind_addr = ":8080"


go get github.com/BurntSushi/toml
```

### Как конфиги передавать?
Хочется запускать :
```
api.exe -path configs/api.toml
```

### Конфигурация http сервера
```
go get -u github.com/gorilla/mux
```

'''
go build -v ./cmd/api/
'''

### Loger
'''
go get github.com/sirupsen/logrus
'''

