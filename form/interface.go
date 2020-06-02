package form

type IFormInterface interface {
	Valid() (bool, error)
}
