package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputs, err := makeVerificationNo(read())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(inputs)
}

func read() []string {
	var inputs []string
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if input := scanner.Text(); input == "" {
			break
		}
		inputs = append(inputs, scanner.Text())
	}
	return inputs
}

func makeVerificationNo(inputs []string) ([]string, error) {
	for i, input := range inputs {
		digitSum := 0
		for digit, s := range input {
			num, err := strconv.Atoi(string(s))
			if err != nil {
				return nil, fmt.Errorf("fail: %w", err)
			}
			if digit%2 != 0 {
				digitSum += num
				continue
			}
			if len(strconv.Itoa(num*2)) == 1 {
				digitSum += num * 2
				continue
			}
			digitSum += sumDigits(num * 2)
		}

		inputs[i] = input + strconv.Itoa(getVerificationNo(digitSum))
	}
	return inputs, nil
}

func sumDigits(i int) int {
	s := strconv.Itoa(i)
	result := 0
	for _, char := range s {
		num, _ := strconv.Atoi(string(char))
		result += num
	}
	return result
}

func getVerificationNo(i int) int {
	if i%10 == 0 {
		return 0
	}
	return 10 - i%10
}
