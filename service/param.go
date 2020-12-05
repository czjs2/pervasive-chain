package service


type Param map[string]interface{}


func (u Param) MarshalJSON() ([]byte, error) {

	panic(u)
}
