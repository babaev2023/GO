## Напишите функцию сложения целых чисел AddInts(a, b int) int. Разместите файл с этой функцией в локальной файловой системе так, чтобы приведённый ниже код успешно отработал. Изменять приведённый код нельзя.

package main

import (
    "fmt"

    "toppackage/middlepackage/bottompackage/mathxxx"
)

func main() {
    if sum := mathxxx.AddInts(3, 2); sum != 5 {
        panic(fmt.Sprintf("sum must be equal 5; got %d", sum))
    }

    fmt.Println("Well done!")
}


## Решение:
C:\Program Files\Go\src\toppackage\middlepackage\bootompackege\mathxxx\math.go
# имя файла может быть другим) со следующим содержимым:
package mathxxx

func AddInts(a, b int) int {
    return a + b
}