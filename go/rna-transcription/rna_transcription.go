package strand

import "bytes"

var (
	guanine  = 'G'
	cytosine = 'C'
	thymine  = 'T'
	adenine  = 'A'
	uracil   = 'U'
)

var rnaNucleotidesByDNANucleotides = map[rune]rune{
	guanine:  cytosine,
	cytosine: guanine,
	thymine:  adenine,
	adenine:  uracil,
}

func ToRNA(dna string) string {
	var b bytes.Buffer

	for _, n := range dna {
		b.WriteRune(rnaNucleotidesByDNANucleotides[n])
	}

	return b.String()
}
