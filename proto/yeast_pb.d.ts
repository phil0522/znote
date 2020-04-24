import * as jspb from "google-protobuf"

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';

export class YTag extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getArchived(): boolean;
  setArchived(value: boolean): void;

  getImplicitTagsList(): Array<string>;
  setImplicitTagsList(value: Array<string>): void;
  clearImplicitTagsList(): void;
  addImplicitTags(value: string, index?: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YTag.AsObject;
  static toObject(includeInstance: boolean, msg: YTag): YTag.AsObject;
  static serializeBinaryToWriter(message: YTag, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YTag;
  static deserializeBinaryFromReader(message: YTag, reader: jspb.BinaryReader): YTag;
}

export namespace YTag {
  export type AsObject = {
    name: string,
    archived: boolean,
    implicitTagsList: Array<string>,
  }
}

export class YNote extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getTitle(): string;
  setTitle(value: string): void;

  getSummary(): string;
  setSummary(value: string): void;

  getContent(): string;
  setContent(value: string): void;

  getTagsList(): Array<string>;
  setTagsList(value: Array<string>): void;
  clearTagsList(): void;
  addTags(value: string, index?: number): void;

  getImplicitTagsList(): Array<string>;
  setImplicitTagsList(value: Array<string>): void;
  clearImplicitTagsList(): void;
  addImplicitTags(value: string, index?: number): void;

  getCreatedTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedTime(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasCreatedTime(): boolean;
  clearCreatedTime(): void;

  getUpdatedTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedTime(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasUpdatedTime(): boolean;
  clearUpdatedTime(): void;

  getArchived(): boolean;
  setArchived(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YNote.AsObject;
  static toObject(includeInstance: boolean, msg: YNote): YNote.AsObject;
  static serializeBinaryToWriter(message: YNote, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YNote;
  static deserializeBinaryFromReader(message: YNote, reader: jspb.BinaryReader): YNote;
}

export namespace YNote {
  export type AsObject = {
    id: string,
    title: string,
    summary: string,
    content: string,
    tagsList: Array<string>,
    implicitTagsList: Array<string>,
    createdTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    archived: boolean,
  }
}

export class YeastUpsertTagRequest extends jspb.Message {
  getYtag(): YTag | undefined;
  setYtag(value?: YTag): void;
  hasYtag(): boolean;
  clearYtag(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YeastUpsertTagRequest.AsObject;
  static toObject(includeInstance: boolean, msg: YeastUpsertTagRequest): YeastUpsertTagRequest.AsObject;
  static serializeBinaryToWriter(message: YeastUpsertTagRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YeastUpsertTagRequest;
  static deserializeBinaryFromReader(message: YeastUpsertTagRequest, reader: jspb.BinaryReader): YeastUpsertTagRequest;
}

export namespace YeastUpsertTagRequest {
  export type AsObject = {
    ytag?: YTag.AsObject,
  }
}

export class YeastUpsertTagResponse extends jspb.Message {
  getStatus(): string;
  setStatus(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YeastUpsertTagResponse.AsObject;
  static toObject(includeInstance: boolean, msg: YeastUpsertTagResponse): YeastUpsertTagResponse.AsObject;
  static serializeBinaryToWriter(message: YeastUpsertTagResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YeastUpsertTagResponse;
  static deserializeBinaryFromReader(message: YeastUpsertTagResponse, reader: jspb.BinaryReader): YeastUpsertTagResponse;
}

export namespace YeastUpsertTagResponse {
  export type AsObject = {
    status: string,
  }
}

export class YeastListTagsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YeastListTagsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: YeastListTagsRequest): YeastListTagsRequest.AsObject;
  static serializeBinaryToWriter(message: YeastListTagsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YeastListTagsRequest;
  static deserializeBinaryFromReader(message: YeastListTagsRequest, reader: jspb.BinaryReader): YeastListTagsRequest;
}

export namespace YeastListTagsRequest {
  export type AsObject = {
  }
}

export class YeastListTagsResponse extends jspb.Message {
  getYtagList(): Array<YTag>;
  setYtagList(value: Array<YTag>): void;
  clearYtagList(): void;
  addYtag(value?: YTag, index?: number): YTag;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YeastListTagsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: YeastListTagsResponse): YeastListTagsResponse.AsObject;
  static serializeBinaryToWriter(message: YeastListTagsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YeastListTagsResponse;
  static deserializeBinaryFromReader(message: YeastListTagsResponse, reader: jspb.BinaryReader): YeastListTagsResponse;
}

export namespace YeastListTagsResponse {
  export type AsObject = {
    ytagList: Array<YTag.AsObject>,
  }
}

export class YeastUpsertNoteRequest extends jspb.Message {
  getYnote(): YNote | undefined;
  setYnote(value?: YNote): void;
  hasYnote(): boolean;
  clearYnote(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YeastUpsertNoteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: YeastUpsertNoteRequest): YeastUpsertNoteRequest.AsObject;
  static serializeBinaryToWriter(message: YeastUpsertNoteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YeastUpsertNoteRequest;
  static deserializeBinaryFromReader(message: YeastUpsertNoteRequest, reader: jspb.BinaryReader): YeastUpsertNoteRequest;
}

export namespace YeastUpsertNoteRequest {
  export type AsObject = {
    ynote?: YNote.AsObject,
  }
}

export class YeastUpsertNoteResponse extends jspb.Message {
  getStatus(): string;
  setStatus(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YeastUpsertNoteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: YeastUpsertNoteResponse): YeastUpsertNoteResponse.AsObject;
  static serializeBinaryToWriter(message: YeastUpsertNoteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YeastUpsertNoteResponse;
  static deserializeBinaryFromReader(message: YeastUpsertNoteResponse, reader: jspb.BinaryReader): YeastUpsertNoteResponse;
}

export namespace YeastUpsertNoteResponse {
  export type AsObject = {
    status: string,
  }
}

export class YeastListNoteRequest extends jspb.Message {
  getUsage(): string;
  setUsage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YeastListNoteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: YeastListNoteRequest): YeastListNoteRequest.AsObject;
  static serializeBinaryToWriter(message: YeastListNoteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YeastListNoteRequest;
  static deserializeBinaryFromReader(message: YeastListNoteRequest, reader: jspb.BinaryReader): YeastListNoteRequest;
}

export namespace YeastListNoteRequest {
  export type AsObject = {
    usage: string,
  }
}

export class YeastListNoteResponse extends jspb.Message {
  getYnoteList(): Array<YNote>;
  setYnoteList(value: Array<YNote>): void;
  clearYnoteList(): void;
  addYnote(value?: YNote, index?: number): YNote;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): YeastListNoteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: YeastListNoteResponse): YeastListNoteResponse.AsObject;
  static serializeBinaryToWriter(message: YeastListNoteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): YeastListNoteResponse;
  static deserializeBinaryFromReader(message: YeastListNoteResponse, reader: jspb.BinaryReader): YeastListNoteResponse;
}

export namespace YeastListNoteResponse {
  export type AsObject = {
    ynoteList: Array<YNote.AsObject>,
  }
}

