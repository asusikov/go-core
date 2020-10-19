package main

import (
	"fmt"
	"log"

	"goondex/pkg/engine"
	"goondex/pkg/spider"
)

func main() {
	const url = "https://thinknetica.com/"
	const query = "Гарантия"

	fmt.Println("[Сканирование] Страница -", url)
	data, err := spider.Scan(url, 3)
	if err != nil {
		log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
	}

	eng := engine.New()
	eng.Index(data)
	fmt.Println("[Поиск] Запрос -", query)
	res := eng.Search(query)

	fmt.Println("[Результат]")
	for k, v := range res {
		fmt.Printf("Страница \"%s\" имеет адрес: %s\n", v, k)
	}
}
