/**
 * @fileoverview gRPC-Web generated client stub for playground
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */


const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.playground = require('./playground_service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.playground.PlaygroundClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.playground.PlaygroundPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.playground.EmptyRequest,
 *   !proto.playground.PingResponse>}
 */
const methodDescriptor_Playground_Ping = new grpc.web.MethodDescriptor(
  '/playground.Playground/Ping',
  grpc.web.MethodType.UNARY,
  proto.playground.EmptyRequest,
  proto.playground.PingResponse,
  /**
   * @param {!proto.playground.EmptyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playground.PingResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.playground.EmptyRequest,
 *   !proto.playground.PingResponse>}
 */
const methodInfo_Playground_Ping = new grpc.web.AbstractClientBase.MethodInfo(
  proto.playground.PingResponse,
  /**
   * @param {!proto.playground.EmptyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playground.PingResponse.deserializeBinary
);


/**
 * @param {!proto.playground.EmptyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.playground.PingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.playground.PingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.playground.PlaygroundClient.prototype.ping =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/playground.Playground/Ping',
      request,
      metadata || {},
      methodDescriptor_Playground_Ping,
      callback);
};


/**
 * @param {!proto.playground.EmptyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.playground.PingResponse>}
 *     A native promise that resolves to the response
 */
proto.playground.PlaygroundPromiseClient.prototype.ping =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/playground.Playground/Ping',
      request,
      metadata || {},
      methodDescriptor_Playground_Ping);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.playground.CodeRequest,
 *   !proto.playground.RunResponse>}
 */
const methodDescriptor_Playground_RunCode = new grpc.web.MethodDescriptor(
  '/playground.Playground/RunCode',
  grpc.web.MethodType.UNARY,
  proto.playground.CodeRequest,
  proto.playground.RunResponse,
  /**
   * @param {!proto.playground.CodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playground.RunResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.playground.CodeRequest,
 *   !proto.playground.RunResponse>}
 */
const methodInfo_Playground_RunCode = new grpc.web.AbstractClientBase.MethodInfo(
  proto.playground.RunResponse,
  /**
   * @param {!proto.playground.CodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playground.RunResponse.deserializeBinary
);


/**
 * @param {!proto.playground.CodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.playground.RunResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.playground.RunResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.playground.PlaygroundClient.prototype.runCode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/playground.Playground/RunCode',
      request,
      metadata || {},
      methodDescriptor_Playground_RunCode,
      callback);
};


/**
 * @param {!proto.playground.CodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.playground.RunResponse>}
 *     A native promise that resolves to the response
 */
proto.playground.PlaygroundPromiseClient.prototype.runCode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/playground.Playground/RunCode',
      request,
      metadata || {},
      methodDescriptor_Playground_RunCode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.playground.CodeRequest,
 *   !proto.playground.FormatCodeResponse>}
 */
const methodDescriptor_Playground_FormatCode = new grpc.web.MethodDescriptor(
  '/playground.Playground/FormatCode',
  grpc.web.MethodType.UNARY,
  proto.playground.CodeRequest,
  proto.playground.FormatCodeResponse,
  /**
   * @param {!proto.playground.CodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playground.FormatCodeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.playground.CodeRequest,
 *   !proto.playground.FormatCodeResponse>}
 */
const methodInfo_Playground_FormatCode = new grpc.web.AbstractClientBase.MethodInfo(
  proto.playground.FormatCodeResponse,
  /**
   * @param {!proto.playground.CodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playground.FormatCodeResponse.deserializeBinary
);


/**
 * @param {!proto.playground.CodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.playground.FormatCodeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.playground.FormatCodeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.playground.PlaygroundClient.prototype.formatCode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/playground.Playground/FormatCode',
      request,
      metadata || {},
      methodDescriptor_Playground_FormatCode,
      callback);
};


/**
 * @param {!proto.playground.CodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.playground.FormatCodeResponse>}
 *     A native promise that resolves to the response
 */
proto.playground.PlaygroundPromiseClient.prototype.formatCode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/playground.Playground/FormatCode',
      request,
      metadata || {},
      methodDescriptor_Playground_FormatCode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.playground.CodeRequest,
 *   !proto.playground.ShareCodeResponse>}
 */
const methodDescriptor_Playground_ShareCode = new grpc.web.MethodDescriptor(
  '/playground.Playground/ShareCode',
  grpc.web.MethodType.UNARY,
  proto.playground.CodeRequest,
  proto.playground.ShareCodeResponse,
  /**
   * @param {!proto.playground.CodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playground.ShareCodeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.playground.CodeRequest,
 *   !proto.playground.ShareCodeResponse>}
 */
const methodInfo_Playground_ShareCode = new grpc.web.AbstractClientBase.MethodInfo(
  proto.playground.ShareCodeResponse,
  /**
   * @param {!proto.playground.CodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playground.ShareCodeResponse.deserializeBinary
);


/**
 * @param {!proto.playground.CodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.playground.ShareCodeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.playground.ShareCodeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.playground.PlaygroundClient.prototype.shareCode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/playground.Playground/ShareCode',
      request,
      metadata || {},
      methodDescriptor_Playground_ShareCode,
      callback);
};


/**
 * @param {!proto.playground.CodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.playground.ShareCodeResponse>}
 *     A native promise that resolves to the response
 */
proto.playground.PlaygroundPromiseClient.prototype.shareCode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/playground.Playground/ShareCode',
      request,
      metadata || {},
      methodDescriptor_Playground_ShareCode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.playground.CommonRequest,
 *   !proto.playground.GetCodeByShareResponse>}
 */
const methodDescriptor_Playground_GetCodeByShare = new grpc.web.MethodDescriptor(
  '/playground.Playground/GetCodeByShare',
  grpc.web.MethodType.UNARY,
  proto.playground.CommonRequest,
  proto.playground.GetCodeByShareResponse,
  /**
   * @param {!proto.playground.CommonRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playground.GetCodeByShareResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.playground.CommonRequest,
 *   !proto.playground.GetCodeByShareResponse>}
 */
const methodInfo_Playground_GetCodeByShare = new grpc.web.AbstractClientBase.MethodInfo(
  proto.playground.GetCodeByShareResponse,
  /**
   * @param {!proto.playground.CommonRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.playground.GetCodeByShareResponse.deserializeBinary
);


/**
 * @param {!proto.playground.CommonRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.playground.GetCodeByShareResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.playground.GetCodeByShareResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.playground.PlaygroundClient.prototype.getCodeByShare =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/playground.Playground/GetCodeByShare',
      request,
      metadata || {},
      methodDescriptor_Playground_GetCodeByShare,
      callback);
};


/**
 * @param {!proto.playground.CommonRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.playground.GetCodeByShareResponse>}
 *     A native promise that resolves to the response
 */
proto.playground.PlaygroundPromiseClient.prototype.getCodeByShare =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/playground.Playground/GetCodeByShare',
      request,
      metadata || {},
      methodDescriptor_Playground_GetCodeByShare);
};


module.exports = proto.playground;

