package helpers

import (
	"runtime"
)

const (
	defaultDepth = 1
	defaultIndex = 0
)

// GetFunctionName Возвращает имяПакета.ИмяФункции.
func GetFunctionName(depthList ...int) string { //nolint:unused // helper func
	var depth int

	if depthList == nil {
		depth = defaultDepth
	} else {
		depth = depthList[defaultIndex]
	}

	function, _, _, ok := runtime.Caller(depth)
	if !ok {
		return "Не удалось получить имя функции"
	}

	return runtime.FuncForPC(function).Name()
}
