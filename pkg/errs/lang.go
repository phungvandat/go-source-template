package errs

// exp: map[en]map[code]{message: abc}
type langMap map[ErrLang]langErrMap

// exp: map[code]{message: "abc"}
type langErrMap map[int]langError

var mapOfLang = make(langMap)

func init() {
	mapOfEnLangErr := make(map[int]langError)
	mapOfLang[EN] = mapOfEnLangErr
}
