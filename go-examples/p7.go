package main

import (
    "fmt"
)

func main() {
    var list []int = []int{3, 9, 1, 4, 2}

    fmt.Println("list[0] =", list[0])
    fmt.Println("list[3] =", list[3])
    fmt.Println("list[4] =", list[4])

    fmt.Println("list =", list)

    fmt.Println("list[0:3] =", list[0:3])
    fmt.Println("list[:3] =", list[:3])
    fmt.Println("list[3:5] =", list[3:5])
    fmt.Println("list[3:] =", list[3:])

    fmt.Println("list[:] =", list[:])
    /*
        This does not work:
        if (list[0:3] == list[:3]) {
            fmt.Println("They are equal.")
        }
    */

    fmt.Println("There are", len(list), "elements in list.")

    var list2 []int

    if list2 == nil {
        fmt.Println("list2 is empty. Length:", len(list2))
    }
    copy(list, list2)
    fmt.Println("list2 is no longer empty.", list2)

    var name string = "Matt"
    fmt.Println(name[0:1],"",name[1:3],"",name[3:4])

    var list3 [][]int = [][]int{[]int{0,1,2},[]int{3,4,5},[]int{6,7,8}}

    fmt.Println(list3)

    fmt.Println("list3[1][0] =", list3[1][0])
    list3[1][0] = 9
    fmt.Println("list3[1][0] =", list3[1][0])
}
