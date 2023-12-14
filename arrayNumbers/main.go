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

const n = 15

//функция, создающая массив с случайными неупорядоченными числами

func arrayCreate() [n]int {
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))
	numbers := [n]int{}
	for i := 0; i < n; i++ {
		numbers[i] = random.Intn(100)
	}
	return numbers
}

//функция, которая осуществляет поиск пользовательского числа (если есть) и считающая количество чисел после него

func finder(array [n]int, userNumber int) int {
	foundNums := 0
	for i, x := range array {
		if x == userNumber {
			foundNums = n - (i + 1)
			break
		}
	}
	return foundNums
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
	foundNums := finder(currentArray, userNumber)

	if foundNums > 0 {
		fmt.Printf("Количество чисел после %d: %d\n", userNumber, foundNums)
	} else {
		fmt.Println("Ваше число не обнаружено в массиве.")
	}

}
