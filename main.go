package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	row, col int
}

func main() {
	var width, height, i, startI, startJ, currentI, currentJ int
	var matriz [][]string
	var fields []string
	var haveStart bool
	var currentCoord string

	lineNumber := 1

	file, err := os.Open("./maze.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the first line
	for scanner.Scan() {

		if lineNumber == 1 {
			firstLine := scanner.Text()

			// Split the line by whitespace
			fields := strings.Fields(firstLine)

			// Parse the values as integers
			width, err = strconv.Atoi(fields[0])
			if err != nil {
				log.Fatal(err)
			}

			height, err = strconv.Atoi(fields[1])
			if err != nil {
				log.Fatal(err)
			}

			matriz = make([][]string, width)

		} else {
			matriz[i] = make([]string, height)
			for j := 0; j < height; j++ {
				line := scanner.Text()
				fields = strings.Fields(line)
				value := fields[j]
				if err != nil {
					log.Fatal(err)
				}
				matriz[i][j] = fmt.Sprint(value)
			}
			i++
		}
		lineNumber++
	}

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			if matriz[i][j] == "X" {
				startI = i
				startJ = j
				haveStart = true
				currentCoord := fmt.Sprint("[", startI, startJ, "]")
				fmt.Println("O " + currentCoord)
				break
			}
		}
	}
	if !haveStart {
		log.Fatal("O labirinto não tem inicio")
	}

	currentI = startI
	currentJ = startJ
	visitou := ""
	for i = 0; matriz[currentI][currentJ] != "O"; i++ {

		if (currentI-1 >= 0 || currentI+1 <= height || currentJ+1 <= width || currentJ-1 >= 0) && (matriz[currentI][currentJ] != "1") {
			if currentI-1 > 0 && matriz[currentI-1][currentJ] != "1" && (!strings.HasSuffix(visitou, "BB") && !strings.HasSuffix(visitou, "CB")) {
				currentCoord += fmt.Sprint("C [", currentI-1, currentJ, "]\n")
				currentI -= 1
				visitou += "C"
			} else if currentJ+1 < width && matriz[currentI][currentJ+1] != "1" && (!strings.HasSuffix(visitou, "EE") || !strings.HasSuffix(visitou, "DE")) {
				currentCoord += fmt.Sprint("D [", currentI, currentJ+1, "]\n")
				currentJ += 1
				visitou += "D"
			} else if currentJ-1 > 0 && matriz[currentI][currentJ-1] != "1" && (!strings.HasSuffix(visitou, "DD") || !strings.HasSuffix(visitou, "ED")) {
				currentCoord += fmt.Sprint("E [", currentI, currentJ-1, "]\n")
				currentJ -= 1
				visitou += "E"
			} else if currentI+1 < height && matriz[currentI+1][currentJ] != "1" && (!strings.HasSuffix(visitou, "CC") || !strings.HasSuffix(visitou, "BC")) {
				currentCoord += fmt.Sprint("B [", currentI+1, currentJ, "]\n")
				currentI += 1
				visitou += "B"
			} else if currentI-1 >= 0 || currentI+1 <= height || currentJ+1 <= width || currentJ-1 >= 0 {
				fmt.Println("Labirinto resolvido!")
				break
			} else {
				fmt.Println("Não é possivel resolver o labirinto")
				break
			}
		}

		fmt.Println(currentCoord)
	}
}
