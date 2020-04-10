package main

import (
    "fmt"
)

func specialPrintln(s ...string) int {
    var count int = 0
    for _, v := range s {
        n, _ := fmt.Print(v + " ")
        count += n
    }

    fmt.Print("\n")
    return count + 1
}

func divideNumbers(a int, b int) (int, int) {
    return a / b, a % b
}

func addNumbers(a int, b int) int {
    return a + b
}

func printHelloName(name string) {
    fmt.Print("Hello, ", name,"!\n")
}

func printHelloWorld() {
    fmt.Println("Hello, World!")
}

func main() {
    printHelloWorld()
    printHelloName("Matt")
    sum := addNumbers(3,9)
    fmt.Println("3 + 9 =", sum)
    d, r := divideNumbers(13,3)
    fmt.Println("13 / 3 =", d)
    fmt.Println("13 % 3 =", r)

    specialPrintln("Hello,", "World!", "How", "Are", "You?")
}
