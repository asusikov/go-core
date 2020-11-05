module cli

go 1.15

replace spider v0.0.0 => ../../pkg/spider

replace engine v0.0.0 => ../../pkg/engine

replace goondex v0.0.0 => ../../pkg/goondex

require (
	golang.org/x/net v0.0.0-20201016165138-7b1cca2348c0 // indirect
	engine v0.0.0 // indirect
	goondex v0.0.0 // indirect
	spider v0.0.0 // indirect
)
