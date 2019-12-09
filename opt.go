package main

import "flag"

// opt
type opt struct {
	a
	t
}

func newOpt() *opt {
	return &opt{
		a: a{
			bool: flag.Bool("a", false, "Include directory entries whose names begin with a dot (.)."),
		},
		t: t{
			bool: flag.Bool("t", false, "Sort by time modified (most recently modified first) before sorting the operands by lexicographical order."),
		},
	}
}
