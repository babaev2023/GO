## Реверс-инжиниринг JSON
# Есть пример API-вызова в формате JSON:
{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
} 

## На входе есть строка с сырыми данными, требуется написать функцию её десериализации:
type Response struct {
    // поля с тегами
}

func ReadResponse(rawResp string) (Response, error) {
    // код десериализации
} 