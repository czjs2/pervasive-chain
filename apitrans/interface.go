package apinode

type ITransService interface {
	QueryTransByGroupId() (interface{}, error)
}
