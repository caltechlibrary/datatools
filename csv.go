// datatools.go is a package for working with various types of data (e.g. CSV, XLSX, JSON) in support
// of the utilities included in the datatools.go package.
//
// Copyright (c) 2021, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package datatools

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"

	// 3rd Party Libraries
	"gopkg.in/yaml.v3"
)

func selectedRow(rowNo int, record []string, rowNos []int) []string {
	if len(rowNos) == 0 {
		return record
	}
	for _, i := range rowNos {
		if i == rowNo {
			return record
		}
	}
	return nil
}

func shuffleRows(rows [][]string, src rand.Source) {
	// Create our random number source
	rn := rand.New(src)
	for a := len(rows) - 1; a > 0; a-- {
		// Pick a random element to swap with
		b := rn.Intn(a + 1)
		// Swap with a random element
		rows[a], rows[b] = rows[b], rows[a]
	}
}

// CSVRandomRows reads a in, creates a csv Reader and Writer and randomly selectes the rowCount
// number of rows to write out.  If showHeader is true it is excluded from the random row selection
// and will be written to out before the randomized rows.
// rowCount is the number of rows to return independent of the header row.
func CSVRandomRows(in io.Reader, out io.Writer, showHeader bool, rowCount int, delimiter string, lazyQuotes, trimLeadingSpace bool) error {
	var err error

	headerRow := []string{}
	rows := [][]string{}

	r := csv.NewReader(in)
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace

	w := csv.NewWriter(out)
	if delimiter != "" {
		r.Comma = NormalizeDelimiterRune(delimiter)
		w.Comma = NormalizeDelimiterRune(delimiter)
	}

	// read in our rows.
	for i := 0; err != io.EOF; i++ {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("%s (%T %+v)", err, rec, rec)
		}
		if i == 0 && showHeader {
			headerRow = rec
		} else {
			rows = append(rows, rec)
		}
	}
	if showHeader && len(headerRow) > 0 {
		if err := w.Write(headerRow); err != nil {
			return fmt.Errorf("Error writing record to csv: %s (Row %T %+v)", err, headerRow, headerRow)
		}
	}

	// Shuffle the rows, then write out the desired number of rows.
	rSrc := rand.NewSource(time.Now().UnixNano())
	shuffleRows(rows, rSrc)

	// Now render the rowCount of the suffled rows
	if rowCount > len(rows) {
		rowCount = len(rows)
	}
	for i := 0; i < rowCount; i++ {
		row := rows[i]
		if row != nil {
			if err := w.Write(row); err != nil {
				return fmt.Errorf("Error writing record to csv: %s (Row %T %+v)", err, row, row)
			}
		}
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		return fmt.Errorf("%s\n", err)
	}
	return nil
}

// CSVRows renders the rows numbers in rowNos using the delimiter to out
func CSVRows(in io.Reader, out io.Writer, showHeader bool, rowNos []int, delimiter string, lazyQuotes, trimLeadingSpace bool) error {
	var err error

	r := csv.NewReader(in)
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace

	w := csv.NewWriter(out)
	if delimiter != "" {
		r.Comma = NormalizeDelimiterRune(delimiter)
		w.Comma = NormalizeDelimiterRune(delimiter)
	}
	for i := 0; err != io.EOF; i++ {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("%s (%T %+v)", err, rec, rec)
		}
		if i == 0 && showHeader {
			if err = w.Write(rec); err != nil {
				return fmt.Errorf("Error writing record to csv: %s (Row %T %+v)", err, rec, rec)
			}
		} else {
			row := selectedRow(i, rec, rowNos)
			if row != nil {
				if err = w.Write(row); err != nil {
					return fmt.Errorf("Error writing record to csv: %s (Row %T %+v)", err, row, row)
				}
			}
		}
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		return fmt.Errorf("%s\n", err)
	}
	return nil
}

