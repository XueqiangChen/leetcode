package utils

type UnionFind struct {
	Count  int
	Parent []int
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{}
	u.Count = n
	u.Parent = make([]int, n)
	for i := 0; i < n; i++ {
		u.Parent[i] = i
	}
	return u
}

func (u *UnionFind) Find(p int) int {
	for p != u.Parent[p] {
		u.Parent[p] = u.Parent[u.Parent[p]]
		p = u.Parent[p]
	}
	return p
}

func (u *UnionFind) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
	}
	u.Parent[rootP] = rootQ
	u.Count--
}
