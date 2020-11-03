module goondex/pkg/engine

go 1.15

replace goondex/pkg/spider v0.0.0 => ../spider

replace goondex/pkg/goondex v0.0.0 => ../goondex

require (
	golang.org/x/net v0.0.0-20201016165138-7b1cca2348c0 // indirect
	goondex/pkg/goondex v0.0.0
	goondex/pkg/spider v0.0.0
)
