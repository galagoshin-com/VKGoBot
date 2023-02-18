package sign

import (
	"fmt"
	"github.com/Galagoshin/GoLogger/logger"
	"github.com/Galagoshin/GoUtils/crypto"
	"github.com/Galagoshin/GoUtils/json"
)

func SignPayload2Json(payloads map[string]any) json.Json {
	payloads["sign"] = make(map[string]any)
	for param, value := range payloads {
		if param != "sign" {
			payloads["sign"].(map[string]any)[param] = crypto.Sha256([]byte(fmt.Sprintf("%s%v%s", param, value, SECRET)))
		}
	}
	jsn, err := json.Encode(payloads)
	if err != nil {
		logger.Error(err)
	}
	return jsn
}

func SignPayload2Map(payloads map[string]any) map[string]any {
	payloads["sign"] = make(map[string]any)
	for param, value := range payloads {
		if param != "sign" {
			payloads["sign"].(map[string]any)[param] = crypto.Sha256([]byte(fmt.Sprintf("%s%v%s", param, value, SECRET)))
		}
	}
	return payloads
}
