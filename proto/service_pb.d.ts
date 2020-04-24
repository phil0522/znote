import * as jspb from "google-protobuf"

export class ZNoteRequest extends jspb.Message {
  getCommand(): string;
  setCommand(value: string): void;

  getBook(): string;
  setBook(value: string): void;

  getNoteId(): string;
  setNoteId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZNoteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ZNoteRequest): ZNoteRequest.AsObject;
  static serializeBinaryToWriter(message: ZNoteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZNoteRequest;
  static deserializeBinaryFromReader(message: ZNoteRequest, reader: jspb.BinaryReader): ZNoteRequest;
}

export namespace ZNoteRequest {
  export type AsObject = {
    command: string,
    book: string,
    noteId: string,
  }
}

export class ZNoteResponse extends jspb.Message {
  getResult(): string;
  setResult(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ZNoteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ZNoteResponse): ZNoteResponse.AsObject;
  static serializeBinaryToWriter(message: ZNoteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ZNoteResponse;
  static deserializeBinaryFromReader(message: ZNoteResponse, reader: jspb.BinaryReader): ZNoteResponse;
}

export namespace ZNoteResponse {
  export type AsObject = {
    result: string,
  }
}

export class QuitServerRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QuitServerRequest.AsObject;
  static toObject(includeInstance: boolean, msg: QuitServerRequest): QuitServerRequest.AsObject;
  static serializeBinaryToWriter(message: QuitServerRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QuitServerRequest;
  static deserializeBinaryFromReader(message: QuitServerRequest, reader: jspb.BinaryReader): QuitServerRequest;
}

export namespace QuitServerRequest {
  export type AsObject = {
  }
}

export class QuitServerResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QuitServerResponse.AsObject;
  static toObject(includeInstance: boolean, msg: QuitServerResponse): QuitServerResponse.AsObject;
  static serializeBinaryToWriter(message: QuitServerResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QuitServerResponse;
  static deserializeBinaryFromReader(message: QuitServerResponse, reader: jspb.BinaryReader): QuitServerResponse;
}

export namespace QuitServerResponse {
  export type AsObject = {
  }
}

