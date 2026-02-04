# Data Structures and Algorithms - Homework
[![CI](https://github.com/hafidbuilds/dsa/actions/workflows/ci.yml/badge.svg)](https://github.com/hafidbuilds/dsa/actions/workflows/ci.yml)

A Go library for learning data structures through implementation.

## Test Status

<!-- STATUS:START -->
| # | Package | Structure | Status |
|---|---|---|---|
| 1 | `dynamicarray` | `DynamicArray[T]` | ✅ |
| 2 | `linkedlist` | `SinglyLinkedList[T]` | ✅ |
| 3 | `linkedlist` | `DoublyLinkedList[T]` | ✅ |
| 4 | `linkedlist` | `CircularDoublyLinkedList[T]` | ✅ |
| 5 | `stack` | `Stack[T]` | ✅ |
| 6 | `queue` | `Queue[T]` | ✅ |
| 7 | `hashmap` | `HashMap[K,V]` | ✅ |
| 8 | `sets` | `HashSet[T]` | ✅ |
| 9 | `graph` | `Graph[V]` | ❌ |
| 10 | `tree` | `BinarySearchTree[T]` | ✅ |
| 11 | `bitsets` | `BitSet` | ❌ |
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
