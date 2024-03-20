package main

type cellList []*cell

func (c cellList) Len() int {
	return len(c)
}

func (c cellList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c cellList) Less(i, j int) bool {
	if c[i].f < c[j].f && c[i].index != c[j].index {
		return c[i].index < c[j].index
	}
	return c[i].f < c[j].f
}
