package hashmap_test

import (
	hm "golc/pkg/leetcode/hashmap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate1(t *testing.T) {
	res := hm.FindDuplicate([]string{"root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)",
													"root/c/d 4.txt(efgh)","root 4.txt(efgh)"})
	assert.Contains(t, res, []string{"root/a/2.txt","root/c/d/4.txt","root/4.txt"},
							[]string{"root/a/1.txt","root/c/3.txt"})
}

func TestFindDuplicate2(t *testing.T) {
	res := hm.FindDuplicate([]string{"root/a 1.txt(abcd) 2.txt(efgh)",
	"root/c 3.txt(abcd)","root/c/d 4.txt(efgh)"})
	assert.Contains(t, res, []string{"root/a/2.txt","root/c/d/4.txt"},
							[]string{"root/a/1.txt","root/c/3.txt"})
}