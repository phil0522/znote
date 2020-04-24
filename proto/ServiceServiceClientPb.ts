/**
 * @fileoverview gRPC-Web generated client stub for phil0522.znote
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  QuitServerRequest,
  QuitServerResponse,
  ZNoteRequest,
  ZNoteResponse} from './service_pb';

export class ZNoteServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoExecuteCommand = new grpcWeb.AbstractClientBase.MethodInfo(
    ZNoteResponse,
    (request: ZNoteRequest) => {
      return request.serializeBinary();
    },
    ZNoteResponse.deserializeBinary
  );

  executeCommand(
    request: ZNoteRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ZNoteResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/phil0522.znote.ZNoteService/ExecuteCommand',
      request,
      metadata || {},
      this.methodInfoExecuteCommand,
      callback);
  }

  methodInfoQuitServer = new grpcWeb.AbstractClientBase.MethodInfo(
    QuitServerResponse,
    (request: QuitServerRequest) => {
      return request.serializeBinary();
    },
    QuitServerResponse.deserializeBinary
  );

  quitServer(
    request: QuitServerRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: QuitServerResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/phil0522.znote.ZNoteService/QuitServer',
      request,
      metadata || {},
      this.methodInfoQuitServer,
      callback);
  }

}

