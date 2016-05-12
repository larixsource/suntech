package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fn := "../../st600/ascii_spec.txt"
	f, ferr := os.Open(fn)
	if ferr != nil {
		log.Panicf("error reading %s: %+v", fn, ferr)
	}
	defer f.Close()

	i := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sz := scanner.Text()
		if len(sz) == 0 || sz[0] == '#' {
			continue
		}
		sz = strings.TrimPrefix(sz, "[command] ")
		sz = strings.TrimPrefix(sz, "[response] ")
		sz = strings.TrimPrefix(sz, "[report] ")
		name := fmt.Sprintf("spec_%d", i)
		writeSample(name, sz+"\r")
		i++
	}
	if scanner.Err() != nil {
		log.Panicf("error scanning %s: %+v", fn, scanner.Err())
	}

	return
}

func writeSample(name string, sz string) {
	fn := fmt.Sprintf("corpus/%s", name)
	wErr := ioutil.WriteFile(fn, []byte(sz), 0644)
	if wErr != nil {
		panic(wErr)
	}
}
