package ds

type DirectedGraph struct {
	V int
	adj []*LinkedList
	connected []int
}

func NewDirectedGraph(v int) *DirectedGraph {
	return &DirectedGraph{V: v, adj: make([]*LinkedList,v),connected:make([]int,v)}
}

func (dg DirectedGraph) AddEdge(v int,w int) {
	dg.adj[v].PushBack(w)
}

func (dg DirectedGraph) DFS(root int) {
	visited := make([]bool,dg.V)
	dg.dfsUtil(root,visited)
}

func (dg DirectedGraph) dfsUtil(root int,visited []bool){
	visited[root] = true
	cur := dg.adj[root].head
	for cur!=nil{
		if !visited[cur.data] {
			dg.dfsUtil(cur.data,visited)
		}
		cur = cur.next
	}
}
