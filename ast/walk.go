package ast

import (
	"fmt"
	"reflect"
)

// Visitor Visit method is invoked for each node encountered by Walk.
// If the result visitor w is not nil, Walk visits each of the children
// of node with the visitor w, followed by a call of w.Visit(nil).
type Visitor interface {
	Visit(n Node) (w Visitor)
}

// Walk traverses an AST in depth-first order: It starts by calling
// v.Visit(node); node must not be nil. If the visitor w returned by
// v.Visit(node) is not nil, Walk is invoked recursively with visitor
// w for each of the non-nil children of node, followed by a call
// of w.Visit(nil).
func Walk(v Visitor, n Node) {
	if n == nil || reflect.ValueOf(n).IsNil() {
		return
	}
	if v = v.Visit(n); v == nil {
		return
	}

	switch n := n.(type) {
	case *ArrayLiteral:
		for _, ex := range n.Value {
			Walk(v, ex)
		}
	case *AssignExpression:
		Walk(v, n.Left)
		Walk(v, n.Right)
	case *BadExpression:
	case *BinaryExpression:
		Walk(v, n.Left)
		Walk(v, n.Right)
	case *BlockStatement:
		for _, s := range n.List {
			Walk(v, s)
		}
	case *BooleanLiteral:
	case *BracketExpression:
		Walk(v, n.Left)
		Walk(v, n.Member)
	case *BranchStatement:
		Walk(v, n.Label)
	case *CallExpression:
		Walk(v, n.Callee)
		for _, a := range n.ArgumentList {
			Walk(v, a)
		}
	case *CaseStatement:
		Walk(v, n.Test)
		for _, c := range n.Consequent {
			Walk(v, c)
		}
	case *CatchStatement:
		Walk(v, n.Parameter)
		Walk(v, n.Body)
	case *ConditionalExpression:
		Walk(v, n.Test)
		Walk(v, n.Consequent)
		Walk(v, n.Alternate)
	case *DebuggerStatement:
	case *DoWhileStatement:
		Walk(v, n.Test)
		Walk(v, n.Body)
	case *DotExpression:
		Walk(v, n.Left)
	case *EmptyExpression:
	case *EmptyStatement:
	case *ExpressionStatement:
		Walk(v, n.Expression)
	case *ForInStatement:
		Walk(v, n.Into)
		Walk(v, n.Source)
		Walk(v, n.Body)
	case *ForStatement:
		Walk(v, n.Initializer)
		Walk(v, n.Update)
		Walk(v, n.Test)
		Walk(v, n.Body)
	case *FunctionLiteral:
		Walk(v, n.Name)
		for _, p := range n.ParameterList.List {
			Walk(v, p)
		}
		Walk(v, n.Body)
	case *FunctionStatement:
		Walk(v, n.Function)
	case *Identifier:
	case *IfStatement:
		Walk(v, n.Test)
		Walk(v, n.Consequent)
		Walk(v, n.Alternate)
	case *LabelledStatement:
		Walk(v, n.Statement)
	case *NewExpression:
		Walk(v, n.Callee)
		for _, a := range n.ArgumentList {
			Walk(v, a)
		}
	case *NullLiteral:
	case *NumberLiteral:
	case *ObjectLiteral:
		for _, p := range n.Value {
			Walk(v, p.Value)
		}
	case *Program:
		for _, b := range n.Body {
			Walk(v, b)
		}
	case *RegExpLiteral:
	case *ReturnStatement:
		Walk(v, n.Argument)
	case *SequenceExpression:
		for _, e := range n.Sequence {
			Walk(v, e)
		}
	case *StringLiteral:
	case *SwitchStatement:
		Walk(v, n.Discriminant)
		for _, c := range n.Body {
			Walk(v, c)
		}
	case *ThisExpression:
	case *ThrowStatement:
		Walk(v, n.Argument)
	case *TryStatement:
		Walk(v, n.Body)
		Walk(v, n.Catch)
		Walk(v, n.Finally)
	case *UnaryExpression:
		Walk(v, n.Operand)
	case *VariableExpression:
		Walk(v, n.Initializer)
	case *VariableStatement:
		for _, e := range n.List {
			Walk(v, e)
		}
	case *WhileStatement:
		Walk(v, n.Test)
		Walk(v, n.Body)
	case *WithStatement:
		Walk(v, n.Object)
		Walk(v, n.Body)
	default:
		panic(fmt.Sprintf("Walk: unexpected node type %T", n))
	}

	Walk(v, nil)
}
