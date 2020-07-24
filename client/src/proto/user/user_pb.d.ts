import * as jspb from 'google-protobuf'



export class Request extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): Request;

  getPassword(): string;
  setPassword(value: string): Request;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Request.AsObject;
  static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
  static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Request;
  static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
}

export namespace Request {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class Result extends jspb.Message {
  getToken(): string;
  setToken(value: string): Result;

  getDisplayMessage(): string;
  setDisplayMessage(value: string): Result;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Result.AsObject;
  static toObject(includeInstance: boolean, msg: Result): Result.AsObject;
  static serializeBinaryToWriter(message: Result, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Result;
  static deserializeBinaryFromReader(message: Result, reader: jspb.BinaryReader): Result;
}

export namespace Result {
  export type AsObject = {
    token: string,
    displayMessage: string,
  }
}

