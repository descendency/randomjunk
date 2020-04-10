package main

import "fmt"

func main() {
    /*
        A multiline comment
        I can write all I want. Including code:
        fmt.Print("Hello, World!\n")
        Note: This did not execute.
    */

    fmt.Print("3 + 2 = ", 3 + 2, "\n")

    fmt.Println("3 - 2 = ", 3 - 2)

    fmt.Println("5 / 2 = ", 5 / 2)
    fmt.Println("5 % 2 = ", 5 % 2)

    fmt.Println("3.2 * 4.7 = ", 3.2 * 4.7)

    fmt.Println("Hello, " + "World!")
}
