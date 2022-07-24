package main

import "fmt"

type Graph map[string][]string

type Stack []string

func (s *Stack) IsEmpty() bool {
    if len(*s) == 0 {
        return true
    }

    return false
}

func (s *Stack) Push(elem string) {
    *s = append(*s, elem)
}

func (s *Stack) Pop() string{
    n := len(*s)
    elem := (*s)[n-1]
    *s = (*s)[:n-1]

    return elem
}

func DFS(graph Graph, start string) {
    visited := make(map[string]bool)
    stack := Stack{}
    
    stack.Push(start)

    for !stack.IsEmpty() {
        currNode := stack.Pop()
        if visited[currNode] {
            continue
        }

        visited[currNode] = true

        fmt.Print(currNode, " ")

        for _, edge := range graph[currNode] {
            stack.Push(edge)
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

    DFS(graph, "a")

}
