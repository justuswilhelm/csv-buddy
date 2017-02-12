package lib

import (
	"fmt"
	"strings"
)

// SummaryCommand stores summary flags
type SummaryCommand struct {
}

func (df *DataFrame) printColumns() {
	names := df.ColumnNames()
	fmt.Printf("---\nColumn names\n---\n%s\n", strings.Join(names, ", "))
	fmt.Printf("---\nColumn summary\n---\n")
	for _, col := range df.Columns {
		fmt.Printf("%s\n", col.Summary())
	}
}

// Execute ...
func (x *SummaryCommand) Execute(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Must specify csv file")
	}

	data, err := ReadCSV(args[0])
	if err != nil {
		return err
	}

	data.printColumns()

	return nil
}
