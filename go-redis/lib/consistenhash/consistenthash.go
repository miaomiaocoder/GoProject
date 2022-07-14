package consistenhash

import (
	"hash/crc32"
	"sort"
)

type HashFunc func(data []byte) uint32

type NodeMap struct {
	hashFunc    HashFunc
	nodeHashs   []int // 12343, 23123, 59898
	nodehashMap map[int]string
}

func NewNodeMap(fn HashFunc) *NodeMap {
	m := &NodeMap{
		hashFunc:    fn,
		nodehashMap: make(map[int]string),
	}
	if m.hashFunc == nil {
		m.hashFunc = crc32.ChecksumIEEE
	}
	return m
}

func (m *NodeMap) IsEmpty() bool {
	return len(m.nodeHashs) == 0
}

func (m *NodeMap) AddNode(keys ...string) {
	for _, key := range keys {
		if key == "" {
			continue
		}
		hash := int(m.hashFunc([]byte(key)))

		// 将虚拟节点添加到环上
		m.nodeHashs = append(m.nodeHashs, hash)

		// 注册虚拟节点到物理节点的映射
		m.nodehashMap[hash] = key
	}
	sort.Ints(m.nodeHashs)
}

func (m *NodeMap) PickNode(key string) string {
	if m.IsEmpty() {
		return ""
	}
	hash := int(m.hashFunc([]byte(key)))

	// sort.Search 会使用二分查找法搜索 keys 中满足 m.nodeHashs[i] >= hash 的最小 i 值
	idx := sort.Search(len(m.nodeHashs), func(i int) bool {
		return m.nodeHashs[i] >= hash
	})

	// 若 key 的 hash 值大于最后一个虚拟节点的 hash 值，则 sort.Search 找不到目标
	// 这种情况下选择第一个虚拟节点
	if idx == len(m.nodeHashs) {
		idx = 0
	}

	// 将虚拟节点映射为实际地址
	return m.nodehashMap[m.nodeHashs[idx]]
}
