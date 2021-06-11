// source: error.proto
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

goog.exportSymbol('proto.protoim.Error', null, global);
/**
 * @enum {number}
 */
proto.protoim.Error = {
  E_UNKNOWN: 0,
  E_SERVER: 1,
  E_WRITE_TIMEOUT: 2,
  E_READ_TIMEOUT: 3,
  E_ROUTE_NOT_EXIST: 4,
  E_DECODE: 5,
  E_ENCODE: 6,
  E_NOT_LOGIN: 7,
  E_PARAMETER: 8,
  E_DB_NOT_DATA: 9,
  E_DB_SELECT: 10,
  E_DB_INSERT: 11,
  E_DB_DELETE: 12,
  E_DB_UPATE: 13,
  E_DB_OPERATION: 14
};

goog.object.extend(exports, proto.protoim);
