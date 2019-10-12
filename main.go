package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	servers := make(map[string][]int)
	fmt.Println(allocate("storage", servers))
	fmt.Println(allocate("storage", servers))
	deallocate("storage-1", servers)
	fmt.Println(allocate("storage", servers))
	fmt.Println(allocate("metrics", servers))
}

func smallestNum(arr []int) int {
	numMaps := make(map[int]bool)
	for _, num := range arr {
		numMaps[num] = true
	}
	bigInt := 10000000
	for x := 1; x < bigInt; x++ {
		if !numMaps[x] {
			return x
		}
	}
	return -1
}

func allocate(name string, servers map[string][]int) string {
	if _, ok := servers[name]; !ok {
		serverId := []int{}
		servers[name] = serverId
	}
	newId := smallestNum(servers[name])
	servers[name] = append(servers[name], newId)
	newName := fmt.Sprintf("%v-%v", name, newId)
	return newName
}

func deallocate(name string, servers map[string][]int) error {
	nameArray := strings.Split(name, "-")
	serverName := nameArray[0]
	serverId := nameArray[1]
	serverIdInt, err := strconv.Atoi(serverId)
	if err != nil {
		panic(err)
	}
	if _, ok := servers[serverName]; !ok {
		return fmt.Errorf("server name does not exist")
	}
	serverIds := servers[serverName]
	newIds := []int{}
	for _, id := range serverIds {
		if id != serverIdInt {
			newIds = append(newIds, id)
		}
	}
	servers[serverName] = newIds
	return nil
}
