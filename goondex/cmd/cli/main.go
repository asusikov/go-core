package main

import (
	"fmt"
	"log"

	"goondex/pkg/engine"
	"goondex/pkg/spider"
)

func main() {
	const url = "https://thinknetica.com/"
	data, err := spider.Scan(url, 3)
	if err != nil {
		log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
	}
	eng := engine.Engine{}
	eng.Index(data)
	res := eng.Search("Гарантия")
	for k, v := range res {
		fmt.Printf("Страница %s имеет адрес: %s\n", v, k)
	}
}
