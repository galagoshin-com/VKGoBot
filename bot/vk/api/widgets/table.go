package widgets

import (
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/requests"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"strings"
)
import "github.com/Galagoshin/GoUtils/json"

type WidgetTable struct {
	data                map[string]any
	Title               string
	TitleUrl            string
	ColumnsDescriptions []string
}

type TableRow struct {
	Index   int
	Columns []TableColumn
}

type TableColumn struct {
	Index int
	Text  string
	Url   requests.URL
	Icon  attachments.Image
}

func (widget *WidgetTable) Init() {
	data := make(map[string]any)
	data["title"] = widget.Title
	data["title_url"] = widget.TitleUrl
	data["head"] = []any{}
	for i, desc := range widget.ColumnsDescriptions {
		data["head"] = append(data["head"].([]any), make(map[string]any))
		data["head"].([]any)[i].(map[string]any)["text"] = desc
		data["body"] = []any{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	}
	widget.data = data
}

func (widjet *WidgetTable) AddRow(row TableRow) *WidgetTable {
	data := widjet.data
	if data["body"].([]any)[row.Index] == nil {
		data["body"].([]any)[row.Index] = []any{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	}
	for _, column := range row.Columns {
		data["body"].([]any)[row.Index].([]any)[column.Index] = make(map[string]any)
		data["body"].([]any)[row.Index].([]any)[column.Index].(map[string]any)["text"] = column.Text
		data["body"].([]any)[row.Index].([]any)[column.Index].(map[string]any)["url"] = column.Url
		if column.Icon.OwnerId != 0 {
			data["body"].([]any)[row.Index].([]any)[column.Index].(map[string]any)["icon_id"] = fmt.Sprintf("id%d", column.Icon.OwnerId)
		}
	}
	widjet.setData(data)
	return widjet
}

func (widjet *WidgetTable) GetType() WidgetType {
	return Table
}

func (widjet *WidgetTable) getData() map[string]any {
	return widjet.data
}

func (widjet *WidgetTable) setData(data map[string]any) {
	widjet.data = data
}

func (widjet *WidgetTable) GetJson() json.Json {
	jsn, err := json.Encode(widjet.data)
	if err != nil {
		logger.Error(err)
	}
	return json.Json(strings.Replace(strings.Replace(string(jsn), ", null", "", -1), ",null", "", -1))
}
