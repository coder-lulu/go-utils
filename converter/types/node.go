package types

// TreeNode represents a generic tree node structure
type TreeNode[T any] struct {
	Data     T
	Children []*TreeNode[T]
}

// IDGetter defines a function type that gets an ID from a data item
type IDGetter[T any] func(item T) int
