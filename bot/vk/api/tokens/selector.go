package tokens

var tokenIndex = 0
var userTokenIndex = 0
var widgetTokenIndex = 0

func GetWidgetToken() string {
	token := WidgetTokensList[widgetTokenIndex]
	if widgetTokenIndex+1 == len(WidgetTokensList) {
		widgetTokenIndex = 0
	} else {
		widgetTokenIndex++
	}
	return token
}

func GetToken() string {
	token := GroupTokensList[tokenIndex]
	if tokenIndex+1 == len(GroupTokensList) {
		tokenIndex = 0
	} else {
		tokenIndex++
	}
	return token
}

func GetUserToken() string {
	token := UserTokensList[userTokenIndex]
	if userTokenIndex+1 == len(UserTokensList) {
		userTokenIndex = 0
	} else {
		userTokenIndex++
	}
	return token
}
