package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type InputParams struct {
	Pos int
	Coords []int
	Matrix [][]string
}

func (i *InputParams) populateInput(b []byte, count int) {
	switch count {
		case 0:
			pos, err := strconv.Atoi(string(b))
			if err != nil {
				fmt.Printf("error in converting param to int: %v", err.Error())
				return
			}
			i.Pos = pos
		case 1:
			err := json.Unmarshal(b, &i.Coords)
			if err != nil {
				fmt.Printf("error in marshal: %v", err.Error())
				return
			}
		default:
			str := string(b)
			var row []string
			for _, c := range str {
				row = append(row, string(c))
			}
			i.Matrix = append(i.Matrix, row)
	}
}

func scanInput() map[int]*InputParams {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	i := &InputParams{}
	inputs := make(map[int]*InputParams)
	for scanner.Scan() {
		b := scanner.Bytes()
		if len(b) == 0 {
			// there is a repeat position, end scanning
			if _, ok := inputs[i.Pos]; ok {
				return inputs
			}
			inputs[i.Pos] = i
			count = 0
			i = &InputParams{}
		} else {
			i.populateInput(b, count)
			count++
		}
	}
	// append the final one, if the last wasn't a blank line
	if count != 0 {
		inputs[i.Pos] = i
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
	}
	return inputs
}

func processInput(input *InputParams) string {
	rows := len(input.Matrix) - 1
	cols := len(input.Matrix[0]) - 1

	col := input.Coords[0]
	rowPrime := input.Coords[1]

	// validations
	if col > cols || rowPrime > rows {
		fmt.Printf("failed validation \n")
		return ""
	}
	row := rows - rowPrime
	res := input.Matrix[row][col]
	return res
}

func main() {
	inputArr := scanInput()
	passArr := make([]string, len(inputArr))
	for _, inputParams := range inputArr {
		fmt.Printf("inputParams: %v \n", inputParams)
		output := processInput(inputParams)
		passArr[inputParams.Pos] = output
	}
	str := ""
	for _, c := range passArr {
		str += c
	}
	fmt.Printf("str: %v \n", str)
}
