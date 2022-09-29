package main

import (
	"flag"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

const newline = "\n"

func main() {
	var fname string
	var out string
	flag.StringVar(&fname, "file", "bsptree.txt", "Pfad zu Textdatei")
	flag.StringVar(&out, "out", "", "Ausgabedatei")
	flag.Parse()

	ws, err := regexp.Compile("([[:digit:]]|[[:space:]]|[[:punct:]])+")
	if err != nil {
		log.Fatalf("regular expression not compiled %w", err)
	}
	text, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalf("could not read file %w", err)
	}
	tokens := ws.Split(string(text), -1)
	unique := make(map[string]struct{}) // go idiom, like Java HashSet
	empty := struct{}{}
	for _, word := range tokens {
		lower := strings.ToLower(word)
		unique[lower] = empty
	}
	list := make([]string, len(unique))
	i := 0
	for word := range unique {
		list[i] = word
		i++
	}
	sort.Strings(list)

	fout := os.Stdout
	if len(out) > 0 {
		fout, err = os.OpenFile(out, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("error opening file %w", err)
		}
		defer fout.Close()
	}
	_, err = fout.WriteString(strings.Join(list, newline) + newline)
	if err != nil {
		log.Fatalf("error writing to buffer %w", err)
	}
}
