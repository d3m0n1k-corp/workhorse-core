package chain_execute_test

import (
	"encoding/json"
	"os"
	"testing"
	"workhorse-core/app"
	"workhorse-core/internal/chain"

	"github.com/stretchr/testify/suite"
)

type ChainTest struct {
	Chain             []chain.ConverterChainLink
	ExpectedResponses []*app.ChainResponse
	input             string
}

type ChainExecuteSuite struct {
	suite.Suite
	ChainTestCases []ChainTest
}

func (s *ChainExecuteSuite) SetupTest() {
	// Read chains.json file and unmarshal it into s.ChainTestCases
	file, err := os.Open("chains.json")
	if err != nil {
		s.Require().NoError(err, "Failed to open chains.json file")
	}
	defer func() {
		err := file.Close()
		s.Require().NoError(err, "Failed to close chains.json file")
	}()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.ChainTestCases); err != nil {
		s.Require().NoError(err, "Failed to decode chains.json file")
	}
}

func (s *ChainExecuteSuite) TestChainExecute() {
	for _, testCase := range s.ChainTestCases {
		// Create a new chain with the provided links
		chain, err := app.ExecuteChain(testCase.Chain, testCase.input)
		s.Require().NoError(err, "Failed to execute chain")
		s.Require().Equal(len(testCase.ExpectedResponses), len(chain), "Chain response length mismatch")
		for i, response := range chain {
			s.Require().Equal(testCase.ExpectedResponses[i].Output, response.Output, "Chain response output mismatch")
			s.Require().Equal(testCase.ExpectedResponses[i].Error, response.Error, "Chain response error mismatch")
		}
	}
}

func TestRun(t *testing.T) {
	suite.Run(t, new(ChainExecuteSuite))
}
