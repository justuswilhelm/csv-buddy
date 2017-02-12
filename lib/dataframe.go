package lib

// DataFrame stores several columns
type DataFrame struct {
	Name    string
	Columns []Column
}

// ColumnNames returns all DataFrame column names as string array
func (df *DataFrame) ColumnNames() []string {
	result := make([]string, len(df.Columns))
	for i, column := range df.Columns {
		result[i] = column.Name
	}
	return result
}
