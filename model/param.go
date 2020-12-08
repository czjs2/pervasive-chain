package model

type Param map[string]interface{}

func (u Param) MarshalJSON() ([]byte, error) {

	return nil, nil
}
