module go-core

go 1.15

replace goondex/pkg/spider v0.0.0 => ./goondex/pkg/spider

replace goondex/pkg/engine v0.0.0 => ./goondex/pkg/engine

require (
	goondex/pkg/engine v0.0.0 // indirect
	goondex/pkg/spider v0.0.0 // indirect
)
