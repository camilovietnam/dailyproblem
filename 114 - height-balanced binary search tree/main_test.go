package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OneElement(t *testing.T) {
	myArr := []int{1}
	want := &node{
		Value: 1,
	}

	tree := arrayToTree(myArr)
	assert.Equal(t, want, tree, "Failed at One Element")
}

func Test_TwoElements(t *testing.T) {
	myArr := []int{3,5}
	want := &node{
		Value: 5,
		Left: &node{
			Value: 3,
		},
	}

	tree := arrayToTree(myArr)
	assert.Equal(t, want, tree, "Failed at Two Elements")
}

func Test_ThreeElements(t *testing.T) {
	myArr := []int{1,2,3}
	want := &node{
		Value: 2,
		Left: &node{
			Value: 1,
		},
		Right: &node {
			Value: 3,
		},
	}

	tree := arrayToTree(myArr)
	assert.Equal(t, want, tree, "Failed at Three Elements")
}

func Test_FourElements(t *testing.T) {
	myArr := []int{1,2,3,4}
	want := &node{
		Value: 3,
		Left: &node{
			Value: 2,
			Left: &node {
				Value: 1,
			},
		},
		Right: &node {
			Value: 4,
		},
	}

	tree := arrayToTree(myArr)
	assert.Equal(t, want, tree, "Failed at Four Elements")
}

func Test_FiveElements(t *testing.T) {
	myArr := []int{6,7,8,9,10}
	want := &node{
		Value: 8,
		Left: &node{
			Value: 7,
			Left: &node{
				Value: 6,
			},
		},
		Right: &node{
			Value: 10,
			Left: &node{
				Value: 9,
			},
		},
	}

	tree := arrayToTree(myArr)
	assert.Equal(t, want, tree, "Failed at Five Elements")
}

func Test_TenElements(t *testing.T) {
	myArr := []int{2,4,6,8,20,30,100,1000,2000,5000}
	want := &node{
		Value: 30,
		Left: &node{
			Value: 6,
			Left: &node{
				Value: 4,
				Left: &node {
					Value: 2,
				},
			},
			Right: &node{
				Value: 20,
				Left: &node{
					Value: 8,
				},
			},
		},
		Right: &node{
			Value: 2000,
			Left: &node{
				Value: 1000,
				Left: &node{
					Value: 100,
				},
			},
			Right: &node{
				Value: 5000,
			},
		},
	}

	tree := arrayToTree(myArr)
	assert.Equal(t, want, tree, "Failed at Five Elements")
}