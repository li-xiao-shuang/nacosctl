package printer

import "github.com/gosuri/uitable"

func NewTable(MaxColWidth uint) *uitable.Table {
	table := uitable.New()
	table.MaxColWidth = MaxColWidth
	return table
}

func NewTableWrap(MaxColWidth uint, wrap bool) *uitable.Table {
	table := uitable.New()
	table.MaxColWidth = MaxColWidth
	table.Wrap = wrap
	return table
}
