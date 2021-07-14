package btree

import "log"

const (
	t = 2
)

type Node struct {
	Key  [2*t - 1]string
	C    [2 * t]*Node
	Leaf bool
	N    int
}

func index(i int) int {
	return i - 1
}

type BTree struct {
	Root *Node
}

func allocateNode() *Node {
	return &Node{Leaf: false}
}

func DiskWrite(x *Node) {
}

func BTreeCreate() *BTree {
	T := BTree{}
	x := allocateNode()
	x.Leaf = true
	x.N = 0
	DiskWrite(x)
	T.Root = x
	return &T
}

func BTreeSplitChild(x *Node, i int) {
	z := allocateNode()
	y := x.C[index(i)]
	z.Leaf = y.Leaf
	z.N = t - 1
	for j := 1; j <= t-1; j++ {
		z.Key[index(j)] = y.Key[index(j+t)]
	}
	if !y.Leaf {
		for j := 1; j <= t; j++ {
			z.C[index(j)] = y.C[index(j+t)]
		}
	}
	y.N = t - 1
	for j := x.N + 1; j >= i+1; j-- {
		x.C[index(j+1)] = x.C[index(j)]
	}
	x.C[index(i+1)] = z
	for j := x.N; j >= i; j-- {
		x.Key[index(j+1)] = x.Key[index(j)]
	}
	x.Key[index(i)] = x.Key[index(t)]
	x.N = x.N + 1
	DiskWrite(x)
	DiskWrite(y)
	DiskWrite(z)
}

func BTreeInsert(T *BTree, k string) {
	r := T.Root
	if r.N == 2*t-1 {
		s := allocateNode()
		T.Root = s
		s.Leaf = false
		s.N = 0
		s.C[index(1)] = r
		BTreeSplitChild(s, 1)
		BTreeInsertNonFull(s, k)
	} else {
		BTreeInsertNonFull(r, k)
	}
}

func BTreeInsertNonFull(x *Node, k string) {
	i := x.N
	if x.Leaf {
		for i >= 1 && k < x.Key[index(i)] {
			x.Key[index(i+1)] = x.Key[index(i)]
			i--
		}
		x.Key[index(i+1)] = k
		x.N = x.N + 1
		DiskWrite(x)
	} else {
		for i >= 1 && k < x.Key[index(i)] {
			i--
		}
		i++
		log.Println(DiskRead(x.C[index(i)]))
		if x.C[index(i)].N == 2*t-1 {
			BTreeSplitChild(x, i)
			if k > x.Key[index(i)] {
				i++
			}
		}
		BTreeInsertNonFull(x.C[index(i)], k)
	}
}

func BTreeSearch(x *Node, k string) (*Node, int) {
	i := 1
	for i <= x.N && k > x.Key[index(i)] {
		i++
	}
	if i <= x.N && k == x.Key[index(i)] {
		return x, i
	} else if x.Leaf {
		return nil, -1
	} else {
		DiskRead(x.C[index(i)])
		return BTreeSearch(x.C[index(i)], k)
	}
}

func DiskRead(x *Node) string {
	return "ok"
}
