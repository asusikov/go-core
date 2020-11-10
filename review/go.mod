module goondex

go 1.15

replace goondex/crawler v0.0.0 => ./pkg/crawler

replace goondex/index v0.0.0 => ./pkg/index

replace goondex/engine v0.0.0 => ./pkg/engine

replace goondex/engine/storage v0.0.0 => ./pkg/engine/storage

replace goondex/web v0.0.0 => ./pkg/web

require (
	goondex/crawler v0.0.0 // indirect
	goondex/engine v0.0.0 // indirect
	goondex/engine/storage v0.0.0 // indirect
	goondex/index v0.0.0 // indirect
	goondex/web v0.0.0 // indirect
)
