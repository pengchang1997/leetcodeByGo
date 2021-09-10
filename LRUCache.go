package leetcodeByGo

// LRU缓存机制

import (
	"container/list"
)

// 链表节点定义
type node struct {
	key   int
	value int
	next  *node
}

// cache定义
type LRUCache struct {
	// cache容量
	capacity int

	// cache大小
	size int

	// 双向链表
	myList *list.List

	// 记录key与链表节点映射关系的哈希表
	m map[int]*list.Element
}

// cache的构造函数定义
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		size:     0,
		myList:   list.New(),
		m:        make(map[int]*list.Element),
	}
}

// 实现cache查询操作
func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.myList.MoveToFront(v)
		return v.Value.(*node).value
	}

	return -1
}

// 实现cache的新增与修改操作
func (this *LRUCache) Put(key int, value int) {
	// 如果key存在于cache中
	if v, ok := this.m[key]; ok {
		node := v.Value.(*node)
		node.value = value
		this.myList.MoveToFront(v)
	} else {
		// 如果key不存在于cache中

		// 首先新建一个节点
		newNode := &node{
			key:   key,
			value: value,
		}

		// 如果cache剩余空间不足，则需要淘汰最久未被访问的元素
		if this.size == this.capacity {
			oldest := this.myList.Back()
			this.myList.Remove(oldest)
			delete(this.m, oldest.Value.(*node).key)

			newElement := this.myList.PushFront(newNode)
			this.m[key] = newElement
		} else {
			newElement := this.myList.PushFront(newNode)
			this.m[key] = newElement
			this.size++
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
