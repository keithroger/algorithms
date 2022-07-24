package main

import "fmt"

type Graph map[string][]string

type Queue []string

func (s *Queue) IsEmpty() bool {
    if len(*s) == 0 {
        return true
    }

    return false
}

func (s *Queue) Push(elem string) {
    *s = append(*s, elem)
}

func (s *Queue) Pop() string{
    elem := (*s)[0]
    *s = (*s)[1:]

    return elem
}

func BFS(graph Graph, start string) {
    visited := make(map[string]bool)
    queue := Queue{}
    
    queue.Push(start)

    for !queue.IsEmpty() {
        currNode := queue.Pop()
        if visited[currNode] {
            continue
        }

        visited[currNode] = true

        fmt.Print(currNode, " ")

        for _, edge := range graph[currNode] {
            queue.Push(edge)
        }
    }
    fmt.Println()

}

func main() {
    graph := Graph{
        "a": []string{"b", "c"},
        "b": []string{"d"},
        "c": []string{"e", "b"},
        "e": []string{"a"},
    }

    BFS(graph, "a")
}
