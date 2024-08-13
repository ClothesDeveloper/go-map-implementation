package main

type MyMap struct {
	buckets [][]node
}

type node struct {
	key   string
	value int
}

func NewMap(bucketsCount int) *MyMap {
	var buckets [][]node

	for i := 0; i < bucketsCount; i++ {
		bucket := []node{}
		buckets = append(buckets, bucket)
	}

	return &MyMap{
		buckets: buckets,
	}
}

func (m *MyMap) Set(key string, val int) {
	index := m.makeHash(key)
	elements := m.buckets[index]
	for nodeIndex, node := range elements {
		if node.key == key {
			node.value = val
			elements[nodeIndex] = node
			m.buckets[index] = elements
			return
		}
	}
	m.buckets[index] = append(m.buckets[index], node{key: key, value: val})
	if len(m.buckets[index]) > 8 {
		m.resize()
	}
}

func (m *MyMap) Get(key string) (int, bool) {
	index := m.makeHash(key)
	elements := m.buckets[index]
	for _, node := range elements {
		if node.key == key {
			return node.value, true
		}
	}

	return 0, false
}

func (m *MyMap) Delete(key string) {
	index := m.makeHash(key)
	elements := m.buckets[index]
	for nodeIndex, node := range elements {
		if node.key == key {
			elements = remove(elements, nodeIndex)
		}
	}
	m.buckets[index] = elements
}

func (m *MyMap) makeHash(key string) int {
	var byteSum int
	for _, char := range key {
		byteSum += int(char)
	}

	return byteSum % len(m.buckets)
}

func (m *MyMap) resize() {
	oldBuckets := m.buckets
	m.buckets = make([][]node, len(m.buckets)*2)
	for _, bucket := range oldBuckets {
		for _, node := range bucket {
			m.Set(node.key, node.value)
		}
	}
}

func remove(slice []node, s int) []node {
	return append(slice[:s], slice[s+1:]...)
}