// CSVRowsAll renders the all rows in rowNos using the delimiter to out
func CSVRowsAll(in io.Reader, out io.Writer, showHeader bool, delimiter string, lazyQuotes bool, trimLeadingSpace bool) error {
	var err error

	r := csv.NewReader(in)
	r.LazyQuotes = lazyQuotes
	r.TrimLeadingSpace = trimLeadingSpace

	w := csv.NewWriter(out)
	if delimiter != "" {
		r.Comma = NormalizeDelimiterRune(delimiter)
		w.Comma = NormalizeDelimiterRune(delimiter)
	}
	for i := 0; err != io.EOF; i++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("%s (%T %+v)", err, row, row)
		}
		if i == 0 && showHeader {
			if err = w.Write(row); err != nil {
				return fmt.Errorf("Error writing record to csv: %s (Row %T %+v)", err, row, row)
			}
			continue
		} else if i > 0 {
			if err = w.Write(row); err != nil {
				return fmt.Errorf("Error writing record to csv: %s (Row %T %+v)", err, row, row)
			}
		}
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		return fmt.Errorf("%s\n", err)
	}
	return nil
}

// JSONObjectsToCSV takes an JSON array of objects mapping to CSV colum/rows. This works a little
// like Python csv.DictWriter. In Go a `map[string]interface{}{}` is used to represent the object.
// If the value is complex then it is rendered as YAML into the cell.
func JSONObjectsToCSV(in io.Reader, out io.Writer, eout io.Writer, quiet bool, showHeader bool, delimiter string) error {
	src, err := io.ReadAll(in)
	if err != nil {
		return err
	}
	objList := []map[string]interface{}{}
	decoder := json.NewDecoder(bytes.NewReader(src))
	decoder.UseNumber()
	if err := decoder.Decode(&objList); err != nil {
		return err
	}

	// Write out CSV
	w := csv.NewWriter(out)
	if delimiter != "" {
		w.Comma = NormalizeDelimiterRune(delimiter)
	}

	// Collection column headers from objects.
	header := []string{}
	if showHeader {
		columnNames := map[string]bool{}
		for _, obj := range objList {
			for k, _ := range obj {
				columnNames[k] = true
			}
		}
		for k, _ := range columnNames {
			header = append(header, k)
		}
		// Order the columns alphabetically
		sort.Strings(header)
		if err = w.Write(header); err != nil {
			if ! quiet {
				fmt.Fprintf(eout, "WARNING: failed to write header: %s (row %T %+v)\n", err, header, header)
			}
		}
	}
	row := []string{}
	fmt.Fprintf(eout, "DEBUG object list length %d\n", len(objList))
	for i, obj := range objList {
		// clear the row before proceeding.
		row = []string{}
		// assemble the row by appending "cells"
		for j, col_name := range header {
			if val, ok := obj[col_name]; ok {
				cell := ""
				switch val.(type) {
					case json.Number:
					  cell = fmt.Sprintf("%s", val)
					case string:
					  cell = fmt.Sprintf("%s", val)
					case bool:
					  cell = fmt.Sprintf("%t", val)
					case int:
					  cell = fmt.Sprintf("%d", val)
					case int64:
					  cell = fmt.Sprintf("%d", val)
					case float64:
					  cell = fmt.Sprintf("%f", val)
					default:
						// we have a complex object, render as YAML.
						src, err := yaml.Marshal(val)
						if err != nil {
							if ! quiet {
								fmt.Fprintf(eout, "failed to convert %+v in row %d, column %d\n", val, i+1, j+1)
							}
						}
						cell = fmt.Sprintf("%s", src)
				}
				row = append(row, cell)
			} else {
				row = append(row, "")
			}
		}
		if err = w.Write(row); err != nil {
			if ! quiet {
				fmt.Fprintf(eout, "failed to write row %d, %+v\n", len(objList)+1, row)
			}
		}
	}
	if err = w.Write(row); err != nil {
		if ! quiet {
			fmt.Fprintf(eout, "failed to write row %d, %+v\n", len(objList)+1, row)
		}
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		return fmt.Errorf("%s\n", err)
	}
	return nil
}
