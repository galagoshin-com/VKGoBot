package widgets

import (
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/GoUtils/requests"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/attachments"
	"strings"
)

type WidgetComactList struct {
	data         map[string]any
	Title        string
	TitleUrl     requests.URL
	FooterText   string
	FooterUrl    requests.URL
	TitleCounter int
}

type ComactListRow struct {
	Index       int
	Title       string
	TitleUrl    requests.URL
	ButtonText  string
	ButtonUrl   requests.URL
	Address     string
	Time        string
	Description string
	Icon        attachments.Image
}

func (widget *WidgetComactList) Init() {
	data := make(map[string]any)
	data["title"] = widget.Title
	data["title_url"] = widget.TitleUrl
	data["title_conter"] = widget.TitleCounter
	data["more"] = widget.FooterText
	data["more_url"] = widget.FooterUrl
	data["rows"] = make([]any, 6)
	widget.data = data
}

func (widjet *WidgetComactList) AddRow(row ListRow) *WidgetComactList {
	data := widjet.data
	data["rows"].([]any)[row.Index] = make(map[string]any)
	data["rows"].([]any)[row.Index].(map[string]any)["title"] = row.Title
	data["rows"].([]any)[row.Index].(map[string]any)["title_url"] = row.TitleUrl
	data["rows"].([]any)[row.Index].(map[string]any)["button"] = row.ButtonText
	data["rows"].([]any)[row.Index].(map[string]any)["button_url"] = row.ButtonUrl
	if row.Icon.OwnerId != 0 {
		data["rows"].([]any)[row.Index].(map[string]any)["icon_id"] = row.Icon.BuildString()
	}
	data["rows"].([]any)[row.Index].(map[string]any)["address"] = row.Address
	data["rows"].([]any)[row.Index].(map[string]any)["time"] = row.Time
	data["rows"].([]any)[row.Index].(map[string]any)["text"] = row.Description
	widjet.setData(data)
	return widjet
}

func (widjet *WidgetComactList) GetType() WidgetType {
	return CompactList
}

func (widjet *WidgetComactList) getData() map[string]any {
	return widjet.data
}

func (widjet *WidgetComactList) setData(data map[string]any) {
	widjet.data = data
}

func (widjet *WidgetComactList) GetJson() json.Json {
	jsn, err := json.Encode(widjet.data)
	if err != nil {
		logger.Error(err)
	}
	return json.Json(strings.Replace(strings.Replace(string(jsn), ", null", "", -1), ",null", "", -1))
}
