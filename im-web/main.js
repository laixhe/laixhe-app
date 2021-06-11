var app_os = require('./protoim/app_os_pb')
var error = require('./protoim/error_pb')
var cmd = require('./protoim/cmd_pb')
var chat_type = require('./protoim/chat_type_pb')
var message_type = require('./protoim/message_type_pb')

var error_info = require('./protoim/error_info_pb')
var message_base = require('./protoim/message_base_pb')

var user_login_request= require('./protoim/user_login_request_pb')
var user_login_response = require('./protoim/user_login_response_pb')
var update_friend_info = require('./protoim/update_friend_info_pb')
var user_info = require('./protoim/user_info_pb')
var get_user_info_request = require('./protoim/get_user_info_request_pb')
var get_user_info_response = require('./protoim/get_user_info_response_pb')

var message_request = require('./protoim/message_request_pb')
var message_response = require('./protoim/message_response_pb')
var message_feedback = require('./protoim/message_feedback_pb')

module.exports = {
    AppOs       : app_os,
    Error       : error,
    Cmd         : cmd,
    ChatType    : chat_type,
    MessageType : message_type,

    ErrorInfo   : error_info,
    MessageBase : message_base,

    UserInfo             : user_info,
    UserLoginRequest     : user_login_request,
    UserLoginResponse    : user_login_response,
    UpdateFriendInfo     : update_friend_info,
    GetUserInfoRequest   : get_user_info_request,
    GetUserInfoResponse  : get_user_info_response,

    MessageRequest   : message_request,
    MessageResponse  : message_response,
    MessageFeedback  : message_feedback,
}