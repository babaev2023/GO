## Напишите код, который будет сериализовывать структуру в json-строку следующего вида:
{
  "Имя": "имя",
  "Почта": "почта"
} 
# Для сериализации используется функция json.Marshal() пакета json. 
# структура:
type Person struct {
    Name        string
    Email       string
    DateOfBirth time.Time
} 
