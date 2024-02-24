package main

import (
	"bufio"
	"fmt"
	"github.com/thoas/go-funk"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Join(makeVerificationNo(read()), "\n"))
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

func makeVerificationNo(inputs []string) []string {
	return funk.Map(inputs, func(input string) string {
		sumOfEachDigit := int(funk.Sum(
			funk.Map(strToMap(input), logic),
		))
		return input + strconv.Itoa(getVerificationNo(sumOfEachDigit))
	}).([]string)
}

func logic(k int, s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	if k%2 != 0 {
		return num
	}
	if len(strconv.Itoa(num*2)) == 1 {
		return num * 2
	}
	return sumDigits(num * 2)
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

func strToMap(s string) map[int]string {
	m := map[int]string{}
	for k, c := range s {
		m[k] = string(c)
	}
	return m
}
