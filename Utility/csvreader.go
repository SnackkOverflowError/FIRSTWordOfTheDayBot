package Utility

import (
	"encoding/csv"
	"io"
	"os"
)

func GetCsv() [][2]string {

	var csvData [][2]string

	csvFile, errRead := os.Open("data/words.csv")
	if errRead != nil {
		panic(errRead)
	}
	r := csv.NewReader(csvFile)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if len(record) != 2 {
			continue
		}
		if err != nil {
			panic(err)
		}


		pair := [2]string{record[0], record[1]}
		csvData = append(csvData,pair)
	}

	return csvData

}

