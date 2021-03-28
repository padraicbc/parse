package main

import (
	"log"

	"github.com/padraicbc/parser"
)

type listener struct {
	parser.BaseJavaScriptParserListener
}

func show(fname, text string) {

	log.Printf("%s %s\n", fname, text)
}
func (l *listener) EnterBlock(ctx *parser.BlockContext) {
	show("EnterBlock", ctx.GetText())
}

func (l *listener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	show("EnterFunctionDeclaration", ctx.FunctionBody().GetText())
}

// EnterImportStatement is called when production importStatement is entered.
func (s *listener) EnterImportStatement(ctx *parser.ImportStatementContext) {

	show("EnterImportStatement", ctx.ImportFromBlock().GetText())

}

// ExitImportStatement is called when production importStatement is exited.
func (s *listener) ExitImportStatement(ctx *parser.ImportStatementContext) {
	show("ExitImportStatement", ctx.GetText())
	// log.Println(ctx.ToStringTree(nil, nil))

}

// EnterImportFromBlock is called when production importFromBlock is entered.
func (s *listener) EnterImportFromBlock(ctx *parser.ImportFromBlockContext) {

	show("EnterImportFromBlock", ctx.GetText())
}

// ExitImportFromBlock is called when production importFromBlock is exited.
func (s *listener) ExitImportFromBlock(ctx *parser.ImportFromBlockContext) {
	show("ExitImportFromBlock", ctx.GetText())

}

// EnterModuleItems is called when production moduleItems is entered.
func (s *listener) EnterModuleItems(ctx *parser.ModuleItemsContext) {
	show("EnterModuleItems", ctx.GetText())

}

// ExitModuleItems is called when production moduleItems is exited.
func (s *listener) ExitModuleItems(ctx *parser.ModuleItemsContext) {
	show("ExitModuleItems", ctx.GetText())

}

// EnterImportDefault is called when production importDefault is entered.
func (s *listener) EnterImportDefault(ctx *parser.ImportDefaultContext) {
	show("EnterImportDefault", ctx.GetText())

}

// ExitImportDefault is called when production importDefault is exited.
func (s *listener) ExitImportDefault(ctx *parser.ImportDefaultContext) {
	show("ExitImportDefault", ctx.GetText())
}

// EnterImportNamespace is called when production importNamespace is entered.
func (s *listener) EnterImportNamespace(ctx *parser.ImportNamespaceContext) {
	show("EnterImportNamespace", ctx.GetText())
}

// ExitImportNamespace is called when production importNamespace is exited.
func (s *listener) ExitImportNamespace(ctx *parser.ImportNamespaceContext) {
	show("ExitImportNamespace", ctx.GetText())
}

// EnterImportFrom is called when production importFrom is entered.
func (s *listener) EnterImportFrom(ctx *parser.ImportFromContext) {
	show("EnterImportFrom", ctx.GetText())
}

// ExitImportFrom is called when production importFrom is exited.
func (s *listener) ExitImportFrom(ctx *parser.ImportFromContext) {
	show("ExitImportFrom", ctx.GetText())
}
