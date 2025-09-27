package chain

import (
	"fmt"
	"reflect"
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
		return fmt.Errorf("converter chain is empty: cannot validate empty chain")
	}
	curr := cl.Head()
	for curr != nil && curr.Next != nil {
		nxt := curr.Next
		if nxt.Value.InputType() != curr.Value.OutputType() {
			return fmt.Errorf("type mismatch in chain: converter '%s' outputs '%s' but next converter '%s' expects '%s'",
				reflect.TypeOf(curr.Value).Elem().Name(), curr.Value.OutputType(),
				reflect.TypeOf(nxt.Value).Elem().Name(), nxt.Value.InputType())
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
