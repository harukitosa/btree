package main

import (
	"btree/btree"
	"fmt"
)

func main() {
	b := btree.BTreeCreate()
	btree.BTreeInsert(b, "a")
	btree.BTreeInsert(b, "b")
	btree.BTreeInsert(b, "c")
	btree.BTreeInsert(b, "d")
	x, y := btree.BTreeSearch(b.Root, "c")
	fmt.Println(x)
	fmt.Println(x.Key[y-1])
}
