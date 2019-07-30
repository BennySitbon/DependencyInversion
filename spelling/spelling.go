package spelling

import (
	"time"
)

type SlowSpellChecker struct {}

func (SlowSpellChecker) SpellCheck(text string) bool {
	time.Sleep(time.Second * 6)
	return true
}

type SpellChecker interface {
	SpellCheck(text string) bool
}

type CloudSpellChecker struct{}

var _ SpellChecker = CloudSpellChecker{}

func (CloudSpellChecker) SpellCheck(text string) bool {
	time.Sleep(time.Second * 1) //assume it does some HTTP calls
	return true
}

func NewCloudSpellChecker() SpellChecker {
	return CloudSpellChecker{}
}