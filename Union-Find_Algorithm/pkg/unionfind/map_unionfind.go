package unionfind

// MapUnionFind represents a Union-Find data structure using hash maps
type MapUnionFind struct {
	parent map[string]string // parent[x] represents the parent of element x
	rank   map[string]int    // rank[x] represents the height of the tree rooted at x
	size   map[string]int    // size[x] represents the size of the set containing x
}

// NewMapUnionFind creates a new Union-Find data structure
func NewMapUnionFind() *MapUnionFind {
	return &MapUnionFind{
		parent: make(map[string]string),
		rank:   make(map[string]int),
		size:   make(map[string]int),
	}
}

// MakeSet creates a new set with element x if it doesn't exist
func (uf *MapUnionFind) MakeSet(x string) {
	if _, exists := uf.parent[x]; !exists {
		uf.parent[x] = x
		uf.rank[x] = 0
		uf.size[x] = 1
	}
}

// Find returns the representative (root) of the set containing x
// Uses path compression for efficiency
func (uf *MapUnionFind) Find(x string) string {
	if _, exists := uf.parent[x]; !exists {
		uf.MakeSet(x)
		return x
	}

	if uf.parent[x] != x {
		// Path compression: make all nodes on path point directly to root
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union merges the sets containing x and y
// Uses union by rank for efficiency
func (uf *MapUnionFind) Union(x, y string) {
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
func (uf *MapUnionFind) Connected(x, y string) bool {
	return uf.Find(x) == uf.Find(y)
}

// GetSize returns the size of the set containing x
func (uf *MapUnionFind) GetSize(x string) int {
	return uf.size[uf.Find(x)]
}

// GetSets returns all disjoint sets as separate maps
func (uf *MapUnionFind) GetSets() map[string][]string {
	sets := make(map[string][]string)
	for element := range uf.parent {
		root := uf.Find(element)
		sets[root] = append(sets[root], element)
	}
	return sets
}
