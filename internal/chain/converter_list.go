package chain

import (
	"fmt"
	"workhorse-core/internal/common/linked_list"
	"workhorse-core/internal/converters"

	"github.com/sirupsen/logrus"
)

var mockableNewConverterFunc = converters.NewConverter

type ConverterList struct {
	linked_list.NonValidatedList[converters.BaseConverter]
}

type ConverterChainLink struct {
	Name       string `json:"name"`
	ConfigJSON string `json:"config_json"`
}

func (cl *ConverterList) Validate() error {
	if cl.Length() == 0 {
		return fmt.Errorf("Converter list is empty")
	}
	curr := cl.Head()
	for curr != nil && curr.Next != nil {
		nxt := curr.Next
		if nxt.Value.InputType() != curr.Value.OutputType() {
			return fmt.Errorf("Input type of converter %s does not match output type of converter %s", nxt.Value, curr.Value)
		}
		curr = nxt
	}

	return nil
}

func NewConverterListFromJSON(chainLinks []ConverterChainLink) (*ConverterList, error) {
	cl := &ConverterList{}
	for _, link := range chainLinks {
		logrus.Tracef("Creating converter %s with config %s", link.Name, link.ConfigJSON)
		c, err := mockableNewConverterFunc(link.Name, link.ConfigJSON)
		if err != nil {
			return nil, err
		}
		cl.Append(c)
	}
	err := cl.Validate()
	if err != nil {
		return nil, err
	}
	return cl, nil
}
