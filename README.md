# Server Name Generator

## Description
Suppose you’re running a pool of servers where each server type is numbered sequentially from 1.

Sometimes, we’ll want to create a new server or decommission an old one.

Write two methods, allocate(host_type) and deallocate(hostname).

The former should reserve and return a new hostname, the latter should release the specified hostname.

## Example
in `main.go` you have

```
servers := make(map[string][]int)
fmt.Println(allocate("storage", servers))
fmt.Println(allocate("storage", servers))
fmt.Println(deallocate("storage-1", servers))
fmt.Println(allocate("storage", servers))
fmt.Println(allocate("metrics", servers))
```

and when you run the app using `go run .` you get

```
storage-1
storage-2
storage-1
metrics-1
```

## Notes
If you deallocate a server that does not exist, you'll get an error
