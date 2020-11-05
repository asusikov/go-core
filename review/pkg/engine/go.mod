module engine

go 1.15

replace spider v0.0.0 => ../spider

replace goondex v0.0.0 => ../goondex

require (
	golang.org/x/net v0.0.0-20201016165138-7b1cca2348c0 // indirect
	goondex v0.0.0
	spider v0.0.0
)
