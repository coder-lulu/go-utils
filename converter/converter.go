package converter

import (
	"github.com/coder-lulu/go-utils/converter/types"
)

// ArrayToTreeWithFields converts a slice into a tree structure using specified ID field getters
func ArrayToTree[T any](
	items []T,
	getID types.IDGetter[T],
	getParentID types.IDGetter[T],
) []*types.TreeNode[T] {
	if getID == nil {
		// Default to assuming there are 'ID' and 'ParentID' fields
		getID = defaultIDGetter[T]
	}
	if getParentID == nil {
		getParentID = defaultParentIDGetter[T]
	}

	// Create a map to store nodes by their ID for quick lookup
	nodeMap := make(map[int]*types.TreeNode[T])
	
	// First pass: create all nodes
	for i := range items {
		node := &types.TreeNode[T]{
			Data:     items[i],
			Children: make([]*types.TreeNode[T], 0),
		}
		nodeMap[getID(items[i])] = node
	}
	
	// Second pass: build the tree structure
	var roots []*types.TreeNode[T]
	for _, item := range items {
		node := nodeMap[getID(item)]
		parentID := getParentID(item)
		
		if parent, exists := nodeMap[parentID]; exists && getID(item) != parentID {
			parent.Children = append(parent.Children, node)
		} else {
			roots = append(roots, node)
		}
	}
	
	return roots
}

// defaultIDGetter tries to get an 'ID' field using reflection
func defaultIDGetter[T any](item T) int {
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Struct {
		if idField := v.FieldByName("ID"); idField.IsValid() {
			return int(idField.Int())
		}
	}
	return 0
}

// defaultParentIDGetter tries to get a 'ParentID' field using reflection
func defaultParentIDGetter[T any](item T) int {
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Struct {
		if parentIDField := v.FieldByName("ParentID"); parentIDField.IsValid() {
			return int(parentIDField.Int())
		}
	}
	return 0
}
