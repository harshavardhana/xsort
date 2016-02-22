package usort

import (
	"sort"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestUniqueSort(c *C) {
	// Test with 2 duplicate entries.
	intSlice := []int{1, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 9, 10, 10}
	n := uniqueSort(sort.IntSlice(intSlice))

	uniqIntSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c.Assert(intSlice[0:n], DeepEquals, uniqIntSlice)

	// Test with multiple duplicate entries.
	intSlice = []int{1, 1, 2, 3, 4, 5, 5, 5, 5, 5, 6, 7, 8, 9, 9, 10, 10}
	n = uniqueSort(sort.IntSlice(intSlice))
	c.Assert(intSlice[0:n], DeepEquals, uniqIntSlice)
}

func (s *MySuite) TestUniqueStrings(c *C) {
	// Test with duplicate string entries.
	strSlice := []string{"test", "test", "test", "test", "test", "hello", "hello", "world", "noop", "foo"}
	uniqStrSlice := []string{"foo", "hello", "noop", "test", "world"}
	c.Assert(uniqStrSlice, DeepEquals, UniqueStrings(strSlice))
}
