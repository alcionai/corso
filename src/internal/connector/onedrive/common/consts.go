package common

import "strings"

const (
	MetaFileSuffix    = ".meta"
	DirMetaFileSuffix = ".dirmeta"
	DataFileSuffix    = ".data"
)

func IsMetaFile(name string) bool {
	return strings.HasSuffix(name, MetaFileSuffix) || strings.HasSuffix(name, DirMetaFileSuffix)
}
