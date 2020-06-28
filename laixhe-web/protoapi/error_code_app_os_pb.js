// source: error_code_app_os.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.protoapi.AppOS', null, global);
goog.exportSymbol('proto.protoapi.ErrorCode', null, global);
/**
 * @enum {number}
 */
proto.protoapi.ErrorCode = {
  E_SUCCESS: 0,
  E_UNKNOWN: 1,
  E_SERVER: 2,
  E_WEBSOCKET: 3,
  E_WEBSOCKET_WRITE: 4,
  E_WEBSOCKET_READ: 5,
  E_WEBSOCKET_TIMEOUT: 6,
  E_DECODE: 7,
  E_ENCODE: 8,
  E_AUTH: 9,
  E_NOT_LOGGED_IN: 10,
  E_OPERATION_FAILURE: 11,
  E_ROUTING_NOT_EXIST: 12,
  E_DB: 13,
  E_NOT_SUPPORT_API: 14,
  E_INVALID_JSON: 15,
  E_ENCODE_JSON: 16,
  E_PARAMETER_EMPTY: 17,
  E_PARAMETER: 18,
  E_RESOURCE: 19,
  E_RESOURCE_REPEAT: 20
};

/**
 * @enum {number}
 */
proto.protoapi.AppOS = {
  O_UNKNOWN: 0,
  O_IOS: 1,
  O_ANDROID: 2,
  O_WEB: 3,
  O_WINDOW: 4,
  O_MAC: 5,
  O_LINUX: 6
};

goog.object.extend(exports, proto.protoapi);
