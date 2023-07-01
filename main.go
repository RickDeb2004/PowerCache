//special features: 1.Timestamp
//                  2.load metrics 
//                  3.

package main

import (
	"fmt"
	"time"
)

const Size = 5

type Node struct {
	Val       string
	Left      *Node
	Right     *Node
	TimeStamp time.Time
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue      Queue
	Hash       Hash
	HitCount   int //represents number of cache hits , it is increamented eachtime as existing entry is found in cache
	TotalCount int //It represents the total number of requests made to the cache, including both hits and misses.
	MissCount  int //total - hit
	Size int
}
type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}
func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail}

}
func (c *Cache) Check(str string, expiration time.Duration) {
	c.TotalCount++
	node := &Node{}
	if val, ok := c.Hash[str]; ok {
		c.HitCount++
		node = c.Remove(val)

	} else {
		c.MissCount++
		node = &Node{Val: str, TimeStamp: time.Now().Add(expiration)}
	}
	c.Add(node)
	c.Hash[str] = node

}
func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove:%s\n", n.Val)
	left := n.Left
	right := n.Right
	right.Left = left
	left.Right = right
	c.Queue.Length -= 1
	delete(c.Hash, n.Val)
	return n

}
func (c *Cache) Add(n *Node) *Node {
	fmt.Printf("add:%s\n", n.Val)
	tmp := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n
	c.Queue.Length++
	if c.Queue.Length > Size {
		c.Remove(c.Queue.Tail.Left)
	}
	return n
}
func (c *Cache) Display() {

	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d-[", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<-->")

		}
		node = node.Right
	}
	fmt.Println("]")
}
func (c *Cache) RemoveExpired() {
	currentTime := time.Now()
	node := c.Queue.Tail.Left
	for node != c.Queue.Head && node.TimeStamp.Before(currentTime) {
		NextNode := node.Left
		c.Remove(node)
		node = NextNode
	}
}
func (c *Cache) GithitRate() float64 {
	if c.TotalCount == 0 {
		return 0
	}
	return float64(c.HitCount) / float64(c.TotalCount) * 100
}
func main() {
	fmt.Println("Start Cache")
	expiration := 2 * time.Second //Set expiration time to 2 seconds
	cache := NewCache()
	for _, word := range []string{"parrot", "avocardo", "tree", "potato", "tree"} {
		cache.Check(word, expiration)
		cache.Display()
	}
	time.Sleep(3 * time.Second) //Sleep for 3 seconds to allow some entries to expire
	cache.RemoveExpired()       //Remove expired entries from the cache
	cache.Display()
	fmt.Printf("HitRate:%.2f%%\n", cache.GithitRate())
	fmt.Printf("MissRate:%.2f%%\n", 100-cache.GithitRate())
	fmt.Printf("TotalRate:%.2f%%\n", cache.TotalCount)

}
