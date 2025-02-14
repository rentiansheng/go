package routine

import (
	_ "unsafe" // for go:linkname
)

/***************************
    @author: tiansheng.ren
    @date: 2/14/25
    @desc:


***************************/

type goroutineDoneFn func() <-chan struct{}

//go:linkname Id runtime.GoID
func Id() uint64

//go:linkname Split runtime.GoSplit
func Split()

//go:linkname Value runtime.GoValue
func Value(key string) any

//go:linkname Set runtime.GoSetValue
func Set(key string, v any)

//go:linkname WithDone runtime.GoWithDone
func WithDone(fn goroutineDoneFn)

//go:linkname IsDone runtime.GoIsDone
func IsDone() bool
