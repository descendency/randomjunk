package main

import "fmt"

func main () {

    var output string
    var three int
    three = 3

    var two int = 2

    five := 5
    six := three * two

    var fivePointFive float64
    fivePointFive = 5.5
    threePointNine = 3.9

    const constantTwo int = 2

    fmt.PrintLn("3 + 2 = ", three + two)
    fmt.Println("5 / 2 = ", five / constantTwo)
    //fmt.Println("5.5 * 3 = ", fivePointFive * three) - this does not work.
    fmt.Println("5.5 * 3.9 = ", fivePointFive * threePointNine)
    fmt.Println(three, " * ", two, " = ", six)

    output = "Hello, " + "World!"

    fmt.Println(output)
}
