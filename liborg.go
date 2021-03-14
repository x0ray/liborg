// Package liborg go library organization demo
package liborg

import (
	"log"
)

const LibVer = "0.0.2"

// Lib is a demo type
type Lib struct {
	name string
}

// New Loborg creates and initializes a new Lib struct
func NewLiborg(name string) *Lib {
	l := new(Lib)
	l.name = name
	return l
}


// SetName alters the name of the Lib struct
func (l *Lib) SetName(name string) {
	l.name = name 
}

// Dump displays the contents of the Lib struct on the log
func (l *Lib) Dump() {
	log.Printf("Package Lib %s Dump",LibVer)
	log.Printf("  Name: %s",l.name)
}
