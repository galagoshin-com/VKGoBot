package widgets

import "github.com/Galagoshin/GoUtils/json"

type WidgetType string

const (
	Table = WidgetType("table")
)

type Widget interface {
	Init()
	GetType() WidgetType
	getData() map[string]interface{}
	setData(map[string]interface{})
	GetJson() json.Json
}
