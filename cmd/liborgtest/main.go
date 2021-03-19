// Package main -  command liborgtest
package main

/* Expected Output:
david@rat:~/go/src/github.com/x0ray/liborg$ liborgtest
2021/03/19 15:47:40 Package Lib 0.0.9 Dump
2021/03/19 15:47:40   Name  : Dave  Tag for lib(0): lib.name
2021/03/19 15:47:40   Name A:   Tag for NestA(0): nesta.name
2021/03/19 15:47:40   Name B: Dave_C  Tag for NestB(0):
2021/03/19 15:47:40   Name C: Dave_C  Tag for NestC(0): nestc.name
2021/03/19 15:47:40 Package Lib 0.0.9 Dump
2021/03/19 15:47:40   Name  : Dave  Tag for lib(0): lib.name
2021/03/19 15:47:40   Name A: Fred_A  Tag for NestA(0): nesta.name
2021/03/19 15:47:40   Name B: Fred_B  Tag for NestB(0):
2021/03/19 15:47:40   Name C: Fred_B  Tag for NestC(0): nestc.name
*/

import (
	o "github.com/x0ray/liborg"
)

func main() {
	d := o.NewLiborg("Dave")
	d.Dump()
	d.SetName("Fred")
	d.Dump()
}
