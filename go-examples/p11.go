package main

import "fmt"

func setZeroRef(i int) {
    i = 0
}

func setZeroPtr(i *int) {
    *i = 0
}

func main() {
    var one int = 1

    fmt.Println("One:", one)

    setZeroRef(one)
    fmt.Println("One:", one)

    setZeroPtr(&one)
    fmt.Println("One:", one)

    *(&one) = 1
    fmt.Println("One:", one)
}
