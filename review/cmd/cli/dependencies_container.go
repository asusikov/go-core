package main

import (
	"goondex/engine"
	"goondex/index"
	"goondex/webpages/storage"
)

type DependenciesContainer struct {
	urls    []string
	storage storage.Interface
	index   index.Interface
	engine  *engine.Engine
}

func (dc *DependenciesContainer) init() {
	dc.urls = []string{
		"https://bash.im",
		"https://go.dev",
	}
	dc.storage = storage.New()
	dc.index = index.New()
	dc.engine = engine.New(dc.index, dc.storage)
}

func InitContainer() *DependenciesContainer {
	dc := &DependenciesContainer{}
	dc.init()
	return dc
}
