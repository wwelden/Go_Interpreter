package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEN_OBJ  = "BOOLEAN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}
type Boolean struct {
	Value bool
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (b *Boolean) Type() ObjectType { return BOOLEN_OBJ }

func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }