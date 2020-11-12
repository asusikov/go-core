package crawler

import "goondex/web"

type Interface interface {
	Scan(string, int) ([]web.Page, error)
}
