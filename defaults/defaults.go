package defaults

import (
	_ "embed"

	"github.com/TimothyStiles/poly/synthesis/codon"
)

//go:embed codon_tables/freqB.json
var ecoliCodonTable []byte

var (
	CodonTables = map[string]codon.Table{
		"Escherichia coli": codon.ParseCodonJSON(ecoliCodonTable),
	}
)
