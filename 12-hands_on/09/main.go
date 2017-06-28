package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Row struct {
	Date     time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int64
	AdjClose float64
}

type data []Row

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	// parse csv
	rows := prs("table.csv")

	// parse template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute template
	err = tpl.Execute(res, rows)
	if err != nil {
		log.Fatalln(err)
	}
}

func prs(filePath string) data {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)
	recs, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	rows := make(data, 0, len(recs))

	for i, row := range recs {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)
		high, _ := strconv.ParseFloat(row[2], 64)
		low, _ := strconv.ParseFloat(row[3], 64)
		close, _ := strconv.ParseFloat(row[4], 64)
		volume, _ := strconv.ParseInt(row[5], 10, 64)
		adjClose, _ := strconv.ParseFloat(row[6], 64)

		rows = append(rows, Row{
			Date:     date,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    close,
			Volume:   volume,
			AdjClose: adjClose,
		})
	}
	return rows
}
