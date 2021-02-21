package object

import (
	"bytes"
	"fmt"
	"lyz-lang/ast"
	"strings"
)

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
)

// Object interface
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer object
type Integer struct {
	Value int64
}

// Inspect function
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type function
func (i *Integer) Type() ObjectType { return ObjectType(INTEGER_OBJ) }

// Boolean object
type Boolean struct {
	Value bool
}

// Inspect function
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type function
func (b *Boolean) Type() ObjectType { return ObjectType(BOOLEAN_OBJ) }

// Null object represents absence of value
type Null struct{}

// Inspect function
func (n *Null) Inspect() string { return "null" }

// Type function
func (n *Null) Type() ObjectType { return ObjectType(NULL_OBJ) }

// ReturnValue object
type ReturnValue struct {
	Value Object
}

// Inspect function
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Type function
func (rv *ReturnValue) Type() ObjectType { return ObjectType(RETURN_VALUE_OBJ) }

// Error object
type Error struct {
	Message string
}

// Inspect function
func (e *Error) Inspect() string { return "Error: " + e.Message }

// Type function
func (e *Error) Type() ObjectType { return ObjectType(ERROR_OBJ) }

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	newEnv := NewEnvironment()
	newEnv.outer = outer
	return newEnv
}

func NewEnvironment() *Environment {
	return &Environment{store: map[string]Object{}}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, obj Object) Object {
	e.store[name] = obj
	return obj
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }
