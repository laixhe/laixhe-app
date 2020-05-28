package servers

type AppOS int
type AppFormat int

// 登录的平台 android/ios/web/pc
const (
	ANDROID AppOS = 1 // android
	IOS     AppOS = 2 // ios
	PC      AppOS = 3 // pc
	WEB     AppOS = 4 // web
)

// 序列化格式 json/protobuf
const (
	JSON     AppFormat = 1 // json
	PROTOBUF AppFormat = 2 // protobuf
)
