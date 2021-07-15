package main

import (
	"btree/btree"
	"fmt"
)

func main() {
	b := btree.BTreeCreate()
	btree.BTreeInsert(b, "a")
	btree.PrintBTree(b.Root)
	fmt.Println("-------------")
	btree.BTreeInsert(b, "b")
	btree.PrintBTree(b.Root)
	fmt.Println("-------------")

	btree.BTreeInsert(b, "c")
	btree.PrintBTree(b.Root)
	fmt.Println("-------------")

	btree.BTreeInsert(b, "d")
	btree.PrintBTree(b.Root)
	fmt.Println("-------------")

	btree.BTreeInsert(b, "e")
	btree.PrintBTree(b.Root)
	fmt.Println("-------------")

	btree.BTreeInsert(b, "f")
	btree.PrintBTree(b.Root)
	fmt.Println("-------------")

	x, y := btree.BTreeSearch(b.Root, "a")
	fmt.Println(x.Key[y-1])
}
