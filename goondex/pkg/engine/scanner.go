package engine

// Позволяет сканировать сайты
type Scanner interface {
	Scan(string, int) (map[string]string, error)
}
