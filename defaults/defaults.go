package defaults

import (
	_ "embed"

	"github.com/TimothyStiles/poly/synthesis/codon"
)

//go:embed codon_tables/freqB.json
var ecoliCodonTable []byte

//go:embed vectors/pET21-PaqCI-no_rbs.txt
var pET21PaqCINoRBS string

var (
	CodonTables = map[string]codon.Table{
		"Escherichia coli": codon.ParseCodonJSON(ecoliCodonTable),
	}

	Vectors = map[string]string{
		"pET21-PaqCI-no_rbs": pET21PaqCINoRBS,
	}

	VectorInstructions = map[string]string{
		"pET21-PaqCI-no_rbs": "Sequences should be flanked with a vectorPrefix of `AACT` and a vectorSuffix of `GGATCC`",
	}
)
