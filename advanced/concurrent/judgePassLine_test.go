package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudgePassLine(t *testing.T) {
	isPass := JudgePassLine(70)
	assert.Equal(t, true, isPass)
}
