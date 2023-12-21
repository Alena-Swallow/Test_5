go
package main

import "fmt"

type MyStruct struct {
    Value int
}

func (m *MyStruct) ChangeValue(newValue int) {
    m.Value = newValue
}

func main() {
    instance := MyStruct{Value: 10}
    fmt.Println("Initial Value:", instance.Value)

    instance.ChangeValue(20)
    fmt.Println("Updated Value:", instance.Value)
}
