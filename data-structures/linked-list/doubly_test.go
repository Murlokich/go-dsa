package linked_list_test

import (
	"github.com/Murlokich/go-dsa.git/data-structures/linked-list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type listTestCase struct {
	name   string
	values []int
}

var commonListTests = []listTestCase{
	{name: "empty", values: []int{}},
	{name: "nil slice", values: nil},
	{name: "single value", values: []int{42}},
	{name: "multiple values", values: []int{10, 20, 30}},
}

func TestDoublyLinkedList_GetHead(t *testing.T) {
	for _, tc := range commonListTests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			linkedList := linked_list.NewDoublyLinkedList(tc.values...)
			head, err := linkedList.GetHead()

			if len(tc.values) == 0 {
				assert.ErrorIs(t, err, linked_list.ErrEmptyList)
				assert.Zero(t, head)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.values[0], head)
			}
		})
	}
}

func TestDoublyLinkedList_GetTail(t *testing.T) {
	for _, tc := range commonListTests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			linkedList := linked_list.NewDoublyLinkedList(tc.values...)
			tail, err := linkedList.GetTail()

			if len(tc.values) == 0 {
				assert.ErrorIs(t, err, linked_list.ErrEmptyList)
				assert.Zero(t, tail)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.values[len(tc.values)-1], tail)
			}
		})
	}
}

func TestDoublyLinkedList_IsEmpty(t *testing.T) {
	for _, tc := range commonListTests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			linkedList := linked_list.NewDoublyLinkedList(tc.values...)

			if len(tc.values) == 0 {
				assert.True(t, linkedList.IsEmpty())
			} else {
				assert.False(t, linkedList.IsEmpty())
			}
		})
	}
}

func TestNewDoublyLinkedList(t *testing.T) {
	for _, tc := range commonListTests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			linkedList := linked_list.NewDoublyLinkedList[int](tc.values...)
			listValues := make([]int, 0, len(tc.values))

			for !linkedList.IsEmpty() {
				value, err := linkedList.GetHead()
				require.NoError(t, err)
				listValues = append(listValues, value)
				err = linkedList.DeleteHead()
				require.NoError(t, err)
			}
			// we expect listValues to be empty for nil slice passed
			if len(tc.values) == 0 {
				assert.Empty(t, listValues)
			} else {
				assert.Equal(t, tc.values, listValues)
			}
		})
	}
}
