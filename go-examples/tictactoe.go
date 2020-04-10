package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

type game struct {
    board [3][3]string              // the gameboard (this is a 3x3 string array)
    playerSymbol string             // player symbol represents which players turn it is
    turnCount int                   // turn count represents how many turns have happened, since 9 turns is a full board
    playerType map[string]bool      // This is for when I use human vs computer players
    rng *rand.Rand                  // The random number generator that I use for computer moves
    printingEnabled bool            // Whether or not the board should be printed.
}

/*
    Returns the result of a game.

    The game must have ended or the program will announce a draw.

    Could implement an error message if it is called without a game being over.
*/
func (g *game) getResult() string {
    if g.hasWon("X") {
        return "!!!Player X has won!!!"
    } else if g.hasWon("O") {
        return "!!!Player O has won!!!"
    } else {
        return "The game was a draw!"
    }
}

func (g game) hasWon(sym string) bool {
    for i := 0; i < 3; i++ {
        // horizontal
        if g.board[i][0] == sym && g.board[i][1] == sym && g.board[i][2] == sym {
            return true;
        }
        // vertical
        if g.board[0][i] == sym && g.board[1][i] == sym && g.board[2][0] == sym {
            return true;
        }
    }
    // backslash
    if g.board[0][0] == sym && g.board[1][1] == sym && g.board[2][2] == sym {
        return true;
    }

    //forward slash
    if g.board[0][2] == sym && g.board[1][1] == sym && g.board[2][0] == sym {
        return true;
    }
    return false;
}

func (g game) isOver() bool {
    return g.turnCount > 7 || g.hasWon(g.playerSymbol)
}

func (g *game) endTurn() {
    if (g.playerSymbol == "X") {
        g.playerSymbol = "O"
    } else {
        g.playerSymbol = "X"
    }

    g.turnCount++
}

func (g *game) makeMove(pos int, player string) {
    var col int = pos % 3
    var row int = pos / 3

    g.board[row][col] = player
}

func (g *game) init() {
    g.printingEnabled = true
    g.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
    g.playerType = map[string]bool{"X": true, "O": true}
    g.turnCount = 0
    g.playerSymbol = "X"
    for i := 0; i < 9; i++ {
        g.board[i / 3][i % 3] = " ";
    }
}

func (g *game) getMove() int {
    if g.playerType[g.playerSymbol] {
        var num1 int
        reader := bufio.NewReader(os.Stdin)
        for {
            fmt.Print("Select a square (" + g.playerSymbol + "): ")
            input1, _ := reader.ReadString('\n')
            input1 = strings.Trim(input1, "\n")
            num1, err := strconv.Atoi(input1)

            if (err != nil) {
                fmt.Println("Invalid square choice. Choose another square.")
                continue;
            }
            pos := num1 - 1

            var col int = pos % 3
            var row int = pos / 3

            if (pos < 0 || pos > 8) {
                fmt.Println("Index out of range. Choose another square.")
                continue;
            }

            if g.board[row][col] == " " {
                return pos
            } else {
                fmt.Println("Square already picked. Choose another square.")
            }
        }
        return (num1 - 1)
    } else {
        for {
            var pos int
            pos = g.rng.Intn(9)
            var col int = pos % 3
            var row int = pos / 3

            if g.board[row][col] == " " {
                return pos
            }
        }
        return 0;
    }
}

func (g game) printBoard() {
    if g.printingEnabled == false {
        return
    }
    for i, v := range g.board {
        fmt.Printf("%d|%d|%d\t\t\t\t",3 * i + 1, 3 * i + 2, 3 * i + 3)
        for i2, v2 := range v {
            fmt.Print(v2)
            if (i2 < 2) {
                fmt.Print("|")
            }
        }
        fmt.Println()
        if (i < 2) {
            fmt.Println("-----\t\t\t\t-----")
        }
    }
    fmt.Println()
}

func main() {
    var output string
    g := game{}
    g.init()
    g.playerType["X"] = true
    g.playerType["O"] = true
    g.printingEnabled = true
    // Game Loop
    for {
        g.printBoard()
        var position int
        position = g.getMove()
        g.makeMove(position, g.playerSymbol)
        if g.isOver() {
            break;
        }
        g.endTurn()
    }
    g.printBoard()
    output = g.getResult()

    if g.printingEnabled {
        fmt.Println(output)
    }
}
