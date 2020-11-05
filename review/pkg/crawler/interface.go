package crawler

import "goondex/goondex"

type Interface interface {
	Scan(string, int) ([]goondex.Page, error)
}
