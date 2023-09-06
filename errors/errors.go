package customErrors

type DataBaseError struct {
	Message string
}

func (e *DataBaseError) Error() string {
	return e.Message
}
