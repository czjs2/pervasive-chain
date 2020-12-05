package form

type IFormValidateInterface interface {
	Valid() (bool, error)
}
