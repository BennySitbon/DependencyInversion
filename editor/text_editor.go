package editor

import (
	"IoCProject/grammar"
	"IoCProject/spelling"
)

type TextEditor struct {}

func (TextEditor) Check(text string) bool {
	spellChecker := &spelling.SlowSpellChecker{}
	grammarChecker := &grammar.SlowGrammarChecker{}

	sRes := spellChecker.SpellCheck(text)
	gRes := grammarChecker.GrammarCheck(text)
	return sRes && gRes
}




type IoCTextEditor struct{
	GrammarChecker grammar.GrammarChecker
	SpellChecker   spelling.SpellChecker
}

func (t *IoCTextEditor) Check(text string) bool {
	sRes := t.SpellChecker.SpellCheck(text)
	gRes := t.GrammarChecker.GrammarCheck(text)
	return sRes && gRes
}



func NewIocTextEditor(g grammar.GrammarChecker, s spelling.SpellChecker) IoCTextEditor {
	return IoCTextEditor{
		GrammarChecker: g,
		SpellChecker:   s,
	}
}