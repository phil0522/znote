/**
 * @fileoverview gRPC-Web generated client stub for phil0522.znote
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';

import {
  YeastListNoteRequest,
  YeastListNoteResponse,
  YeastListTagsRequest,
  YeastUpsertNoteResponse,
  YeastUpsertTagRequest,
  YeastUpsertTagResponse} from './yeast_pb';

export class YeastServiceClient {
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

  methodInfoupsertTag = new grpcWeb.AbstractClientBase.MethodInfo(
    YeastUpsertTagResponse,
    (request: YeastUpsertTagRequest) => {
      return request.serializeBinary();
    },
    YeastUpsertTagResponse.deserializeBinary
  );

  upsertTag(
    request: YeastUpsertTagRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: YeastUpsertTagResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/phil0522.znote.YeastService/upsertTag',
      request,
      metadata || {},
      this.methodInfoupsertTag,
      callback);
  }

  methodInfoListTags = new grpcWeb.AbstractClientBase.MethodInfo(
    YeastListTagsRequest,
    (request: YeastListTagsRequest) => {
      return request.serializeBinary();
    },
    YeastListTagsRequest.deserializeBinary
  );

  listTags(
    request: YeastListTagsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: YeastListTagsRequest) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/phil0522.znote.YeastService/ListTags',
      request,
      metadata || {},
      this.methodInfoListTags,
      callback);
  }

  methodInfoupsertNote = new grpcWeb.AbstractClientBase.MethodInfo(
    YeastUpsertNoteResponse,
    (request: YeastUpsertTagResponse) => {
      return request.serializeBinary();
    },
    YeastUpsertNoteResponse.deserializeBinary
  );

  upsertNote(
    request: YeastUpsertTagResponse,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: YeastUpsertNoteResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/phil0522.znote.YeastService/upsertNote',
      request,
      metadata || {},
      this.methodInfoupsertNote,
      callback);
  }

  methodInfoListNotes = new grpcWeb.AbstractClientBase.MethodInfo(
    YeastListNoteResponse,
    (request: YeastListNoteRequest) => {
      return request.serializeBinary();
    },
    YeastListNoteResponse.deserializeBinary
  );

  listNotes(
    request: YeastListNoteRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: YeastListNoteResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/phil0522.znote.YeastService/ListNotes',
      request,
      metadata || {},
      this.methodInfoListNotes,
      callback);
  }

}

