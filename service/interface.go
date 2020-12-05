package service

type IFormValidateInterface interface {
	Valid() (bool, error)
}
