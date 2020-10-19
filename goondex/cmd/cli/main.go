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

	data, err := spider.Scan(url, 3)
	if err != nil {
		log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
	}

	eng := engine.New()
	eng.Index(data)
	res := eng.Search(query)

	for k, v := range res {
		fmt.Printf("Страница \"%s\" имеет адрес: %s\n", v, k)
	}
}
