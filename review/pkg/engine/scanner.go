package engine

import "goondex/goondex"

// Позволяет сканировать сайты
type Scanner interface {
	Scan(string, int) ([]goondex.Page, error)
}
