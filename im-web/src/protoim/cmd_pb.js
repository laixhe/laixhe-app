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
  C_ERROR: 0,
  PING: 1,
  PONG: 2,
  USER_LOGIN_REQUEST: 1000,
  USER_LOGIN_RESPONSE: 1001,
  UPDATE_USER_FRIEND_INFO: 1002,
  GET_USER_INFO_REQUEST: 1003,
  GET_USER_INFO_RESPONSE: 1004
};

goog.object.extend(exports, proto.protoim);
