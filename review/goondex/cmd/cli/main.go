package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"goondex/crawler"
	"goondex/engine"
	"goondex/engine/storage"
	"goondex/index"
	"goondex/webpages"
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
			res, err := eng.Search(query)
			if err == nil {
				printResult(res)
			} else {
				fmt.Printf("[Ошибка] Поиск вернул ошибку %v\n", err)
			}
		}
	}
}

func initEngine(url string) (eng *engine.Engine, err error) {
	index := index.New()
	storage := storage.New()
	fmt.Println("[Сканирование] Страница -", url)
	err = scan(url, storage, index)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сканировании сайта %s: %v", url, err)
	}
	eng = engine.New(index, storage)
	return eng, nil
}

func printResult(result []webpages.Page) {
	fmt.Println("[Результат]")
	for index, page := range result {
		fmt.Printf("%d. %s\n", index, page.URL)
		fmt.Println(page.Title)
	}
}

func scan(url string, storage storage.Interface, index index.Interface) error {
	const depth = 2
	pages, err := crawler.Scan(url, depth)
	if err != nil {
		return err
	}
	for _, page := range pages {
		storage.Insert(page)
		index.Add(page)
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
