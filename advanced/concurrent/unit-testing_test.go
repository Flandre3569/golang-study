package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func HelloTom() string {
	return "Jerry"
}

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	//if output != expectOutput {
	//	t.Errorf("Expected %s do not match actual %s", expectOutput, output)
	//}

	// 使用断言实现单测
	assert.Equal(t, output, expectOutput)
}
