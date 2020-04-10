package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    fmt.Print("Input a number: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.Trim(input, "\n")
    num1, _ := strconv.Atoi(input)

    var start int = 1
    var stop int = num1
    var step int = 1

    fmt.Println("\nDefinite Loop:\n")
    for i := start; i < stop; i = i + step {
        if i % 2 == 0  {
            fmt.Println(i, "is even.")
        } else {
            fmt.Println(i, "is odd.")
        }
    }
    fmt.Println("\n---------------\n\nIndefinite Loop:")

    j := start
    for {
        if (j >= stop) {
            break
        }

        if j % 2 == 0  {
            fmt.Println(j, "is even.")
        } else {
            fmt.Println(j, "is odd.")
        }

        j += step
    }

    fmt.Println("\n---------------\n\nContinue keyword:")
    for k := -3; k <= 3; k++ {
        if k == 0 {
            fmt.Println("You can't divide by 0!")
            continue
        }
        fmt.Println("3 / ", k, "=", 3 / k)
    }
}
