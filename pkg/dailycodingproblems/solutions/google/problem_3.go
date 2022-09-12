package google

import (
	"fmt"
	"strconv"
)

/*
This problem was asked by Google.

Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

For example, given the following Node class

class TreeNode {
  constructor(val, left = null, right = null) {
    this.val = val;
    this.left = left;
    this.right = right;
  }
}
The following test should pass:

const node = new TreeNode(
  'root',
  new TreeNode('left', new TreeNode('left.left'), new TreeNode('right'))
);
expect(deserialize(serialize(node)).left.left.val).toEqual('left.left'); // Jest Testing

Leetcode
297. Serialize and Deserialize Binary Tree
https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	result := ""
	if root == nil {
		return result
	}
	result = fmt.Sprintf("[%s]%d[%s]", c.serialize(root.Left), root.Val, c.serialize(root.Right))
	return result
}

func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" || data == "[]" {
		return nil
	}
	if data[0] != '[' {
		v, err := strconv.Atoi(data)
		if err != nil {
			panic("Invalid string")
		}
		return &TreeNode{v, nil, nil}
	}
	cBIdx := findClosingBracketIndex(data, 0)
	if cBIdx == len(data)-1 {
		return c.deserialize(data[1 : len(data)-1])
	}
	oBIdx := findSymbolStartingFromIndex(data, '[', cBIdx)
	left := data[1:cBIdx]
	center := data[cBIdx+1 : oBIdx]
	right := data[oBIdx:]
	root := c.deserialize(center)
	root.Left = c.deserialize(left)
	root.Right = c.deserialize(right)
	return root

}

func findSymbolStartingFromIndex(s string, c byte, startIndex int) int {
	for i := startIndex; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	panic("Shouldnt get there")
}

func findClosingBracketIndex(s string, openIdx int) int {
	openBracketsCount := 1
	for i := openIdx + 1; i < len(s); i++ {
		if s[i] == '[' {
			openBracketsCount += 1
		} else if s[i] == ']' {
			openBracketsCount -= 1
			if openBracketsCount == 0 {
				return i
			}
		}
	}
	panic("Shouldnt get there")

}

func (c *Codec) Serialize(root *TreeNode) string {
	return c.serialize(root)
}

func (c *Codec) Deserialize(data string) *TreeNode {
	return c.deserialize(data)
}
