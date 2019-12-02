package ws

type AppOS int

//登录的平台 android/ios/web/pc
const (
	ANDROID AppOS = 1 // android
	IOS     AppOS = 2 // ios
	PC      AppOS = 3 // pc
	WEB     AppOS = 4 // web
)
