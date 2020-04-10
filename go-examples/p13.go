package main

import (
    "fmt"
)

type person struct {
    name string
    age int
    spouse *person
    children []person
}

func newPerson() person {
    return person{"",0,nil,nil}
}

func (p *person) getMarried(p2 *person) {
    p.spouse = p2
}

func createChild(p *person, p2 *person, name string) person {
    child := newPerson()
    child.name = name
    p.children = append(p.children, child)
    p2.children = append(p2.children, child)

    return child
}

func (p *person) getOlderPtr() {
    (*p).age++
}

func (p person) getOlder() {
    p.age++
}

func main() {
    fmt.Println("Let's meet two people.\n")

    var bob person = person{"Bob", 21, nil, nil}
    tammy := person{"Tammy", 20, nil, nil}
    fmt.Println("We have", bob.name, "and",tammy.name,"\n")

    bob.getMarried(&tammy)
    tammy.getMarried(&bob)
    fmt.Println("Bob and Tammy got married...")
    fmt.Print(bob.name, "'s spouse's name is ", bob.spouse.name, ".\n" )
    fmt.Print(tammy.name, "'s spouse's name is ", tammy.spouse.name, ".\n" )
    fmt.Println("")

    timmy := createChild(&bob, &tammy, "Timmy")
    fmt.Print(bob.name," and ", tammy.name, " had a child named ", timmy.name, ".\n")
    fmt.Print(timmy.name, " is currently ",timmy.age," years old and has ", len(timmy.children), " children.\n")
}
