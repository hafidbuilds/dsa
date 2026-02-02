# Data Structures and Algorithms - Homework
[![CI](https://github.com/hafidbuilds/dsa/actions/workflows/ci.yml/badge.svg)](https://github.com/hafidbuilds/dsa/actions/workflows/ci.yml)

A Go library for learning data structures through implementation.

## Test Status

<!-- STATUS:START -->
| # | Package | Structure | Status |
|---|---|---|---|
| 1 | `dynamicarray` | `DynamicArray[T]` | PASS |
| 2 | `linkedlist` | `SinglyLinkedList[T]` | PASS |
| 3 | `linkedlist` | `DoublyLinkedList[T]` | PASS |
| 4 | `linkedlist` | `CircularDoublyLinkedList[T]` | PASS |
| 5 | `stack` | `Stack[T]` | PASS |
| 6 | `queue` | `Queue[T]` | PASS |
| 7 | `hashmap` | `HashMap[K,V]` | PASS |
| 8 | `sets` | `HashSet[T]` | PASS |
| 9 | `graph` | `Graph[V]` | FAIL |
| 10 | `tree` | `BinarySearchTree[T]` | FAIL |
| 11 | `bitsets` | `BitSet` | FAIL |
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
