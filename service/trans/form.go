package trans

type TransGroupFrom struct {
	Height    int    `json:"height"`
	FromShard string `json:"fromShard"`
	ToShard   string `json:"toShard"`
}

func (t *TransGroupFrom) Valid() (bool, error) {
	if t.Height == 0 || t.FromShard == "" || t.ToShard == "" {
		return false, nil
	}
	return true, nil
}

type TransFrom struct {
	Hash string `json:"hash"`
}

func (t *TransFrom) Valid() (bool, error) {
	return t.Hash != "", nil
}
