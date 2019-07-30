package editor

import (
	"IoCProject/grammar"
	"IoCProject/spelling"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTextEditor_Check(t *testing.T) {
	editor := &TextEditor{}
	res := editor.Check("some valid text")
	if res != true {
		t.Errorf("expect true but got %t", res)
	} //So slow... we need a mock
	//What if it required networking - DB/rabbit/S3?
}




func TestIoCTextEditor_Check(t *testing.T) {
	passingGrammarChecker := &MockGrammarChecker{}
	passingSpellChecker := &MockSpellChecker{}

	passingGrammarChecker.On("GrammarCheck", mock.Anything).Return(true)
	passingSpellChecker.On("SpellCheck", mock.Anything).Return(true)

	editor := &IoCTextEditor{
		GrammarChecker: passingGrammarChecker,
		SpellChecker:   passingSpellChecker,
	}

	res := editor.Check("some valid text")
	if res != true {
		t.Errorf("expect true but got %t", res)
	} // Much better - fast & a real unit test!
}

type MockGrammarChecker struct{
	grammar.GrammarChecker
	mock.Mock
}

func (mg *MockGrammarChecker) GrammarCheck(text string) bool {
	args := mg.Called(text)
	return args.Bool(0)
}

type MockSpellChecker struct{
	spelling.SpellChecker
	mock.Mock
}

func (mg *MockSpellChecker) SpellCheck(text string) bool {
	args := mg.Called(text)
	return args.Bool(0)
}
