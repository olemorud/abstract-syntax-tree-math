package main

import "fmt"

type mode int

const (
	operator mode = iota
	integer
	parentheses
	function
)

type Ast struct {
	t     mode
	val   int64
	left  *Ast
	right *Ast
}

func main() {

	testAst :=
		&Ast{operator, '*',
			&Ast{parentheses, 0,
				&Ast{operator, '+',
					&Ast{integer, 10, nil, nil},
					&Ast{integer, 5, nil, nil},
				},
				nil,
			},
			&Ast{integer, 3, nil, nil},
		}

	printTree(testAst)
	fmt.Println()
	solveTree(testAst)
	fmt.Printf("%+v\n", testAst.val)
}

func printTree(tree *Ast) {
	switch tree.t {
	case integer:
		fmt.Printf("%d", tree.val)
	case operator:
		if tree.left != nil {
			printTree(tree.left)
		}
		fmt.Printf("%c", tree.val)
		if tree.right != nil {
			printTree(tree.right)
		}
	case parentheses:
		fmt.Printf("(")
		if tree.left != nil {
			printTree(tree.left)
		}
		if tree.right != nil {
			printTree(tree.right)
		}
		fmt.Printf(")")
	}

}

// returns true if both child nodes are empty
func solveTree(tree *Ast) bool {
	// return true if leaf node (no children)
	if tree.left == nil && tree.right == nil {
		return true
	}

	var hasLeafChild bool = false

	if tree.left != nil {
		hasLeafChild = solveTree(tree.left) || hasLeafChild
	}

	if tree.right != nil {
		hasLeafChild = solveTree(tree.right) || hasLeafChild
	}

	if hasLeafChild || tree.t == parentheses {
		solveBranch(tree)
	}

	return false
}

func solveBranch(tree *Ast) {
	if tree.t == parentheses {
		tree.val = tree.left.val
		return
	}

	switch tree.val {
	case '+':
		tree.val = tree.left.val + tree.right.val
	case '*':
		tree.val = tree.left.val * tree.right.val
	}

	tree.t = integer
	tree.left = nil
	tree.right = nil
}
