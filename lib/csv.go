package lib

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// ReadCSV reads in a .csv file and outputs a DataFrame
func ReadCSV(path string) (*DataFrame, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(fd)
	header, err := r.Read()
	if err != nil {
		return nil, err
	}

	columnsLen := len(header)
	if columnsLen == 0 {
		return nil, fmt.Errorf("Needs to contain at least one column")
	}

	columns := make([]Column, columnsLen)
	result := &DataFrame{
		Columns: columns,
		Name:    path,
	}

	for i, name := range header {
		columns[i] = Column{
			Name: name,
		}
	}

	contents, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	rowsLen := len(contents)
	if rowsLen == 0 {
		return nil, fmt.Errorf("Needs to contain at least one row")
	}

	for i := range result.Columns {
		result.Columns[i].Content = make([]float64, rowsLen)
	}

	for i, row := range contents {
		for j, column := range row {
			columns[j].Content[i], err = strconv.ParseFloat(column, 64)
			if err != nil {
				return nil, fmt.Errorf(
					"Error in %d:%d: %v", i, j, err)
			}
		}
	}
	return result, nil
}
