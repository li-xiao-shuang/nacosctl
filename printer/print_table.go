package printer

import "github.com/gosuri/uitable"

func NewTable(MaxColWidth uint) *uitable.Table {
	table := uitable.New()
	table.MaxColWidth = MaxColWidth
	return table
}
