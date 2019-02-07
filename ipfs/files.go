package ipfs

import (
	"bytes"
	"github.com/ipfs/go-ipfs-files"
	"strconv"
)

//  Create files from duras and structure files with directory hierarchy,
// where each directory has max `dirMaxWidth` entries, using top-bottom breakdown
//
//  Returns root directory
func getAsDirHierarchy(duras []string, dirMaxWidth int) files.Directory {

	// construct dir
	if len(duras) <= dirMaxWidth {
		return constructDir(duras)
	}

	dirsAsEntries := []files.DirEntry{}
	dirs := getChildDirectories(duras, dirMaxWidth)

	for i, dir := range dirs {
		dirsAsEntries = append(dirsAsEntries, files.FileEntry(strconv.Itoa(i), dir))
	}

	return files.NewSliceDirectory(dirsAsEntries)
}

func constructDir(duras []string) files.Directory {

	entries := []files.DirEntry{}
	for i, dura := range duras {
		file := files.NewBytesFile(bytes.NewBufferString(dura).Bytes())
		entries = append(entries, files.FileEntry(strconv.Itoa(i), file))
	}
	return files.NewSliceDirectory(entries)
}

func getChildDirectories(duras []string, dirMaxWidth int) []files.Directory {

	i := 0
	dirs := []files.Directory{}
	subdirEntriesNum := len(duras)/dirMaxWidth + 1
	for i < len(duras) {
		if len(duras)-i < subdirEntriesNum {
			dirs = append(dirs, getAsDirHierarchy(duras[i:], dirMaxWidth))
		} else {
			dirs = append(dirs, getAsDirHierarchy(duras[i:i+subdirEntriesNum], dirMaxWidth))
		}
		i += subdirEntriesNum
	}
	return dirs
}
