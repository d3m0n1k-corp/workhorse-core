package linked_list

import "workhorse-core/internal/converters/basec"

type ConverterList struct {
	baseList[basec.BaseConverter]
}

func NewConverterList() *ConverterList {
	return &ConverterList{}
}

func (cl *ConverterList) Validate() error {

	curr := cl.Head()
	for curr != nil && curr.Next != nil {
		nxt := curr.Next
		if nxt.Value.Inputtype() == curr.Value.Outputtype() {

		}

	}

	return nil
}
