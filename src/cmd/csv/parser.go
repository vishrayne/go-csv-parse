package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// documentation for csv is at http://golang.org/pkg/encoding/csv/
// TODO: could not find
func main() {
	file, err := os.Open("data/sample.csv")
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		fmt.Println("Error:", err)
		return
	}
	// automatically call Close() at the end of current method
	defer file.Close()
	//
	reader := csv.NewReader(file)
	// options are available at:
	// http://golang.org/src/pkg/encoding/csv/reader.go?s=3213:3671#L94
	reader.Comma = ','
	reader.LazyQuotes = true

	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(record)
		fmt.Println(len(record))
		// record is an array of string so is directly printable
		// // and we can iterate on top of that
		// for i := 0; i < len(record); i++ {
		// 	fmt.Println(" ", record[i])
		// }

		fmt.Println(record[4], ",", record[5], ",", record[8:13])
		fmt.Println("----------------------------")
	}
}
