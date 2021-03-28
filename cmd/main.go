package main

import (
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/padraicbc/parser"
)

func main() {
	// fmt.Println(getScript("/home/padraic/test/layouts/components/home.svelte"))
	f, err := os.Open("/home/padraic/test/layouts/components/home.svelte")
	// f, err := os.Open("/home/padraic/test/layouts/global/head.svelte")
	if err != nil {
		log.Fatal(err)
	}
	_ = f

	log.SetFlags(log.Lshortfile)
	// ndes, err := GetNodes(f, true, "script")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // f.Seek(0, io.SeekStart)
	// // r, err := io.ReadAll(f)
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// for _, a := range ndes {
	// 	log.Println(a.Data.String(), a.Attrs)
	// 	// log.Println(string(r[a.startOffest:a.endOffset]))

	// }
	// log.Println(parser.JustOneContent(f, []byte("style")))
	jsparse()
}

func jsparse() {

	b, err := os.ReadFile("./examples/Promises.js")
	if err != nil {
		log.Fatal(err)
	}
	_ = b
	stream := antlr.NewInputStream(`
	import Home from "../components/test1.svelte";
	import {foo, bar} from "../components/text2.svelte";
	export let allContent;
  

	odds  = evens.map(v => v + 1);
	pairs = evens.map(v => ({ even: v, odd: v + 1 }));
	nums  = evens.map((v, i) => v + i)
	let top2 = allContent
	  .filter((content) => content.type == "article")
	  .sort((a, b) => a.date > b.date)
	  .slice(0, 2);
	  	  `)
	// Create the js Lexer
	lexer := parser.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJavaScriptParser(tokenStream)

	tree := p.Program()

	pr := antlr.NewParseTreeWalker()

	pr.Walk(&listener{}, tree)

}

func WriteTree(tree antlr.Tree) {
	log.Printf("%+v\n", tree.GetChildCount())
	for _, tr := range tree.GetChildren() {
		WriteTree(tr)
	}
}
