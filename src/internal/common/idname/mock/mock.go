package mock

import "github.com/alcionai/corso/src/internal/common/idname"

var _ idname.Provider = &in{}

func NewProvider(id, name string) *in {
	return &in{
		id:   id,
		name: name,
	}
}

type in struct {
	id   string
	name string
}

func (i in) ID() string   { return i.id }
func (i in) Name() string { return i.name }
