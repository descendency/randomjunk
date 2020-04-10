package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    var list []int
    sum := 0

    fmt.Print("Input numbers for this program to sum up.\n\n")

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter a number (negative number to stop):")
        input, _ := reader.ReadString('\n')
        input = strings.Trim(input, "\n")
        num, _ := strconv.Atoi(input)
        if (num < 0) {
            break
        }
        list = append(list, num)
    }

    for i := 0; i < len(list); i++ {
        sum += list[i]
    }

    for k, value := range list {
        if (k == len(list) - 1) {
            // the last index of an array/slice is always length - 1.
            fmt.Print(value, " = ")
        } else {
            fmt.Print(value, " + ")
        }
    }
    fmt.Println(sum)
}
