module goondex/cmd/cli

go 1.15

replace goondex/pkg/spider v0.0.0 => ../../pkg/spider

replace goondex/pkg/engine v0.0.0 => ../../pkg/engine

require (
	golang.org/x/net v0.0.0-20201016165138-7b1cca2348c0 // indirect
	goondex/pkg/engine v0.0.0 // indirect
	goondex/pkg/spider v0.0.0 // indirect
)
