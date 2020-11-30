package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"goondex/crawler"
	"goondex/engine"
	"goondex/index"
	"goondex/webpages"
	"goondex/webpages/storage"
)

var urls = []string{
	"https://bash.im",
	"https://go.dev",
}

func main() {
	storage := storage.New()
	index := index.New()
	eng := engine.New(index, storage)

	err := scan(urls, storage, index)
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
			res, err := eng.Search(query)
			if err == nil {
				printResult(res)
			} else {
				fmt.Printf("[Ошибка] Поиск вернул ошибку %v\n", err)
			}
		}
	}
}

func printResult(result []webpages.Page) {
	fmt.Println("[Результат]")
	for index, page := range result {
		fmt.Printf("%d. %s\n", index, page.URL)
		fmt.Println(page.Title)
	}
}

func scan(urls []string, storage storage.Interface, index index.Interface) error {
	const depth = 2
	for _, url := range urls {
		fmt.Println("[Сканирование] Страница -", url)
		err := crawler.Scan(url, depth, storage, index)
		if err != nil {
			return fmt.Errorf("ошибка при сканировании сайта %s: %v", url, err)
		}
	}
	return nil
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
