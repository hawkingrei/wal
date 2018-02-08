package main

import (
	"fmt"

	"github.com/hawkingrei/wal/fileutil"
)

func Exist(dirpath string) bool {
	names, err := fileutil.ReadDir(dirpath)
	if err != nil {
		return false
	}
	return len(names) != 0
}

func walName(seq, index uint64) string {
	return fmt.Sprintf("%016x-%016x.wal", seq, index)
}
