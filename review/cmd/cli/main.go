package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"goondex/pkg/engine"
)

func main() {
	const depth = 2

	var url string
	flag.StringVar(&url, "url", "https://bash.im", "Адрес сайта для сканирования")
	var query string
	flag.StringVar(&query, "q", "", "Строка для поиска")
	var isInteractive bool
	flag.BoolVar(&isInteractive, "i", false, "Интерактивный режим")
	flag.Parse()

	eng := engine.New()
	fmt.Println("[Сканирование] Страница -", url)
	err := eng.Scan(url)
	if err != nil {
		log.Fatalf("ошибка при сканировании сайта %s: %v\n", url, err)
	}

	runLoop := true
	runSearch := false

	for runLoop {
		if isInteractive {
			query, runSearch = interactiveInput()
		}

		if !runSearch {
			fmt.Println("[Поиск] Запрос -", query)
			res := eng.Search(query)
			printResult(res)
		}
		runLoop = isInteractive && !runSearch
	}
}

func printResult(result map[string]string) {
	fmt.Println("[Результат]")
	for k, v := range result {
		fmt.Printf("Страница \"%s\" имеет адрес: %s\n", v, k)
	}
}

func interactiveInput() (input string, isExitCommand bool) {
	const exitCommand = "/q"

	fmt.Println("[Ввод запроса] Введите запрос или /q для выхода:")
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	return input, input == exitCommand
}
