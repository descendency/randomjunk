package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
    _ "strconv"
)

func printResults(results string) {
    fmt.Println(results)
}

func parseCommand(input string) (string, []string, error) {

    commandFields := strings.Split(inputCommand, " ")

    command := commandFields[0]

    if len(command) == 0 {
        return "","", errors.New("ERROR: No command input.")
    }

    if len(commandFields) > 1 {
        arguments := commandFields[1:]
    }

    return command, arguments, nil
}

func execute(inputCommand string) (string, error) {
    results := ""


    command, arguments, err := commandFields[0]


    fmt.Println(command, arguments, err)

    return results, nil
}

func getInput(prompt string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)
    input, _ := reader.ReadString('\n')
    input = strings.Trim(input, "\n")

    return input
}

func main() {
    // REPL Loop
    for {
        // Read (input from user)
        input := getInput("> ")

        // Execute (the input command)
        results, err := execute(input)

        // Print (the results of the execution)
        if (err != nil) {
            //
            fmt.Println(err.Error())

            // because there are no results, skip printing them.
            continue
        }
        printResults(results)
    } // Loop back to the top.
}
