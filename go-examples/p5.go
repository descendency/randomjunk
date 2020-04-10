package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main () {
    fmt.Print("Input a number: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.Trim(input, "\n")
    num1, _ := strconv.Atoi(input)

    /*
        There are 6 logical comparison operators in programming:
            - Equal to: ==
            - Not Equal to: !=
            - Less than: <
            - Greater than: >
            - Less than or equal to: <=
            - Greater than or equal to: >=
    */
    var positive bool = num1 > 0    // This is kind of rare to see, but some functions return boolean values.

    if positive {
        // only executes if true
        fmt.Println(num1, "is positive.")
    } else if num1 < 0 {
        // only executes if the first statement is false and this statement is true
        fmt.Println(num1, "is negative.")
    } else {
        // only executes if all if and if-else statements are false
        fmt.Println("You input zero.")
    }

    /*
        Along with the comparison opertors, there are 3 major ways to join boolean Types:
            && (and)
            || (or)
            ! (not)
    */

    fmt.Println("\n--------------\n")
    fmt.Println("And Operator:")
    fmt.Println("true && true = ", true && true)
    fmt.Println("true && false = ", true && false)
    fmt.Println("false && true = ", false && true)
    fmt.Print("false && false = ", false && false, "\n")
    fmt.Println("\n")

    fmt.Println("Or Operator:")
    fmt.Println("true || true = ", true || true)
    fmt.Println("true || false = ", true || false)
    fmt.Println("false || true = ", false || true)
    fmt.Println("false || false = ", false || false)
    fmt.Println("\n")

    fmt.Println("Not Operator:")
    fmt.Println("!true =", !true)
    fmt.Println("!false =", !false)
}
