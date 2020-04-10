package main

import "fmt"

type person struct {
    name string
    age int
    spouse *person
    children []person
}

func getMarried(p *person, p2 *person) {
    p.spouse = p2
    p2.spouse = p
}

func createChild(p *person, p2 *person, name string) person {
    child := person{name, 0, nil, nil}
    p.children = append(p.children, child)
    p2.children = append(p2.children, child)

    return child
}

func setSpouse(p *person, p2 *person) {
    p.spouse = p2
}

func getOlderPtr(p *person) {
    (*p).age++
}

func getOlder(p person) {
    p.age++
}

func main() {
    var matt person = person{name: "Matt", age: 32, spouse: nil, children: nil}
    fmt.Println(matt, "\n")

    fmt.Print(matt.name, "'s age is ", matt.age, ". Some time passes . . .\n")
    getOlder(matt)
    fmt.Print(matt.name, "'s age is now ", matt.age, " years old?!\n")

    fmt.Println("\nCough... Let's try this again. Cough...\n")

    fmt.Print(matt.name, "'s age is ", matt.age, ". Some time passes . . .\n")
    getOlderPtr(&matt)
    fmt.Print(matt.name, "'s age is now ", matt.age," years old!\n")

    fmt.Println("\n----------------\n")

    fmt.Println("Let's meet two people.\n")

    var bob person = person{"Bob", 21, nil, nil}
    tammy := person{"Tammy", 20, nil, nil}
    fmt.Println("We have", bob.name, "and",tammy.name,"\n")

    getMarried(&bob, &tammy)
    fmt.Println("Bob and Tammy got married...")
    fmt.Print(bob.name, "'s spouse's name is ", bob.spouse.name, ".\n" )
    fmt.Print(tammy.name, "'s spouse's name is ", tammy.spouse.name, ".\n" )
    fmt.Println("")

    timmy := createChild(&bob, &tammy, "Timmy")
    fmt.Print(bob.name," and ", tammy.name, " had a child named ", timmy.name, ".\n")
    fmt.Print(timmy.name, " is currently ",timmy.age," years old and has ", len(timmy.children), " children.\n")
}
