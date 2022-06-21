package main

import "fmt"

type Tries struct {
	root *Node
}

type Node struct {
	children [26]*Node
	isEnd    bool
}

func Init() *Tries {
	return &Tries{root: &Node{}}
}

// insert 往Tries中插入数据
func (t *Tries) insert(w string) {
	wordsCount := len(w)
	currentNode := t.root
	for i := 0; i < wordsCount; i++ {
		// 获取每个字符当前的索引
		currentIndex := w[i] - 'a'
		// 如果不存在该字符索引 则新增Node
		if currentNode.children[currentIndex] == nil {
			currentNode.children[currentIndex] = &Node{}
		}
		// 继续
		currentNode = currentNode.children[currentIndex]
	}
	currentNode.isEnd = true
}

// search 搜索
func (t *Tries) search(w string) bool {
	wordsCount := len(w)
	currentNode := t.root
	for i := 0; i < wordsCount; i++ {
		currentIndex := w[i] - 'a'
		if currentNode.children[currentIndex] == nil {
			return false
		}
		currentNode = currentNode.children[currentIndex]
	}

	return currentNode.isEnd

}

func main() {
	fmt.Println("Tries")

	r := Init()
	toAdd := []string{
		"chenkuan",
		"chenking",
		"lichao",
		"lck",
		"argon",
		"echo",
		"ahke",
		"eooo",
	}
	for _, s := range toAdd {
		r.insert(s)
	}
	fmt.Println(r.search("chenking"))
}
