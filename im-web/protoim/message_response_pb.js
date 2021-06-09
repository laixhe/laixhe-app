// source: message_response.proto
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

var chat_type_pb = require('./chat_type_pb.js');
goog.object.extend(proto, chat_type_pb);
var message_type_pb = require('./message_type_pb.js');
goog.object.extend(proto, message_type_pb);
goog.exportSymbol('proto.protoim.MessageResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protoim.MessageResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protoim.MessageResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protoim.MessageResponse.displayName = 'proto.protoim.MessageResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protoim.MessageResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.protoim.MessageResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protoim.MessageResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protoim.MessageResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    msgId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    localMsgId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    pts: jspb.Message.getFieldWithDefault(msg, 3, 0),
    fromId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    chatTypeId: jspb.Message.getFieldWithDefault(msg, 5, 0),
    toId: jspb.Message.getFieldWithDefault(msg, 6, ""),
    messageTypeId: jspb.Message.getFieldWithDefault(msg, 7, 0),
    content: jspb.Message.getFieldWithDefault(msg, 8, ""),
    dataTime: jspb.Message.getFieldWithDefault(msg, 9, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protoim.MessageResponse}
 */
proto.protoim.MessageResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protoim.MessageResponse;
  return proto.protoim.MessageResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protoim.MessageResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protoim.MessageResponse}
 */
proto.protoim.MessageResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setMsgId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLocalMsgId(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPts(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setFromId(value);
      break;
    case 5:
      var value = /** @type {!proto.protoim.ChatType} */ (reader.readEnum());
      msg.setChatTypeId(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setToId(value);
      break;
    case 7:
      var value = /** @type {!proto.protoim.MessageType} */ (reader.readEnum());
      msg.setMessageTypeId(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setContent(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setDataTime(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protoim.MessageResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protoim.MessageResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protoim.MessageResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protoim.MessageResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMsgId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getLocalMsgId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getPts();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getFromId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getChatTypeId();
  if (f !== 0.0) {
    writer.writeEnum(
      5,
      f
    );
  }
  f = message.getToId();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getMessageTypeId();
  if (f !== 0.0) {
    writer.writeEnum(
      7,
      f
    );
  }
  f = message.getContent();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getDataTime();
  if (f !== 0) {
    writer.writeInt64(
      9,
      f
    );
  }
};


/**
 * optional string msg_id = 1;
 * @return {string}
 */
proto.protoim.MessageResponse.prototype.getMsgId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protoim.MessageResponse} returns this
 */
proto.protoim.MessageResponse.prototype.setMsgId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string local_msg_id = 2;
 * @return {string}
 */
proto.protoim.MessageResponse.prototype.getLocalMsgId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.protoim.MessageResponse} returns this
 */
proto.protoim.MessageResponse.prototype.setLocalMsgId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int32 pts = 3;
 * @return {number}
 */
proto.protoim.MessageResponse.prototype.getPts = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.protoim.MessageResponse} returns this
 */
proto.protoim.MessageResponse.prototype.setPts = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional string from_id = 4;
 * @return {string}
 */
proto.protoim.MessageResponse.prototype.getFromId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.protoim.MessageResponse} returns this
 */
proto.protoim.MessageResponse.prototype.setFromId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional ChatType chat_type_id = 5;
 * @return {!proto.protoim.ChatType}
 */
proto.protoim.MessageResponse.prototype.getChatTypeId = function() {
  return /** @type {!proto.protoim.ChatType} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {!proto.protoim.ChatType} value
 * @return {!proto.protoim.MessageResponse} returns this
 */
proto.protoim.MessageResponse.prototype.setChatTypeId = function(value) {
  return jspb.Message.setProto3EnumField(this, 5, value);
};


/**
 * optional string to_id = 6;
 * @return {string}
 */
proto.protoim.MessageResponse.prototype.getToId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.protoim.MessageResponse} returns this
 */
proto.protoim.MessageResponse.prototype.setToId = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional MessageType message_type_id = 7;
 * @return {!proto.protoim.MessageType}
 */
proto.protoim.MessageResponse.prototype.getMessageTypeId = function() {
  return /** @type {!proto.protoim.MessageType} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {!proto.protoim.MessageType} value
 * @return {!proto.protoim.MessageResponse} returns this
 */
proto.protoim.MessageResponse.prototype.setMessageTypeId = function(value) {
  return jspb.Message.setProto3EnumField(this, 7, value);
};


/**
 * optional string content = 8;
 * @return {string}
 */
proto.protoim.MessageResponse.prototype.getContent = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.protoim.MessageResponse} returns this
 */
proto.protoim.MessageResponse.prototype.setContent = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional int64 data_time = 9;
 * @return {number}
 */
proto.protoim.MessageResponse.prototype.getDataTime = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/**
 * @param {number} value
 * @return {!proto.protoim.MessageResponse} returns this
 */
proto.protoim.MessageResponse.prototype.setDataTime = function(value) {
  return jspb.Message.setProto3IntField(this, 9, value);
};


goog.object.extend(exports, proto.protoim);
