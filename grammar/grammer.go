package grammar

import (
	"time"
)

type SlowGrammarChecker struct{}

func (*SlowGrammarChecker) GrammarCheck(text string) bool {
	// check line by line
	time.Sleep(time.Second * 6)
	return true
}

type GrammarChecker interface{
	GrammarCheck(text string) bool
}

type CloudGrammerChecker struct{}

func (CloudGrammerChecker) GrammarCheck(text string) bool {
	time.Sleep(time.Second * 1)
	return true
}

func NewCloudGrammerChecker() GrammarChecker {
	return CloudGrammerChecker{}
}