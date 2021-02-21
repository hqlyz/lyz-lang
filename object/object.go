package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
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
