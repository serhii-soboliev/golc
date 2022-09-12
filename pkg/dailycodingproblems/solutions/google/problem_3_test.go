package google_test

import (
	g "golc/pkg/dailycodingproblems/solutions/google"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestProblem2Serialize1(t *testing.T){
  left := g.TreeNode{2, nil, nil}
  right := g.TreeNode{3, nil, nil}
  root := g.TreeNode{1, &left, &right}
  c := g.Constructor()
  assert.Equal(t,"[[]2[]]1[[]3[]]", c.Serialize(&root))
}

func TestProblem2Serialize2(t *testing.T){
  left := g.TreeNode{2, &g.TreeNode{4, nil, nil}, nil}
  right := g.TreeNode{3, &g.TreeNode{6, nil, nil}, nil}
  root := g.TreeNode{1, &left, &right}
  c := g.Constructor()
  assert.Equal(t,"[[[]4[]]2[]]1[[[]6[]]3[]]", c.Serialize(&root))
}

func TestProblem21(t *testing.T){
  left := g.TreeNode{2, nil, nil}
  right := g.TreeNode{3, nil, nil}
  root := g.TreeNode{1, &left, &right}
  c := g.Constructor()
  dRoot := c.Deserialize(c.Serialize(&root))
  assert.Equal(t, 1, dRoot.Val)
  assert.Equal(t, 2, dRoot.Left.Val)
  assert.Equal(t, 3, dRoot.Right.Val)
}

func TestProblem22(t *testing.T){
  left := g.TreeNode{2, &g.TreeNode{4, nil, nil}, nil}
  right := g.TreeNode{3, &g.TreeNode{6, nil, nil}, nil}
  root := g.TreeNode{1, &left, &right}
  c := g.Constructor()
  dRoot := c.Deserialize(c.Serialize(&root))
  assert.Equal(t, 1, dRoot.Val)
  assert.Equal(t, 2, dRoot.Left.Val)
  assert.Equal(t, 3, dRoot.Right.Val)
  assert.Equal(t, 4, dRoot.Left.Left.Val)
  assert.Nil(t, dRoot.Left.Right)
  assert.Equal(t, 4, dRoot.Left.Left.Val)
  assert.Equal(t, 6, dRoot.Right.Left.Val)
  assert.Nil(t, dRoot.Right.Right)
}