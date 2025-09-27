package chain_execute_test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

type ChainTestCase struct {
	Name                    string                     `json:"name"`
	Description             string                     `json:"description"`
	Input                   string                     `json:"input"`
	Chain                   []chain.ConverterChainLink `json:"chain"`
	ExpectedResponses       []*app.ChainResponse       `json:"expected_responses"`
	ExpectedValidationError *string                    `json:"expected_validation_error,omitempty"`
}

type ChainExecuteSuite struct {
	suite.Suite
	ChainTestCases []ChainTest
	TestCaseFiles  []ChainTestCase
}

func (s *ChainExecuteSuite) SetupTest() {
	// Read test case files from test_cases subdirectories
	s.TestCaseFiles = []ChainTestCase{}

	testCasesDir := "test_cases"
	err := filepath.Walk(testCasesDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".json") {
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("failed to open test case file %s: %w", path, err)
			}
			defer file.Close()

			var testCase ChainTestCase
			decoder := json.NewDecoder(file)
			if err := decoder.Decode(&testCase); err != nil {
				return fmt.Errorf("failed to decode test case file %s: %w", path, err)
			}

			s.TestCaseFiles = append(s.TestCaseFiles, testCase)
		}

		return nil
	})

	s.Require().NoError(err, "Failed to load test case files")
	s.Require().Greater(len(s.TestCaseFiles), 0, "No test case files found")

	// Read legacy chains.json file for backward compatibility
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

func (s *ChainExecuteSuite) TestChainExecuteFromFiles() {
	for _, testCase := range s.TestCaseFiles {
		s.Run(testCase.Name, func() {
			// If this test case expects a validation error, test that separately
			if testCase.ExpectedValidationError != nil {
				_, err := app.ExecuteChain(testCase.Chain, testCase.Input)
				s.Require().Error(err, "Expected validation error for test case: %s", testCase.Name)
				s.Require().Contains(err.Error(), *testCase.ExpectedValidationError,
					"Expected validation error message mismatch for test case: %s", testCase.Name)
				return
			}

			// Execute the chain normally
			chain, err := app.ExecuteChain(testCase.Chain, testCase.Input)
			s.Require().NoError(err, "Failed to execute chain for test case: %s", testCase.Name)
			s.Require().Equal(len(testCase.ExpectedResponses), len(chain),
				"Chain response length mismatch for test case: %s", testCase.Name)

			for i, response := range chain {
				expectedOutput := testCase.ExpectedResponses[i].Output
				expectedError := testCase.ExpectedResponses[i].Error

				s.Require().Equal(expectedOutput, response.Output,
					"Chain response output mismatch at step %d for test case: %s", i, testCase.Name)
				s.Require().Equal(expectedError, response.Error,
					"Chain response error mismatch at step %d for test case: %s", i, testCase.Name)
			}
		})
	}
}

func TestRun(t *testing.T) {
	suite.Run(t, new(ChainExecuteSuite))
}
