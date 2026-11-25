package anago

import "fmt"
import "os"
import "slices"

func Anagrammi(s string) (risultato []string) {
	if len(s) == 1 {
		risultato = []string{s}
		return
	}
	for i := 0; i < len(s); i++ {
		letteraIniziale := string(s[i])
		resto := s[:i] + s[i+1:]
		anag := Anagrammi(resto)
		for _, anagramma := range anag {
			nuovoAnagramma := letteraIniziale + anagramma
			if !slices.Contains(risultato, nuovoAnagramma) {
				risultato = append(risultato, nuovoAnagramma)
			}
		}
	}
	return risultato
}

func finta(s string) string {
	return "*" + s + "*"
}

func main() {
	s := os.Args[1]
	x := Anagrammi(s)
	for _, anagramma := range x {
		fmt.Println(anagramma)
	}
}
