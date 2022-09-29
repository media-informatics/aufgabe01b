package main

import (
	"flag"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	var fname string
	var out string
	flag.StringVar(&fname, "file", "bsptree.txt", "Pfad zu Textdatei")
	flag.StringVar(&out, "out", "unique.txt", "Ausgabedatei")
	flag.Parse()

	ws, err := regexp.Compile("([[:digit:]]|[[:space:]]|[[:punct:]])+")
	if err != nil {
		log.Fatal(err)
	}
	text, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
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
	os.WriteFile(out, []byte(strings.Join(list, "\n")), 0664)
}
