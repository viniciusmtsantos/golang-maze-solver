package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var width, height, i, startI, startJ, currentI, currentJ int
	var matriz [][]string
	var fields []string
	var haveStart bool
	var currentCoord string

	lineNumber := 1

	file, err := os.Open("./entrada-labirinto3.txt")
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
			height, err = strconv.Atoi(fields[0])
			if err != nil {
				log.Fatal(err)
			}

			width, err = strconv.Atoi(fields[1])
			if err != nil {
				log.Fatal(err)
			}

			matriz = make([][]string, height)

		} else {
			matriz[i] = make([]string, width)
			for j := 0; j < width; j++ {
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
				currentCoord = fmt.Sprintf("O [%d,%d]\n", startI+1, startJ+1)
				break
			}
		}
	}
	if !haveStart {
		log.Fatal("O labirinto não tem inicio")
	}

	currentI = startI
	currentJ = startJ
	visited := make(map[string]bool)
	deadLine := make(map[string]bool)
	visited[fmt.Sprintf("%d-%d", currentI, currentJ)] = true

	for {

		if (currentI == 0 || currentJ == 0) && matriz[currentI][currentJ] != "X" {
			break
		} else if (currentI+1 >= height || currentJ+1 >= width) && matriz[currentI][currentJ] != "X" {
			break
		}

		if (currentI-1 >= 0 || currentI+1 <= height || currentJ+1 <= width || currentJ-1 >= 0) && (matriz[currentI][currentJ] != "1") {
			if currentI-1 > 0 && matriz[currentI-1][currentJ] != "1" && (!visited[fmt.Sprintf("%d-%d", currentI-1, currentJ)] || deadLine[fmt.Sprintf("%d-%d", currentI, currentJ)]) {
				currentCoord += fmt.Sprintf("C [%d,%d]\n", currentI-1+1, currentJ+1)
				currentI -= 1
			} else if currentJ+1 <= width && matriz[currentI][currentJ+1] != "1" && (!visited[fmt.Sprintf("%d-%d", currentI, currentJ+1)] || deadLine[fmt.Sprintf("%d-%d", currentI, currentJ)]) {
				currentCoord += fmt.Sprintf("D [%d,%d]\n", currentI+1, currentJ+1+1)
				currentJ += 1
			} else if currentJ-1 >= 0 && matriz[currentI][currentJ-1] != "1" && (!visited[fmt.Sprintf("%d-%d", currentI, currentJ-1)] || deadLine[fmt.Sprintf("%d-%d", currentI, currentJ)]) {
				currentCoord += fmt.Sprintf("E [%d,%d]\n", currentI+1, currentJ-1+1)
				currentJ -= 1
			} else if currentI+1 <= height && matriz[currentI+1][currentJ] != "1" && (!visited[fmt.Sprintf("%d-%d", currentI+1, currentJ)] || deadLine[fmt.Sprintf("%d-%d", currentI, currentJ)]) {
				currentCoord += fmt.Sprintf("B [%d,%d]\n", currentI+1+1, currentJ+1)
				currentI += 1
			} else {
				deadLine[fmt.Sprintf("%d-%d", currentI, currentJ)] = true
				continue
			}
		}

		visited[fmt.Sprintf("%d-%d", currentI, currentJ)] = true

	}
	fmt.Println(currentCoord)

	fileToCreatte, err := os.Create("./resolvido.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileToCreatte.WriteString(currentCoord)
}
