package main

import (
	"bufio"
	"os"
	"domain-serch/thesaurus"
	"log"
	"fmt"
)

func main()  {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("%q synonym search failed: %v\n", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("did not have any synonyms \n")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}