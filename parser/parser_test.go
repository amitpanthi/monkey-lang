package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("Parser was unable to parse any tokens")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got %d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, test := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, test.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("parser was unable to parse any tokens")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements has length %d, expected %d",
			len(program.Statements), 3)
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Encountered statement was not *ast.ReturnStatement, got %T",
				stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("statement's token does not match, expected %s got %s",
				token.RETURN, returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("len(program.Statements) is not 1, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("given input is not of type *ast.ExpressionStatement, got %T",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("given input is not of type *ast.Identifier, got %T",
			ident)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral() is not 'foobar', got %s instead",
			ident.TokenLiteral())
	}

	if ident.Value != "foobar" {
		t.Fatalf("ident.Value is not 'foobar', got %s instead", ident.Value)
	}
}

func TestIntegerLiteralExpressions(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("len(program.Statements) is not 1, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("given input is not of type *ast.ExpressionStatement, got %T",
			program.Statements[0])
	}

	intStmt, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("given input is not of type *ast.IntegerLiteral, got %T",
			stmt.Expression)
	}

	if intStmt.TokenLiteral() != "5" {
		t.Fatalf("ident.TokenLiteral() is not '5', got %s instead",
			intStmt.TokenLiteral())
	}

	if intStmt.Value != 5 {
		t.Fatalf("ident.Value is not '5', got %d instead", intStmt.Value)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Fatalf("s.TokenLiteral() not let, got %s", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Fatalf("s is not an ast.LetStatement, got %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Fatalf("letStmt.Name.Value not %s, got %s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Fatalf("letStmt.Name.TokenLiteral() not %s, got %s", name,
			letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, err := range errors {
		t.Errorf("parser error: %q", err)
	}
	t.FailNow()
}
