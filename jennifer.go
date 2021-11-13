// Package jennifer is a code generator for Go
package jennifer

//go:generate go get github.com/pedidopago/jennifer/genjen
//go:generate genjen
//go:generate go get github.com/pedidopago/jennifer/gennames
//go:generate gennames -output "jen/hints.go" -package "jen" -name "standardLibraryHints" -standard -novendor -path "./..."
//go:generate go get github.com/dave/rebecca/cmd/becca
//go:generate becca -package=github.com/pedidopago/jennifer/jen
