module goondex

go 1.15

replace goondex/spider v0.0.0 => ./pkg/spider

replace goondex/index v0.0.0 => ./pkg/index

replace goondex/engine v0.0.0 => ./pkg/engine

replace goondex/goondex v0.0.0 => ./pkg/goondex

require (
	goondex/engine v0.0.0 // indirect
	goondex/index v0.0.0 // indirect
)
