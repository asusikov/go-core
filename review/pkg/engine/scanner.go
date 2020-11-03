package engine

import "goondex/pkg/goondex"

// Позволяет сканировать сайты
type Scanner interface {
	Scan(string, int) ([]goondex.Page, error)
}
