package common

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
)

func BuildFilepathSanitree(
	ctx context.Context,
	rootDir string,
) *Sanitree[fs.FileInfo] {
	var root *Sanitree[fs.FileInfo]

	walker := func(
		path string,
		info os.FileInfo,
		err error,
	) error {
		if err != nil {
			Fatal(ctx, "param in filepath walker", err)
		}

		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			Fatal(ctx, "getting relative filepath", err)
		}

		if root == nil {
			root = &Sanitree[fs.FileInfo]{
				Self:     info,
				ID:       info.Name(),
				Name:     info.Name(),
				Leaves:   map[string]*Sanileaf[fs.FileInfo]{},
				Children: map[string]*Sanitree[fs.FileInfo]{},
			}

			return nil
		}

		node := root.NodeAt(ctx, relPath)

		if info.IsDir() {
			node.Children[info.Name()] = &Sanitree[fs.FileInfo]{
				Parent:   node,
				Self:     info,
				ID:       info.Name(),
				Name:     info.Name(),
				Leaves:   map[string]*Sanileaf[fs.FileInfo]{},
				Children: map[string]*Sanitree[fs.FileInfo]{},
			}
		} else {
			node.Leaves[info.Name()] = &Sanileaf[fs.FileInfo]{
				Parent: node,
				Self:   info,
				ID:     info.Name(),
				Name:   info.Name(),
				Size:   info.Size(),
			}
		}

		return nil
	}

	err := filepath.Walk(rootDir, walker)
	if err != nil {
		Fatal(ctx, "walking filepath", err)
	}

	return root
}
