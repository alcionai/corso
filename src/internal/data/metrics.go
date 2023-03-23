package data

type CollectionStats struct {
	Folders            int
	Objects, Successes int
	Bytes              int64
	Details            string
}

func (cs CollectionStats) IsZero() bool {
	return cs.Folders+cs.Objects+cs.Successes+int(cs.Bytes) == 0
}

func (cs CollectionStats) String() string {
	return cs.Details
}
