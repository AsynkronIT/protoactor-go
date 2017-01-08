package cluster

import (
	"hash/fnv"
	"log"
	"math"

	"github.com/AsynkronIT/protoactor-go/actor"
)

const (
	hashSize = uint32(math.MaxUint32)
)

func getNode(key, kind string) string {
	v := hash(key)
	members := getMembers(kind)
	if members == nil {
		log.Printf("[CLUSTER] Failed to getNode")
		return actor.ProcessRegistry.Host
	}

	bestV := hashSize
	bestI := 0

	//walk all members and find the node with the closest distance to the id hash
	for i, n := range members {
		h := hash(n)     //hash the node address
		b := delta(h, v) //calculate the delta between key and node address
		if b < bestV {
			bestV = b
			bestI = i
		}
	}

	member := members[bestI]
	return member
}

func delta(l uint32, r uint32) uint32 {
	if l > r {
		return l - r
	}
	return r - l
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
