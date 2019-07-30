package main

import (
	"IoCProject/editor"
	"IoCProject/grammar"
	"IoCProject/spelling"
	"fmt"
	"os"
)

func main() {
	//Non IoC Editor
	t := &editor.TextEditor{}
	fmt.Println("text received: ")
	fmt.Println(os.Args[1])
	fmt.Println(t.Check(os.Args[1]))

	// Dependency Inversion Principle (DIP)
	t2 := &editor.IoCTextEditor{
		GrammarChecker: grammar.CloudGrammerChecker{},
		SpellChecker: spelling.CloudSpellChecker{},
	}
	fmt.Println(os.Args[1])
	fmt.Println(t2.Check(os.Args[1]))

	// Automatic dependency injection
	t3 := InitializeTextEditor()

	fmt.Println(os.Args[1])
	fmt.Println(t3.Check(os.Args[1]))
}
