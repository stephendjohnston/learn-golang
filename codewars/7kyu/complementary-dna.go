package main

import (
	"fmt"
)

func DNAStrand(dna string) (res string) {
	var complements = map[string]string{"A": "T", "T": "A", "G": "C", "C": "G"}
	for _, char := range dna {
		res += complements[string(char)]
	}
	return
}

func main() {
	fmt.Println(DNAStrand("AAAA") == "TTTT")
	fmt.Println(DNAStrand("ATTGC") == "TAACG")
	fmt.Println(DNAStrand("GTAT") == "CATA")
}
