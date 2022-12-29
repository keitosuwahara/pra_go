package main

import (
    "fmt"
    "./animals"
)

func main() {
    for i := 0; i < 5; i++ { 
        fmt.Println(i) 
    }
    fmt.Print("山田孝之")
    fmt.Println(animals.dog_feed(), animals.monkey_feed())
}