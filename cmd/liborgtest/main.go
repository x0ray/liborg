// Package main -  command liborgtest
package main

import (
	o "github.com/x0ray/liborg"
)

func main() {
	d := o.NewLiborg("Dave")
	d.Dump()
	d.SetName("Fred")
	d.Dump() 
}
