package main

import "fmt"

const ArraySize = 7

// HashTable (array + linkedList)
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket 桶
type bucket struct {
	head *bucketNode
}

// bucketNode 桶node
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert 插入
func (t *HashTable) Insert(key string) {
	h := hash(key)
	t.array[h].insert(key)
}

// Search 搜索
func (t *HashTable) Search(key string) bool {
	// 根据hash算法得到key所在的index
	h := hash(key)
	// 通过index获取到使用哪个bucket桶来遍历查找
	return t.array[h].search(key)
}

// Delete 删除
func (t *HashTable) Delete(key string) {
	h := hash(key)
	t.array[h].delete(key)
}

// hash 算法
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// insert 其实bucket 就是使用 linkedlist 来实现的
// 往链表中插入数据都是下面的套路，这里不关心bucket的length
func (b *bucket) insert(key string) {
	second := b.head
	b.head = &bucketNode{key: key}
	b.head.next = second
}

// search 某个桶下找key是否存在
func (b *bucket) search(key string) bool {
	// 如果桶下面没有node直接返回false
	if b.head == nil {
		return false
	}
	// 如果桶下有node，检查是否存在key，这里直接while loop
	for b.head.key != key {
		// 当前node不相等，如果没有下一个node则直接返回false，没有找到
		if b.head.next == nil {
			return false
		}
		// 将指向当前node的指针移动到下一个node，再次循环查找
		b.head = b.head.next
	}
	// 如果找到了则返回true
	return true
}

// delete
func (b *bucket) delete(key string) {
	if b.head == nil {
		return
	}
	if b.head.key == key {
		b.head = b.head.next
		return
	}

	for b.head.key != key {
		if b.head.next == nil {
			return
		}
		b.head = b.head.next
	}
	b.head.next = b.head.next.next
}

// Init 初始化,给hashtable array 给个空桶
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	fmt.Println("哈希表简单例子")

	h := Init()
	fmt.Println(h)
	h.Insert("Stand")
	h.Insert("Envy")
	h.Insert("Hello")
	h.Insert("World")
	h.Insert("Wade")
	h.Insert("Chan")
	fmt.Println(h.Search("Envy"))
	h.Delete("Hello")
	fmt.Println(h.Search("Hello"))
}
