package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Введите строку: ")
	str := readInput()

	strs := strings.Split(str, " ")

	fmt.Print(len(strs))

}

func readInput() string {

	buff := bufio.NewReader(os.Stdin)
	str, _ := buff.ReadString('\n')

	return str

}
