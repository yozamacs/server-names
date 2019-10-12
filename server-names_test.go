package main

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

var smallestNumTests = []struct {
	in  []int
	out int
}{
	{[]int{1, 2, 3}, 4},
	{[]int{1, 3}, 2},
	{[]int{4, 3, 1, 4}, 2},
	{[]int{2, 3, 5, 1}, 4},
	{[]int{4, 5, 3}, 1},
	{[]int{1}, 2},
	{[]int{1, 9999}, 2},
	{[]int{}, 1},
}

func TestSmallestNum(t *testing.T) {
	g := NewGomegaWithT(t)
	for _, smt := range smallestNumTests {
		g.Expect(smallestNum(smt.in)).To(Equal(smt.out))
	}
}

func TestAllocate(t *testing.T) {
	g := NewGomegaWithT(t)
	servers := make(map[string][]int)
	newServer := allocate("storage", servers)
	g.Expect(newServer).To(Equal("storage-1"))
}

func TestAllocateAndDeallocate(t *testing.T) {
	g := NewGomegaWithT(t)
	servers := make(map[string][]int)
	server1 := allocate("storage", servers)
	g.Expect(server1).To(Equal("storage-1"))
	server2 := allocate("storage", servers)
	g.Expect(server2).To(Equal("storage-2"))
	fmt.Printf("servers: %+v\n", servers)
	g.Expect(deallocate("storage-1", servers)).To(BeNil())
	newServer := allocate("storage", servers)
	g.Expect(newServer).To(Equal("storage-1"))
}

func TestDeallocateServerNotFound(t *testing.T) {
	g := NewGomegaWithT(t)
	servers := make(map[string][]int)
	g.Expect(deallocate("storage-1", servers)).To(MatchError("server name does not exist"))
}
