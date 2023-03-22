package keyboards

import (
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/json"
	"github.com/Galagoshin/VKGoBot/bot/vk/api/sign"
)

type Payload map[string]any

func (payload Payload) Sign() json.Json {
	return sign.SignPayload2Json(payload)
}

func (payload Payload) Verify() bool {
	signed := payload["sign"].(map[string]any)
	valid := sign.SignPayload2Map(payload)["sign"].(map[string]any)
	logger.Debug(2, false, fmt.Sprintf("Required sign: %v", valid["sign"]))
	for param, value := range valid {
		if signed[param].(string) != value.(string) {
			return false
		}
	}
	return true
}
