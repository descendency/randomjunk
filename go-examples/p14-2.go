package main

import "fmt"

func PrintFunction(a ...interface{}) int {
    sum := 0
    for _, val := range a {
        temp, _ := fmt.Print(val)
        sum += temp
    }

    return sum
}

func main() {
    PrintFunction("Hello", 1, true, "End\n")
}
