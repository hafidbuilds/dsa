#!/usr/bin/env bash
set -u

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$root_dir"

if ! grep -q "<!-- STATUS:START -->" README.md || ! grep -q "<!-- STATUS:END -->" README.md; then
  echo "README status markers not found." >&2
  exit 1
fi

rows=()

run_test() {
  local index="$1"
  local pkg="$2"
  local structure="$3"
  shift 3
  "$@"
  local status="$?"
  local result="FAIL"
  if [ "$status" -eq 0 ]; then
    result="PASS"
  fi
  rows+=("| ${index} | \`${pkg}\` | \`${structure}\` | ${result} |")
  return 0
}

run_test 1 dynamicarray "DynamicArray[T]" go test ./dynamicarray
run_test 2 linkedlist "SinglyLinkedList[T]" go test ./linkedlist -run '^TestSinglyLinkedList$'
run_test 3 linkedlist "DoublyLinkedList[T]" go test ./linkedlist -run '^TestDoublyLinkedList$'
run_test 4 linkedlist "CircularDoublyLinkedList[T]" go test ./linkedlist -run '^TestCircularDoublyLinkedList'
run_test 5 stack "Stack[T]" go test ./stack
run_test 6 queue "Queue[T]" go test ./queue
run_test 7 hashmap "HashMap[K,V]" go test ./hashmap
run_test 8 sets "HashSet[T]" go test ./sets
run_test 9 graph "Graph[V]" go test ./graph
run_test 10 tree "BinarySearchTree[T]" go test ./tree
run_test 11 bitsets "BitSet" go test ./bitsets

table_header="| # | Package | Structure | Status |"$'\n'"|---|---|---|---|"
table_rows="$(printf "%s\n" "${rows[@]}")"
table="${table_header}"$'\n'"${table_rows}"

tmp_file="$(mktemp)"
awk -v table="$table" '
  $0 ~ /<!-- STATUS:START -->/ {
    print
    print table
    in_block=1
    next
  }
  $0 ~ /<!-- STATUS:END -->/ {
    in_block=0
    print
    next
  }
  in_block == 1 { next }
  { print }
' README.md > "$tmp_file"

mv "$tmp_file" README.md
