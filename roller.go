// roller
package main

type RollSpec struct {
	Sides    int64
	BestOf   int64
	DieCount int64
	Modifier int64
	Times    int64
}

type SetResult struct {
	Total int
	Count int
	Dies  []int
}

type RollResults struct {
	Count int
	Rolls []SetResult
}
