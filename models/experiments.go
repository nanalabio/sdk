package models

type Experiment struct {
	Synthesize []SynthesisRequest `json:"synthesis_request"`
	Logs       []string           `json:"logs"`
	Error      string             `json:"error"`
}

func (e *Experiment) Println(str string) {
	e.Logs = append(e.Logs, sr)
}

func (e *Experiment) Error(str string) {
	e.Error = str
}

//*****************************************************************************
//
//                          Synthesis management
//
//
// Part type's sequence definitions are appended to the parts given to the
// synthesis request in all cases except in the case of PART_TYPE_NONE, which
// can be used to define your own custom part type. PART_TYPE_NONE does not get
// any sequenced appended to its ends.
//
// In the case that you want to synthesize a part and might use it in a multi
// gene construct later, use PART_TYPE_IDENTITY.
//
// === Linkers ===
// A |  CGAG-TACT   | Prefix linker.
// B |  CGCT-GTCT   | Suffix linker
//
// === PartTypes between linkers ===
// P |  TACA-AACT    | Promoter part type.
// R |  AACT-AATG    | RBS part type.
// C |  A[ATG]-*ATCC | CDS part type. Protein should always start with an ATG
//                      and will be appended with A to make a complete prefix.
//                      The suffix can either be after a stop codon, or * can
//                      be replaced with [GG] to make an in-frame GS amino acid
//                      pair with GGA-TCC for protein fusions with a C terminal
//                      tag (Ct).
// T |  ATCC-CGCT   | Terminator part type.
// PR|  TACA-AATG   | Promoter+RBS part type. Commonly used for eukaryotic
//                      promoters, as kozak signals aren't commonly used in
//                      engineering toolkits.
// Cn|  A[ATG]-TTCT | N terminal tag part type. For tagging the front of
//                      proteins. The [TCT] should be in frame with the target
//                      protein - so tags should optimally end with [GG] in
//                      order to make an in-frame GS amino acid pair with
//                      GGT-TCT for protein fusions with Ct parts.
// Ct|  TTCT-*ATCC  | CDS part type for N terminal tagging. Should be in-frame
//                      with the prefix's TCT. * can be replaced with a stop
//                      codon or [GG] to make an in-frame GS amino acid pair.
// Cc|  ATCC-ATGT   | C tag. Used with C/Ct and Tt parts to C tag proteins.
// Tt|  ATGT-CGCT   | Terminator for C tag. Used with a C tags.
// PCRT|TACA-CGCT   | Identity part type. This part type is special in that it
//                      when retrieved, it has the part overhangs of a standard
//                      MoClo complete assembly. Used in combination with
//                      linkers to create useful multi-gene constructs.
//
// === PartTypes outside of linkers ===
// D |  GTCT-AAGC   | Origin of replication for target organism or downstream
//                      homology.
// S |  AAGC-ATAG   | Selection marker for target organism. For example, LEU2
//                      could be used in S. cerevisiae.
// U |  ATTA-CGAG   | Upstream homology.
//
// === Cloning organism parts ===
// E1|  GTCT-CGAG   | Vector type 1. Used for basic MoClo assemblies
// E2|  ATAG-CGAG   | Vector type 2. Used for simple plasmids between organisms
// E3|  ATAG-ATTA   | Vector type 3. Used for integration vectors
//*****************************************************************************

type SynthesisRequest struct {
	Name     string `json:"name"`
	Sequence string `json:"sequence"`
	Vector   string `json:"vector"`
	PartType string `json:"part_type"`
}
