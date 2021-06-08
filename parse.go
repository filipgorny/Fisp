package main

type Ast struct {
	children []Ast
	elements []Element
}

type Tree interface {
	Pop() Ast
	Push(ast Ast)
}

func (ast *Ast) Pop() Ast {
	lastElement := ast.children[len(ast.children)-1]

	ast.children = ast.children[:len(ast.children)-1]

	return lastElement
}

func (ast *Ast) Push(newAst Ast) {
	ast.children = append(ast.children, newAst)
}

func parse(tokens []LToken) Ast {
	var ast Ast
	ast.children = append(ast.children, Ast{})

	for _, token := range tokens {
		if token.kind == LT_OPEN_LIST {
			var newAst Ast
			ast.Push(newAst)
		} else if token.kind == LT_CLOSE_LIST {
			oldAst := ast.Pop()

			ast.children[len(ast.children)-1].Push(oldAst)
		} else {
			currentExpression := ast.Pop()

			if token.kind == LT_NUMBER {
				var number Number
				number.numberValue = token.numberValue
				currentExpression.elements = append(currentExpression.elements, number)
			} else if token.kind == LT_SYMBOL {
				var symbol Symbol
				symbol.stringValue = token.stringValue
				currentExpression.elements = append(currentExpression.elements, symbol)
			}

			ast.Push(currentExpression)
		}
	}

	return ast
}
