package disjoinSet

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	equationSize := len(equations)
	uf := NewUnionFind(2 * equationSize)

  // step 1: preprocess, map string to id
	hashmap := make(map[string]int, 2*equationSize)
  id := 0
  for i, equation := range equations {
    var1 , var2 := equation[0], equation[1]

    if _,ok := hashmap[var1];!ok {
      hashmap[var1] = id
      id++
    }
    if _,ok := hashmap[var2];!ok {
      hashmap[var2] = id
      id++
    }
    uf.Union(hashmap[var1],hashmap[var2], values[i])
  }
  // step2: query the
  querySize := len(queries)
    res := make([]float64, querySize)
  for i, query := range queries {
    var1, var2 := query[0], query[1]
    id1, ok1 := hashmap[var1]
    id2, ok2 := hashmap[var2]
    // there is one doesnt exist
    if !(ok1&&ok2) {
      res[i] = -1.0
    } else {
      // weight rate of vars
      res[i] = uf.isConnected(id1,id2)
    }
  }
  return res



}

// UnionFind something

type UnionFind struct {
	parent []int
// weight to parent
weight []float64
}

//NewUnionFind something
func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{parent: make([]int, size), weight:make([]float64,size)}
for i:=0;i<size;i++ {
  uf.parent[i] = i
  uf.weight[i] = 1.0
}
	return uf
}

// Union something
func (uf *UnionFind) Union(x int, y int, value float64) {
rootX := uf.Find(x)
rootY := uf.Find(y)
if rootX == rootY {
  return
}
uf.parent[rootX] = uf.parent[rootY]
uf.weight[rootX] = uf.weight[y] * value / uf.weight[x]
}

// Find something
func (uf *UnionFind) Find(x int) int {
if x!=uf.parent[x] {
  origin := uf.parent[x]
  uf.parent[x] = uf.Find(uf.parent[x])
  uf.weight[x] *= uf.weight[origin]
}
return uf.parent[x]
}

//
func (uf *UnionFind) isConnected(x int, y int) float64 {
rootX, rootY := uf.Find(x), uf.Find(y)
if rootY == rootX {
  return uf.weight[x]/ uf.weight[y]
}
return -1.0
}
