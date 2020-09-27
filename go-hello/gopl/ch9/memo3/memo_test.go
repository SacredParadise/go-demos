package memo_test

import (
	memo "gopl/ch9/memo3"
	"gopl/ch9/memotest"
	"testing"
)


var httpGetBody = memotest.HttpGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}