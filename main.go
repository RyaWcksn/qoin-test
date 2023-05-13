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

func main() {
    rand.Seed(time.Now().UnixNano()) 

    fmt.Print("Enter number of players: ")
    numPlayersStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    numPlayers, _ := strconv.Atoi(strings.TrimSpace(numPlayersStr))

    fmt.Print("Enter number of dice per player: ")
    numDiceStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    numDice, _ := strconv.Atoi(strings.TrimSpace(numDiceStr))

    players := make([]int, numPlayers) 
    dice := make([][]int, numPlayers)  
    for i := range dice {
        dice[i] = make([]int, numDice)
    }

    for {
        numActivePlayers := 0 
        for i := range dice {
            if len(dice[i]) > 0 {
                numActivePlayers++
            }
        }
        if numActivePlayers == 1 {
            break
        }

        for i := range dice {
            if len(dice[i]) > 0 {
                for j := range dice[i] {
                    dice[i][j] = rand.Intn(6) + 1 
                }
            }
        }

        for i := range dice {
            if len(dice[i]) > 0 {
                for j := 0; j < len(dice[i]); {
                    if dice[i][j] == 6 {
                        players[i]++
                        dice[i] = append(dice[i][:j], dice[i][j+1:]...) 
                    } else if dice[i][j] == 1 {
                        nextPlayer := (i + 1) % numPlayers
                        dice[nextPlayer] = append(dice[nextPlayer], 1) 
                        dice[i] = append(dice[i][:j], dice[i][j+1:]...) 
                    } else {
                        j++
                    }
                }
            }
        }
    }

    maxScore := 0
    maxPlayer := -1
    for i, score := range players {
        if score > maxScore {
            maxScore = score
            maxPlayer = i
        }
    }
    fmt.Printf("Player %d wins with a score of %d!\n", maxPlayer+1, maxScore)
}

