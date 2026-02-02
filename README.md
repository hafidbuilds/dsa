# Data Structures and Algorithms - Homework
[![CI](https://github.com/hafidbuilds/dsa/actions/workflows/ci.yml/badge.svg)](https://github.com/hafidbuilds/dsa/actions/workflows/ci.yml)

A Go library for learning data structures through implementation.

## Test Status

<!-- STATUS:START -->
| #  | Package        | Structure                     | Status  | Key Concepts                   |
|----|----------------|-------------------------------|---------|--------------------------------|
| 1  | `dynamicarray` | `DynamicArray[T]`             | PENDING | Amortized growth, index access |
| 2  | `linkedlist`   | `SinglyLinkedList[T]`         | PENDING | Node pointers, traversal       |
| 3  | `linkedlist`   | `DoublyLinkedList[T]`         | PENDING | Bidirectional links            |
| 4  | `linkedlist`   | `CircularDoublyLinkedList[T]` | PENDING | Circular references            |
| 5  | `stack`        | `Stack[T]`                    | PENDING | LIFO, backend abstraction      |
| 6  | `queue`        | `Queue[T]`                    | PENDING | FIFO, backend abstraction      |
| 7  | `hashmap`      | `HashMap[K,V]`                | PENDING | Hashing, collision resolution  |
| 8  | `sets`         | `HashSet[T]`                  | PENDING | Set operations, uniqueness     |
| 9  | `graph`        | `Graph[V]`                    | PENDING | Adjacency list, traversal      |
| 10 | `tree`         | `BinarySearchTree[T]`         | PENDING | Tree structure, ordering       |
| 11 | `bitsets`      | `BitSet`                      | PENDING | Bit manipulation (optional)    |
<!-- STATUS:END -->

## Philosophy

This library uses small, composable interfaces defined in `adt/adt.go`. 
Each method you implement satisfies one or more interfaces. 

Tests in `adt/prop/` validate your implementations against these contracts.

## Getting Started

```bash
# run all tests (most will fail initially)
go test ./...

# run tests for a specific package
go test ./dynamicarray/...
go test ./linkedlist/...
```

## Homework Order

Complete the data structures in this order. Each builds on concepts from the previous.
