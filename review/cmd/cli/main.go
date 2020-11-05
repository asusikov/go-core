package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"goondex/engine"
	"goondex/web"
)

func main() {
	const depth = 2

	var url string
	flag.StringVar(&url, "url", "https://bash.im", "Адрес сайта для сканирования")
	flag.Parse()

	eng, err := initEngine(url)
	if err != nil {
		log.Fatal(err)
	}

	// Флаг, который определяет выполнять ли поиск или нет
	runSearch := true
	query := ""

	for runSearch {
		query, runSearch = userInput()

		if runSearch {
			fmt.Println("[Поиск] Запрос -", query)
			res := eng.Search(query)
			printResult(res)
		}
	}
}

func initEngine(url string) (eng *engine.Engine, err error) {
	eng = engine.New()
	fmt.Println("[Сканирование] Страница -", url)
	err = eng.Scan(url)
	if err != nil {
		return eng, fmt.Errorf("ошибка при сканировании сайта %s: %v\n", url, err)
	}
	return eng, nil
}

func printResult(result []web.Page) {
	fmt.Println("[Результат]")
	for index, page := range result {
		fmt.Printf("%d. %s\n", index, page.Url)
		fmt.Println(page.Title)
	}
}

func userInput() (input string, runSearch bool) {
	const exitCommand = "/q"

	fmt.Println("[Ввод запроса] Введите запрос или /q для выхода:")
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	return input, input != exitCommand
}
