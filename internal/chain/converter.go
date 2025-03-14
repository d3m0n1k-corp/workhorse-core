package chain

import (
	"fmt"
	"workhorse-core/internal/common/linked_list"
	"workhorse-core/internal/converters"
)

type ConverterList struct {
	linked_list.NonValidatedList[converters.BaseConverter]
}

func NewConverterList() *ConverterList {
	return &ConverterList{}
}

func (cl *ConverterList) Validate() error {

	curr := cl.Head()
	for curr != nil && curr.Next != nil {
		nxt := curr.Next
		if nxt.Value.InputType() != curr.Value.OutputType() {
			return fmt.Errorf("Input type of converter %s does not match output type of converter %s", nxt.Value, curr.Value)
		}
	}

	return nil
}
