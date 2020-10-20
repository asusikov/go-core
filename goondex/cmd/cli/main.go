package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"goondex/pkg/engine"
	"goondex/pkg/spider"
)

func main() {
	// const url = "https://thinknetica.com/"
	// const query = "Гарантия"
	const depth = 3

	url := flag.String("url", "", "Адрес сайта для сканирования")
	query := flag.String("q", "", "Строка для поиска")
	isInteractive := flag.Bool("i", false, "Интерактивный режим")
	flag.Parse()

	fmt.Println("[Сканирование] Страница -", *url)
	data, err := spider.Scan(*url, depth)
	if err != nil {
		log.Printf("ошибка при сканировании сайта %s: %v\n", *url, err)
	}

	eng := engine.New()
	eng.Index(data)

	if *isInteractive {
		runInteractive(eng)
	} else {
		search(eng, *query)
	}
}

func search(eng *engine.Engine, query string) {
	fmt.Println("[Поиск] Запрос -", query)
	res := eng.Search(query)

	fmt.Println("[Результат]")
	for k, v := range res {
		fmt.Printf("Страница \"%s\" имеет адрес: %s\n", v, k)
	}
}

func runInteractive(eng *engine.Engine) {
	const exitCommand = "/q"

	for {
		fmt.Println("[Ввод запроса] Введите запрос или /q для выхода:")
		fmt.Print("> ")

		reader := bufio.NewReader(os.Stdin)
		query, _ := reader.ReadString('\n')
		query = strings.Replace(query, "\n", "", -1)

		if query == exitCommand {
			return
		}
		search(eng, query)
	}
}
