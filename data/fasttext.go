package data

import (
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"strings"
)

type AddressSet struct {
	set map[string]bool
}

func (s *AddressSet) Add(str string) bool {
	_, found := s.set[str]
	if found {
		return false
	}
	s.set[str] = true
	return true
}

type FastText struct {
}

func (f *FastText) GenClassifyData() {
	pcdt, err := os.Open("./data/pcdt.csv")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer pcdt.Close()
	reader := csv.NewReader(bufio.NewReader(pcdt))

	classify, err := os.Create("./data/classify.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer classify.Close()
	writer := bufio.NewWriter(classify)

	provinces := AddressSet{make(map[string]bool)}
	cities := AddressSet{make(map[string]bool)}
	districts := AddressSet{make(map[string]bool)}
	towns := AddressSet{make(map[string]bool)}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		province := strings.Replace(record[0], "\"", "", -1)
		city := strings.Replace(record[1], "\"", "", -1)
		district := strings.Replace(record[2], "\"", "", -1)
		town := strings.Replace(record[3], "\"", "", -1)
		provinces.Add(province)
		cities.Add(city)
		districts.Add(district)
		towns.Add(town)

	}
	for p := range provinces.set {
		if "" == p {
			continue
		}
		fmt.Fprintf(writer, "__label__province %s\n", p)
	}
	for c := range cities.set {
		fmt.Fprintf(writer, "__label__city %s\n", c)
	}
	for d := range districts.set {
		fmt.Fprintf(writer, "__label__district %s\n", d)
	}
	for t := range towns.set {
		fmt.Fprintf(writer, "__label__town %s\n", t)
	}
	writer.Flush()
}
