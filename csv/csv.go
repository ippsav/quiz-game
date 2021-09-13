package csv

import (
	"encoding/csv"
	"log"
	"os"
)

type CSVFile struct {
	f *os.File
	l *log.Logger
}

func NewCSVFile(f *os.File, l *log.Logger) *CSVFile {
	return &CSVFile{f, l}
}

func (cFile *CSVFile) ParseCSVFile() [][]string {
	csvReader := csv.NewReader(cFile.f)
	lines, err := csvReader.ReadAll()
	if err != nil {
		cFile.l.Fatal(err)
	}
	return lines
}
