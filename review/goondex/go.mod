module goondex

go 1.15

replace goondex/crawler v0.0.0 => ./pkg/crawler

replace goondex/index v0.0.0 => ./pkg/index

replace goondex/engine v0.0.0 => ./pkg/engine

replace goondex/engine/storage v0.0.0 => ./pkg/engine/storage

replace goondex/webpages v0.0.0 => ./pkg/webpages

require (
	goondex/crawler v0.0.0 // indirect
	goondex/engine v0.0.0 // indirect
	goondex/engine/storage v0.0.0 // indirect
	goondex/index v0.0.0 // indirect
	goondex/webpages v0.0.0 // indirect
)
