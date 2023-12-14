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

const n = 12

//функция, создающая и сортирующая массив по возрастанию с случайными числами

func arrayCreate() [n]int {
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))
	numbers := [n]int{}
	for i := 0; i < n; i++ {
		numbers[i] = random.Intn(10)
	}
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}

//функция, которая определяет индекс, под которым впервые появляется пользовательское число в отсортированном массиве

func indexFinder(array [n]int, userNumber int) int {
	foundIndex := -1
	for i, x := range array {
		if x == userNumber {
			foundIndex = i
			break
		}
	}
	return foundIndex
}

func main() {
	userNumber := 0
	fmt.Println("Введите число:")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		userNumber, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите корректное целое число.")
			continue
		} else {
			break
		}
	}

	currentArray := arrayCreate()
	fmt.Println("Ваш массив:", currentArray)
	foundIndex := indexFinder(currentArray, userNumber)

	if foundIndex > -1 {
		fmt.Printf("Первое появление в массиве числа %d приходится на индекс: %d\n", userNumber, foundIndex)
	} else {
		fmt.Println("Ваше число не обнаружено в массиве.")
	}

}
