package _02_testify_测试框架

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestTesting(t *testing.T) {
	t.Log("start ...")
	as := assert.New(t)
	as.Equal(1, 2, "Equal")
	as.Equalf(1, 2, "Equalf, %s", "12")
	as.EqualError(errors.New("error"), "erro", "EqualError")
	as.EqualErrorf(errors.New("error"), "erro", "EqualErrorf: %d", 12)
	as.EqualValues([]int{1, 2}, []int{1, 3}, "EqualValues")
	as.EqualValuesf([]int{1, 2}, []int{1, 3}, "EqualValuesf: %d", 12)
	as.Error(errors.New("error"), "Error")
	as.Errorf(errors.New("error"), "Errorf: %d", 12)
	as.NotEqual(1, 2, "NotEqual")
	as.NotNil(nil, "Nil")
	as.IsType([]string{}, []string{})

	as.Contains("Hello World", "World")
	as.Contains([]string{"Hello", "World"}, "Hello")
	as.Contains([2]string{"Hello", "World"}, "Hello")
	as.Contains(map[string]string{"Hello": "World"}, "Hello")

	as.True(true)
	as.True(false)
	var s []string
	as.Empty(s)

	path, _ := filepath.Abs(".")
	as.DirExists(path+"/tmp", "DirExists")
	as.DirExistsf(path+"/tmp1", "DirExistsf %d", 12)
	t.Log("end ...")
}

//func TestCalculate(t *testing.T) {
//	as := assert.New(t)
//
//	var tests = []struct {
//		input    int
//		expected int
//	}{
//		{2, 4},
//		{-1, 1},
//		{0, 2},
//		{-5, -3},
//		{99999, 100001},
//	}
//
//	for _, test := range tests {
//		as.Equal(Calculate(test.input), test.expected)
//	}
//}
