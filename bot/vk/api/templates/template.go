package templates

import (
	"github.com/Galagoshin/GoUtils/json"
)

type Template interface {
	Init()
	getData() map[string]any
	setData(map[string]any)
	GetJson() json.Json
}
