module goondex

go 1.15

replace goondex/crawler v0.0.0 => ./pkg/crawler

replace goondex/index v0.0.0 => ./pkg/index

replace goondex/engine v0.0.0 => ./pkg/engine

replace goondex/goondex v0.0.0 => ./pkg/goondex

require (
	goondex/crawler v0.0.0 // indirect
	goondex/engine v0.0.0 // indirect
	goondex/goondex v0.0.0 // indirect
	goondex/index v0.0.0 // indirect
)
