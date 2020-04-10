package main

import "testpackage"

func main() {
    testpackage.PrintMessage()
    // The line below does not work, because it has a lower case first letter.
    //testpackage.printMessage()

}
