package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func userNum(max, rightAnswer, answerLang int) {
	answer := 0
	if answerLang == 1 {
		fmt.Println("Введіть число від 0 до", max-1)
	} else {
		fmt.Println("Enter a number from 0 to", max-1)
	}
	fmt.Scanln(&answer)
	if rightAnswer == answer {
		if answerLang == 1 {
			fmt.Println("Вітаю Ви вгадали число!")
		} else {
			fmt.Println("Congratulations! You guessed the number!")
		}
		main()
		return
	} else if rightAnswer > answer {
		if answerLang == 1 {
			fmt.Println("Загадане число більше.")
		} else {
			fmt.Println("The guessed number is greater.")
		}
		userNum(max, rightAnswer, answerLang)
	} else {
		if answerLang == 1 {
			fmt.Println("Загадане число менше.")
		} else {
			fmt.Println("The guessed number is smaller.")
		}
		userNum(max, rightAnswer, answerLang)
	}
}
func userMax(answerLang int) int {
	maxN := 100
	if answerLang == 1 {
		fmt.Println("До якого числа загадати?")
	} else {
		fmt.Println("Until what number to guess?")
	}
	fmt.Scanln(&maxN)
	if maxN <= 0 {
		userMax(answerLang)
	}
	return maxN
}
func userMode(answerLang int) int {
	if answerLang == 1 {
		fmt.Println("Оберіть режим.")
		fmt.Println("1 - Бот загадує число, Вам треба вгадати.")
		fmt.Println("2 - Ви загадуєте число, боту треба вгадати.")
	} else {
		fmt.Println("Choose a mode.")
		fmt.Println("1 - The bot guesses the number, you have to guess it.")
		fmt.Println("2 - You guess the number, the bot must guess.")
	}
	selectMode := 1
	fmt.Scanln(&selectMode)
	if selectMode != 1 && selectMode != 2 {
		userMode(answerLang)
	}
	return selectMode
}
func checkAnswer(randN, answerLang int) int {
	answer := 3
	if answerLang == 1 {
		fmt.Println("Бот загадав:", randN)
		fmt.Println("1 - Так, воно дорівнює", randN)
		fmt.Println("2 - Ні, воно більше", randN)
		fmt.Println("3 - Ні, воно менше", randN)
	} else {
		fmt.Println("The bot guessed:", randN)
		fmt.Println("1 - Yes, it is", randN)
		fmt.Println("2 - No, it is greater than", randN)
		fmt.Println("3 - No, it is smaller than", randN)
	}
	fmt.Scanln(&answer)
	if answer != 1 && answer != 2 && answer != 3 {
		checkAnswer(randN, answerLang)
	}
	return answer
}
func randomaiser(min, max int) int {
	return rand.Intn(max-min) + min
}
func botNum(min, max, answerLang int) {
	randN := randomaiser(min, max)
	answer := checkAnswer(randN, answerLang)
	if answer == 1 {
		if answerLang == 1 {
			fmt.Println("Бот переміг.")
		} else {
			fmt.Println("The bot won.")
		}
		main()
	} else if answer == 2 {
		min = randN + 1
		botNum(min, max, answerLang)
	} else {
		max = randN
		botNum(min, max, answerLang)
	}
}
func memoryLang(answer int) {

}
func chooseLang() int {
	answer := 2
	_, err := os.Stat("lang.txt")
	isExist := !os.IsNotExist(err)
	if isExist {
		data, err := os.ReadFile("lang.txt") // Читаємо весь файл
		if err != nil {
			fmt.Println("Error 2 (reading file)", err)
			return answer
		}
		answer, _ = strconv.Atoi(string(data))
		return answer
	} else {
		fmt.Println("Choose language")
		fmt.Println("1 - Ukrainian")
		fmt.Println("2 - English")
		fmt.Scanln(&answer)
		if answer != 1 && answer != 2 {
			chooseLang()
		}
		file, err := os.Create("lang.txt")
		if err != nil {
			fmt.Println("Error 1 (create a file)")
			return answer
		}
		defer file.Close()
		_, err = file.WriteString(strconv.Itoa(answer))
	}
	return answer
}
func main() {
	fmt.Println("Made by adattoweb")
	answerLang := chooseLang()
	if answerLang == 1 {
		fmt.Println("Вас вітає RandGame. Вам потрібно буде вгадати число, або загадати самим.")
	} else {
		fmt.Println("RandGame welcomes you. You will need to guess the number or guess yourself.")
	}
	selectMode := userMode(answerLang)
	maxN := userMax(answerLang) + 1
	if selectMode == 1 {
		randN := rand.Intn(maxN)
		userNum(maxN, randN, answerLang)
	} else {
		botNum(0, maxN, answerLang)
	}
}

// Made by adattoweb 20.02.2025 11:29
