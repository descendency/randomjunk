package main

import "bufio"
import "fmt"
import "os"
import "strings"
import "strconv"

func main() {

    fmt.Print("Input a number: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.Trim(input, "\n")
    num1, _ := strconv.Atoi(input)
    fmt.Print("Input a second number: ")
    input, _ = reader.ReadString('\n')
    input = strings.Trim(input, "\n")
    num2, _ := strconv.Atoi(input)

    sum := num1 + num2

    fmt.Println(num1, " + ", num2, " = ", sum)
}
