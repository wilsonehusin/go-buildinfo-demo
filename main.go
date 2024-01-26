package main

import (
	"fmt"
	"runtime/debug"
)

type vcsInfo struct {
	time     string
	revision string
	modified string
}

func (v *vcsInfo) Print() {
	fmt.Printf("%s\n%s", v.time, v.revision)
	if v.modified == "true" {
		fmt.Print("-dirty")
	}
	fmt.Print("\n")
}

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		panic("cannot retrieve buildinfo")
	}

	v := &vcsInfo{}
	for _, i := range bi.Settings {
		switch i.Key {
		case "vcs.time":
			v.time = i.Value
		case "vcs.revision":
			v.revision = i.Value
		case "vcs.modified":
			v.modified = i.Value
		default:
			continue
		}
	}
	v.Print()
}
