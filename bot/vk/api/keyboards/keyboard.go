package keyboards

import (
	"github.com/Galagoshin/GoUtils/json"
)

type Keyboard interface {
	Init()
	getData() map[string]any
	setData(map[string]any)
	GetJson() json.Json
}
