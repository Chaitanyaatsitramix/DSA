package unionfind

// ArrayUnionFind represents a Union-Find data structure using arrays
type ArrayUnionFind struct {
	parent []int // parent[i] represents the parent of element i
	rank   []int // rank[i] represents the height of the tree rooted at i
	size   []int // size[i] represents the size of the set containing i
}

// NewArrayUnionFind creates a new Union-Find data structure with n elements
func NewArrayUnionFind(n int) *ArrayUnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	size := make([]int, n)

	// Initialize each element as its own set
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
		size[i] = 1
	}

	return &ArrayUnionFind{
		parent: parent,
		rank:   rank,
		size:   size,
	}
}

// Find returns the representative (root) of the set containing x
// Uses path compression for efficiency
func (uf *ArrayUnionFind) Find(x int) int {
	if uf.parent[x] != x {
		// Path compression: make all nodes on path point directly to root
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union merges the sets containing x and y
// Uses union by rank for efficiency
func (uf *ArrayUnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return // Already in same set
	}

	// Union by rank: attach smaller rank tree under root of higher rank tree
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
		uf.size[rootX] += uf.size[rootY]
	}
}

// Connected checks if x and y are in the same set
func (uf *ArrayUnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// GetSize returns the size of the set containing x
func (uf *ArrayUnionFind) GetSize(x int) int {
	return uf.size[uf.Find(x)]
}

// GetSets returns all disjoint sets as separate slices
func (uf *ArrayUnionFind) GetSets() [][]int {
	sets := make(map[int][]int)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		sets[root] = append(sets[root], i)
	}

	result := make([][]int, 0, len(sets))
	for _, set := range sets {
		result = append(result, set)
	}
	return result
}
