package utils

import (
	"crypto/sha1"
	"math"
	"sort"
	"strconv"
)

const (
	DefaultVirtualSpots = 40
)

type node struct {
	nodeKey   string
	spotValue uint32
}

type nodesArray []node

func (p nodesArray) Len() int           { return len(p) }
func (p nodesArray) Less(i, j int) bool { return p[i].spotValue < p[j].spotValue }
func (p nodesArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p nodesArray) Sort()              { sort.Sort(p) }

type Hashring struct {
	virtualSpots int //每个节点对应的虚拟节点数
	nodes        nodesArray
	weights      map[string]int
}

func NewHashRing(spots int) *Hashring {
	if spots == 0 {
		spots = DefaultVirtualSpots
	}
	h := &Hashring{
		virtualSpots: spots,
		weights:      map[string]int{},
	}
	return h
}

func genValue(bs []byte) uint32 {
	if len(bs) < 4 {
		return 0
	}
	v := (uint32(bs[3] << 24)) | (uint32(bs[2] << 16)) | (uint32(bs[1] << 8)) | (uint32(bs[0]))
}

func (h *Hashring) gen() {
	var totalW int
	for _, w := range h.weights {
		totalW += w
	}
	totalVirtualSpots := h.virtualSpots * len(h.weights)
	h.nodes = nodesArray{}
	for nodeKey, w := range h.weights {
		spots := int(math.Floor(float64(w) / float64(totalW) * float64(totalVirtualSpots)))
		for i := 1; i <= spots; i++ {
			hash := sha1.New()
			hash.Write([]byte(nodeKey + ":" + strconv.Itoa(i)))
			hashBytes := hash.Sum(nil)
			n := node{
				nodeKey:   nodeKey,
				spotValue: genValue(hashBytes[6:10]),
			}
			h.nodes = append(h.nodes, n)
			hash.Reset()
		}
	}
	h.nodes.Sort()
}

func (h *Hashring) AddNodes(nodeWeight map[string]int) {
	for nodeKey, w := range nodeWeight {
		h.weights[nodeKey] = w
	}
	h.gen()
}

func (h *Hashring) AddNode(nodekey string, weight int) {
	h.weights[nodekey] = weight
	h.gen()
}

func (h *Hashring) RemoveNode(nodeKey string) {
	delete(h.weights, nodeKey)
	h.gen()
}

func (h *Hashring) UpdateNode(nodeKey string, weight int) {
	h.weights[nodeKey] = weight
	h.gen()
}

func (h *Hashring) GetNode(s string) string {
	if len(h.nodes) == 0 {
		return ""
	}

	hash := sha1.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)
	v := genValue(hashBytes[6:10])
	i := sort.Search(len(h.nodes), func(i int) bool { return h.nodes[i].spotValue >= v })
	for i == len(h.nodes) {
		i = 0
	}
	return h.nodes[i].nodeKey
}
