//+build wireinject

package main

import (
	"IoCProject/editor"
	"github.com/google/wire"
	"IoCProject/grammar"
	"IoCProject/spelling"
)

func InitializeTextEditor() editor.IoCTextEditor{
	wire.Build(editor.NewIocTextEditor, grammar.NewCloudGrammerChecker, spelling.NewCloudSpellChecker)
	return editor.IoCTextEditor{}
}