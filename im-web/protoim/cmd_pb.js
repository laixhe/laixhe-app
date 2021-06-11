// source: cmd.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.protoim.CMD', null, global);
/**
 * @enum {number}
 */
proto.protoim.CMD = {
  C_UNKNOWN: 0,
  C_ERROR: 1,
  C_PING: 2,
  C_PONG: 3,
  C_USER_LOGIN_REQUEST: 1000,
  C_USER_LOGIN_RESPONSE: 1001,
  C_GET_USER_INFO_REQUEST: 1002,
  C_GET_USER_INFO_RESPONSE: 1003,
  C_UPDATE_FRIEND_INFO: 1004,
  C_MESSAGE_REQUEST: 1005,
  C_MESSAGE_RESPONSE: 1006,
  C_MESSAGE_FEEDBACK: 1007
};

goog.object.extend(exports, proto.protoim);
