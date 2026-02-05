// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_chatlog_pb = require('../proto/chatlog_pb.js');

function serialize_chatlog_BackupRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.BackupRequest)) {
    throw new Error('Expected argument of type chatlog.BackupRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_BackupRequest(buffer_arg) {
  return proto_chatlog_pb.BackupRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_BackupResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.BackupResponse)) {
    throw new Error('Expected argument of type chatlog.BackupResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_BackupResponse(buffer_arg) {
  return proto_chatlog_pb.BackupResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_CommandDecryptRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.CommandDecryptRequest)) {
    throw new Error('Expected argument of type chatlog.CommandDecryptRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_CommandDecryptRequest(buffer_arg) {
  return proto_chatlog_pb.CommandDecryptRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_CommandDecryptResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.CommandDecryptResponse)) {
    throw new Error('Expected argument of type chatlog.CommandDecryptResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_CommandDecryptResponse(buffer_arg) {
  return proto_chatlog_pb.CommandDecryptResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_CommandHTTPServerRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.CommandHTTPServerRequest)) {
    throw new Error('Expected argument of type chatlog.CommandHTTPServerRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_CommandHTTPServerRequest(buffer_arg) {
  return proto_chatlog_pb.CommandHTTPServerRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_CommandHTTPServerResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.CommandHTTPServerResponse)) {
    throw new Error('Expected argument of type chatlog.CommandHTTPServerResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_CommandHTTPServerResponse(buffer_arg) {
  return proto_chatlog_pb.CommandHTTPServerResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_CommandKeyRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.CommandKeyRequest)) {
    throw new Error('Expected argument of type chatlog.CommandKeyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_CommandKeyRequest(buffer_arg) {
  return proto_chatlog_pb.CommandKeyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_CommandKeyResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.CommandKeyResponse)) {
    throw new Error('Expected argument of type chatlog.CommandKeyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_CommandKeyResponse(buffer_arg) {
  return proto_chatlog_pb.CommandKeyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_DecryptDBFilesRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.DecryptDBFilesRequest)) {
    throw new Error('Expected argument of type chatlog.DecryptDBFilesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_DecryptDBFilesRequest(buffer_arg) {
  return proto_chatlog_pb.DecryptDBFilesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_DecryptDBFilesResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.DecryptDBFilesResponse)) {
    throw new Error('Expected argument of type chatlog.DecryptDBFilesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_DecryptDBFilesResponse(buffer_arg) {
  return proto_chatlog_pb.DecryptDBFilesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_DecryptRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.DecryptRequest)) {
    throw new Error('Expected argument of type chatlog.DecryptRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_DecryptRequest(buffer_arg) {
  return proto_chatlog_pb.DecryptRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_DecryptResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.DecryptResponse)) {
    throw new Error('Expected argument of type chatlog.DecryptResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_DecryptResponse(buffer_arg) {
  return proto_chatlog_pb.DecryptResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_GetDataKeyRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.GetDataKeyRequest)) {
    throw new Error('Expected argument of type chatlog.GetDataKeyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_GetDataKeyRequest(buffer_arg) {
  return proto_chatlog_pb.GetDataKeyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_GetDataKeyResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.GetDataKeyResponse)) {
    throw new Error('Expected argument of type chatlog.GetDataKeyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_GetDataKeyResponse(buffer_arg) {
  return proto_chatlog_pb.GetDataKeyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_GetKeyRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.GetKeyRequest)) {
    throw new Error('Expected argument of type chatlog.GetKeyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_GetKeyRequest(buffer_arg) {
  return proto_chatlog_pb.GetKeyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_GetKeyResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.GetKeyResponse)) {
    throw new Error('Expected argument of type chatlog.GetKeyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_GetKeyResponse(buffer_arg) {
  return proto_chatlog_pb.GetKeyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_GetWeChatInstancesRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.GetWeChatInstancesRequest)) {
    throw new Error('Expected argument of type chatlog.GetWeChatInstancesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_GetWeChatInstancesRequest(buffer_arg) {
  return proto_chatlog_pb.GetWeChatInstancesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_GetWeChatInstancesResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.GetWeChatInstancesResponse)) {
    throw new Error('Expected argument of type chatlog.GetWeChatInstancesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_GetWeChatInstancesResponse(buffer_arg) {
  return proto_chatlog_pb.GetWeChatInstancesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_RefreshSessionRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.RefreshSessionRequest)) {
    throw new Error('Expected argument of type chatlog.RefreshSessionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_RefreshSessionRequest(buffer_arg) {
  return proto_chatlog_pb.RefreshSessionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_RefreshSessionResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.RefreshSessionResponse)) {
    throw new Error('Expected argument of type chatlog.RefreshSessionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_RefreshSessionResponse(buffer_arg) {
  return proto_chatlog_pb.RefreshSessionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_RunRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.RunRequest)) {
    throw new Error('Expected argument of type chatlog.RunRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_RunRequest(buffer_arg) {
  return proto_chatlog_pb.RunRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_RunResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.RunResponse)) {
    throw new Error('Expected argument of type chatlog.RunResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_RunResponse(buffer_arg) {
  return proto_chatlog_pb.RunResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_SetHTTPAddrRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.SetHTTPAddrRequest)) {
    throw new Error('Expected argument of type chatlog.SetHTTPAddrRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_SetHTTPAddrRequest(buffer_arg) {
  return proto_chatlog_pb.SetHTTPAddrRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_SetHTTPAddrResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.SetHTTPAddrResponse)) {
    throw new Error('Expected argument of type chatlog.SetHTTPAddrResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_SetHTTPAddrResponse(buffer_arg) {
  return proto_chatlog_pb.SetHTTPAddrResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_SetLogLevelRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.SetLogLevelRequest)) {
    throw new Error('Expected argument of type chatlog.SetLogLevelRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_SetLogLevelRequest(buffer_arg) {
  return proto_chatlog_pb.SetLogLevelRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_SetLogLevelResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.SetLogLevelResponse)) {
    throw new Error('Expected argument of type chatlog.SetLogLevelResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_SetLogLevelResponse(buffer_arg) {
  return proto_chatlog_pb.SetLogLevelResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_StartAutoDecryptRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.StartAutoDecryptRequest)) {
    throw new Error('Expected argument of type chatlog.StartAutoDecryptRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_StartAutoDecryptRequest(buffer_arg) {
  return proto_chatlog_pb.StartAutoDecryptRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_StartAutoDecryptResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.StartAutoDecryptResponse)) {
    throw new Error('Expected argument of type chatlog.StartAutoDecryptResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_StartAutoDecryptResponse(buffer_arg) {
  return proto_chatlog_pb.StartAutoDecryptResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_StartServiceRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.StartServiceRequest)) {
    throw new Error('Expected argument of type chatlog.StartServiceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_StartServiceRequest(buffer_arg) {
  return proto_chatlog_pb.StartServiceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_StartServiceResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.StartServiceResponse)) {
    throw new Error('Expected argument of type chatlog.StartServiceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_StartServiceResponse(buffer_arg) {
  return proto_chatlog_pb.StartServiceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_StopAutoDecryptRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.StopAutoDecryptRequest)) {
    throw new Error('Expected argument of type chatlog.StopAutoDecryptRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_StopAutoDecryptRequest(buffer_arg) {
  return proto_chatlog_pb.StopAutoDecryptRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_StopAutoDecryptResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.StopAutoDecryptResponse)) {
    throw new Error('Expected argument of type chatlog.StopAutoDecryptResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_StopAutoDecryptResponse(buffer_arg) {
  return proto_chatlog_pb.StopAutoDecryptResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_StopServiceRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.StopServiceRequest)) {
    throw new Error('Expected argument of type chatlog.StopServiceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_StopServiceRequest(buffer_arg) {
  return proto_chatlog_pb.StopServiceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_StopServiceResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.StopServiceResponse)) {
    throw new Error('Expected argument of type chatlog.StopServiceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_StopServiceResponse(buffer_arg) {
  return proto_chatlog_pb.StopServiceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_SwitchRequest(arg) {
  if (!(arg instanceof proto_chatlog_pb.SwitchRequest)) {
    throw new Error('Expected argument of type chatlog.SwitchRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_SwitchRequest(buffer_arg) {
  return proto_chatlog_pb.SwitchRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_chatlog_SwitchResponse(arg) {
  if (!(arg instanceof proto_chatlog_pb.SwitchResponse)) {
    throw new Error('Expected argument of type chatlog.SwitchResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_chatlog_SwitchResponse(buffer_arg) {
  return proto_chatlog_pb.SwitchResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// ManagerService provides gRPC interface for WeChat operations
var ManagerServiceService = exports.ManagerServiceService = {
  setLogLevel: {
    path: '/chatlog.ManagerService/SetLogLevel',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.SetLogLevelRequest,
    responseType: proto_chatlog_pb.SetLogLevelResponse,
    requestSerialize: serialize_chatlog_SetLogLevelRequest,
    requestDeserialize: deserialize_chatlog_SetLogLevelRequest,
    responseSerialize: serialize_chatlog_SetLogLevelResponse,
    responseDeserialize: deserialize_chatlog_SetLogLevelResponse,
  },
  run: {
    path: '/chatlog.ManagerService/Run',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.RunRequest,
    responseType: proto_chatlog_pb.RunResponse,
    requestSerialize: serialize_chatlog_RunRequest,
    requestDeserialize: deserialize_chatlog_RunRequest,
    responseSerialize: serialize_chatlog_RunResponse,
    responseDeserialize: deserialize_chatlog_RunResponse,
  },
  switch: {
    path: '/chatlog.ManagerService/Switch',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.SwitchRequest,
    responseType: proto_chatlog_pb.SwitchResponse,
    requestSerialize: serialize_chatlog_SwitchRequest,
    requestDeserialize: deserialize_chatlog_SwitchRequest,
    responseSerialize: serialize_chatlog_SwitchResponse,
    responseDeserialize: deserialize_chatlog_SwitchResponse,
  },
  startService: {
    path: '/chatlog.ManagerService/StartService',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.StartServiceRequest,
    responseType: proto_chatlog_pb.StartServiceResponse,
    requestSerialize: serialize_chatlog_StartServiceRequest,
    requestDeserialize: deserialize_chatlog_StartServiceRequest,
    responseSerialize: serialize_chatlog_StartServiceResponse,
    responseDeserialize: deserialize_chatlog_StartServiceResponse,
  },
  stopService: {
    path: '/chatlog.ManagerService/StopService',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.StopServiceRequest,
    responseType: proto_chatlog_pb.StopServiceResponse,
    requestSerialize: serialize_chatlog_StopServiceRequest,
    requestDeserialize: deserialize_chatlog_StopServiceRequest,
    responseSerialize: serialize_chatlog_StopServiceResponse,
    responseDeserialize: deserialize_chatlog_StopServiceResponse,
  },
  setHTTPAddr: {
    path: '/chatlog.ManagerService/SetHTTPAddr',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.SetHTTPAddrRequest,
    responseType: proto_chatlog_pb.SetHTTPAddrResponse,
    requestSerialize: serialize_chatlog_SetHTTPAddrRequest,
    requestDeserialize: deserialize_chatlog_SetHTTPAddrRequest,
    responseSerialize: serialize_chatlog_SetHTTPAddrResponse,
    responseDeserialize: deserialize_chatlog_SetHTTPAddrResponse,
  },
  getDataKey: {
    path: '/chatlog.ManagerService/GetDataKey',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.GetDataKeyRequest,
    responseType: proto_chatlog_pb.GetDataKeyResponse,
    requestSerialize: serialize_chatlog_GetDataKeyRequest,
    requestDeserialize: deserialize_chatlog_GetDataKeyRequest,
    responseSerialize: serialize_chatlog_GetDataKeyResponse,
    responseDeserialize: deserialize_chatlog_GetDataKeyResponse,
  },
  decryptDBFiles: {
    path: '/chatlog.ManagerService/DecryptDBFiles',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.DecryptDBFilesRequest,
    responseType: proto_chatlog_pb.DecryptDBFilesResponse,
    requestSerialize: serialize_chatlog_DecryptDBFilesRequest,
    requestDeserialize: deserialize_chatlog_DecryptDBFilesRequest,
    responseSerialize: serialize_chatlog_DecryptDBFilesResponse,
    responseDeserialize: deserialize_chatlog_DecryptDBFilesResponse,
  },
  startAutoDecrypt: {
    path: '/chatlog.ManagerService/StartAutoDecrypt',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.StartAutoDecryptRequest,
    responseType: proto_chatlog_pb.StartAutoDecryptResponse,
    requestSerialize: serialize_chatlog_StartAutoDecryptRequest,
    requestDeserialize: deserialize_chatlog_StartAutoDecryptRequest,
    responseSerialize: serialize_chatlog_StartAutoDecryptResponse,
    responseDeserialize: deserialize_chatlog_StartAutoDecryptResponse,
  },
  stopAutoDecrypt: {
    path: '/chatlog.ManagerService/StopAutoDecrypt',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.StopAutoDecryptRequest,
    responseType: proto_chatlog_pb.StopAutoDecryptResponse,
    requestSerialize: serialize_chatlog_StopAutoDecryptRequest,
    requestDeserialize: deserialize_chatlog_StopAutoDecryptRequest,
    responseSerialize: serialize_chatlog_StopAutoDecryptResponse,
    responseDeserialize: deserialize_chatlog_StopAutoDecryptResponse,
  },
  refreshSession: {
    path: '/chatlog.ManagerService/RefreshSession',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.RefreshSessionRequest,
    responseType: proto_chatlog_pb.RefreshSessionResponse,
    requestSerialize: serialize_chatlog_RefreshSessionRequest,
    requestDeserialize: deserialize_chatlog_RefreshSessionRequest,
    responseSerialize: serialize_chatlog_RefreshSessionResponse,
    responseDeserialize: deserialize_chatlog_RefreshSessionResponse,
  },
  commandKey: {
    path: '/chatlog.ManagerService/CommandKey',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.CommandKeyRequest,
    responseType: proto_chatlog_pb.CommandKeyResponse,
    requestSerialize: serialize_chatlog_CommandKeyRequest,
    requestDeserialize: deserialize_chatlog_CommandKeyRequest,
    responseSerialize: serialize_chatlog_CommandKeyResponse,
    responseDeserialize: deserialize_chatlog_CommandKeyResponse,
  },
  commandDecrypt: {
    path: '/chatlog.ManagerService/CommandDecrypt',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.CommandDecryptRequest,
    responseType: proto_chatlog_pb.CommandDecryptResponse,
    requestSerialize: serialize_chatlog_CommandDecryptRequest,
    requestDeserialize: deserialize_chatlog_CommandDecryptRequest,
    responseSerialize: serialize_chatlog_CommandDecryptResponse,
    responseDeserialize: deserialize_chatlog_CommandDecryptResponse,
  },
  commandHTTPServer: {
    path: '/chatlog.ManagerService/CommandHTTPServer',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.CommandHTTPServerRequest,
    responseType: proto_chatlog_pb.CommandHTTPServerResponse,
    requestSerialize: serialize_chatlog_CommandHTTPServerRequest,
    requestDeserialize: deserialize_chatlog_CommandHTTPServerRequest,
    responseSerialize: serialize_chatlog_CommandHTTPServerResponse,
    responseDeserialize: deserialize_chatlog_CommandHTTPServerResponse,
  },
  getWeChatInstances: {
    path: '/chatlog.ManagerService/GetWeChatInstances',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.GetWeChatInstancesRequest,
    responseType: proto_chatlog_pb.GetWeChatInstancesResponse,
    requestSerialize: serialize_chatlog_GetWeChatInstancesRequest,
    requestDeserialize: deserialize_chatlog_GetWeChatInstancesRequest,
    responseSerialize: serialize_chatlog_GetWeChatInstancesResponse,
    responseDeserialize: deserialize_chatlog_GetWeChatInstancesResponse,
  },
  getKey: {
    path: '/chatlog.ManagerService/GetKey',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.GetKeyRequest,
    responseType: proto_chatlog_pb.GetKeyResponse,
    requestSerialize: serialize_chatlog_GetKeyRequest,
    requestDeserialize: deserialize_chatlog_GetKeyRequest,
    responseSerialize: serialize_chatlog_GetKeyResponse,
    responseDeserialize: deserialize_chatlog_GetKeyResponse,
  },
  decrypt: {
    path: '/chatlog.ManagerService/Decrypt',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.DecryptRequest,
    responseType: proto_chatlog_pb.DecryptResponse,
    requestSerialize: serialize_chatlog_DecryptRequest,
    requestDeserialize: deserialize_chatlog_DecryptRequest,
    responseSerialize: serialize_chatlog_DecryptResponse,
    responseDeserialize: deserialize_chatlog_DecryptResponse,
  },
  backup: {
    path: '/chatlog.ManagerService/Backup',
    requestStream: false,
    responseStream: false,
    requestType: proto_chatlog_pb.BackupRequest,
    responseType: proto_chatlog_pb.BackupResponse,
    requestSerialize: serialize_chatlog_BackupRequest,
    requestDeserialize: deserialize_chatlog_BackupRequest,
    responseSerialize: serialize_chatlog_BackupResponse,
    responseDeserialize: deserialize_chatlog_BackupResponse,
  },
};

exports.ManagerServiceClient = grpc.makeGenericClientConstructor(ManagerServiceService, 'ManagerService');
