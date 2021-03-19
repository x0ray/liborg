// Package liborg go library organization demo
// This package only serves as a demonstration of:
//   An object library
//    - Constant name LibVer can be seen in the importing package
//    - Types  can be seen in the importing package
//   Nested structure member visibility
//   Struct member tags
package liborg

import (
	"fmt"
	"log"
	"reflect"
)

const LibVer = "0.0.9"

// NestA is a nested demo type
type NestA struct {
	Name string `mytag:"nesta.name"` // dont use blanks around colon :
}

// NestB is a nested demo type
type NestB struct {
	NestC
}

// NestC is a nested demo type
type NestC struct {
	Name string `mytag:"nestc.name"`
}

// lib is a demo type
type lib struct {
	Name string `mytag:"lib.name"`
	NestA
	NestB
}

type Face interface {
	SetName(string)
}

// New Loborg creates and initializes a new Lib struct
func NewLiborg(name string) *lib {
	l := new(lib)
	l.Name = name
	l.NestB.NestC.Name = name + "_C"
	return l
}

// SetName alters the name of the lib struct
func (l *lib) SetName(name string) {
	l.NestA.Name = name + "_A"
	l.NestB.Name = name + "_B"
}

// Dump displays the contents of the Lib struct on the log
func (l *lib) Dump() {
	// get the tag info for the field of the struct
	t := func(ifce interface{}, f int) string {
		var mytag string
		st := reflect.TypeOf(ifce)
		n := st.Name()
		field := st.Field(f)
		mytag = field.Tag.Get("mytag")
		return fmt.Sprintf("%s(%d): %s", n, f, mytag)
	}
	// print out the lib details
	log.Printf("Package Lib %s Dump", LibVer)
	log.Printf("  Name  : %s  Tag for %s", l.Name, t(lib{}, 0))
	log.Printf("  Name A: %s  Tag for %s", l.NestA.Name, t(NestA{}, 0))
	log.Printf("  Name B: %s  Tag for %s", l.NestB.Name, t(NestB{}, 0))
	log.Printf("  Name C: %s  Tag for %s", l.NestB.NestC.Name, t(NestC{}, 0))
}
