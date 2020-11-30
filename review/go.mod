module goondex

go 1.15

replace goondex/crawler v0.0.0 => ./pkg/crawler

replace goondex/crawler/site v0.0.0 => ./pkg/crawler/site

replace goondex/index v0.0.0 => ./pkg/index

replace goondex/engine v0.0.0 => ./pkg/engine

replace goondex/webpages v0.0.0 => ./pkg/webpages

replace goondex/webpages/storage v0.0.0 => ./pkg/webpages/storage

replace goondex/bitree => ./pkg/bitree

replace goondex/cli => ./cmd/cli

require (
	goondex/bitree v0.0.0-00010101000000-000000000000 // indirect
	goondex/cli v0.0.0-00010101000000-000000000000 // indirect
	goondex/crawler v0.0.0 // indirect
	goondex/crawler/site v0.0.0 // indirect
	goondex/engine v0.0.0 // indirect
	goondex/index v0.0.0 // indirect
	goondex/webpages v0.0.0 // indirect
	goondex/webpages/storage v0.0.0 // indirect
)
