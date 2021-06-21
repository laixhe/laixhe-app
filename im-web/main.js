var app_os       = require('./protoim/app_os_pb')
var error        = require('./protoim/error_pb')
var cmd          = require('./protoim/cmd_pb')
var chat_type    = require('./protoim/chat_type_pb')
var message_type = require('./protoim/message_type_pb')
var user_type    = require('./protoim/user_type_pb')

var error_base   = require('./protoim/error_base_pb')
var message_base = require('./protoim/message_base_pb')

var user_info                  = require('./protoim/user_info_pb')
var user_login_request         = require('./protoim/user_login_request_pb')
var user_login_response        = require('./protoim/user_login_response_pb')
var get_user_request           = require('./protoim/get_user_request_pb')
var get_user_response          = require('./protoim/get_user_response_pb')
var get_friends_request        = require('./protoim/get_friends_request_pb')
var get_friends_response       = require('./protoim/get_friends_response_pb')
var update_friend_type         = require('./protoim/update_friend_type_pb')
var update_friend              = require('./protoim/update_friend_pb')
var update_friends             = require('./protoim/update_friends_pb')
var get_conversations_request  = require('./protoim/get_conversations_request_pb')
var get_conversations_response = require('./protoim/get_conversations_response_pb')
var message_entity_uri         = require('./protoim/message_entity_uri_pb')
var message_media              = require('./protoim/message_media_pb')
var message_request            = require('./protoim/message_request_pb')
var message_response           = require('./protoim/message_response_pb')
var message_feedback           = require('./protoim/message_feedback_pb')

module.exports = {
    AppOs                : app_os,
    Error                : error,
    Cmd                  : cmd,
    ChatType             : chat_type,
    MessageType          : message_type,
    UserType             : user_type,
    
    ErrorBase            : error_base,
    MessageBase          : message_base,

    UserInfo             : user_info,
    UserLoginRequest     : user_login_request,
    UserLoginResponse    : user_login_response,
    GetUserRequest       : get_user_request,
    GetUserResponse      : get_user_response,
    GetFriendsRequest    : get_friends_request,
    GetFriendsResponse   : get_friends_response,
    UpdateFriendType     : update_friend_type,
    UpdateFriend         : update_friend,
    UpdateFriends        : update_friends,
    GetConversationsRequest  : get_conversations_request,
    GetConversationsResponse : get_conversations_response,
    MessageEntityUri     : message_entity_uri,
    MessageMedia         : message_media,
    MessageRequest       : message_request,
    MessageResponse      : message_response,
    MessageFeedback      : message_feedback,
}