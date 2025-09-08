package errs

import "github.com/rotisserie/eris"

func ToJson(err error) map[string]interface{} {
	return eris.ToJSON(err, true)
}

func ToString(err error) string {
	return eris.ToString(err, true)
}

type ErrNextMap struct {
	Message string
}

func NewErrNextMap(message string) *ErrNextMap {
	return &ErrNextMap{Message: message}
}

func (e ErrNextMap) Error() string {
	return e.Message
}
