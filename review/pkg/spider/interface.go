package spider

// Позволяет сканировать сайты
type Interface interface {
	Scan(string, int) (map[string]string, error)
}
