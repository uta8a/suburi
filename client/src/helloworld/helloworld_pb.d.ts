import * as jspb from 'google-protobuf'



export class HelloRequest extends jspb.Message {
  getName(): string;
  setName(value: string): HelloRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HelloRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HelloRequest): HelloRequest.AsObject;
  static serializeBinaryToWriter(message: HelloRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HelloRequest;
  static deserializeBinaryFromReader(message: HelloRequest, reader: jspb.BinaryReader): HelloRequest;
}

export namespace HelloRequest {
  export type AsObject = {
    name: string,
  }
}

export class HelloReply extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): HelloReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HelloReply.AsObject;
  static toObject(includeInstance: boolean, msg: HelloReply): HelloReply.AsObject;
  static serializeBinaryToWriter(message: HelloReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HelloReply;
  static deserializeBinaryFromReader(message: HelloReply, reader: jspb.BinaryReader): HelloReply;
}

export namespace HelloReply {
  export type AsObject = {
    message: string,
  }
}

export class TellMeSecretRequest extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): TellMeSecretRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TellMeSecretRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TellMeSecretRequest): TellMeSecretRequest.AsObject;
  static serializeBinaryToWriter(message: TellMeSecretRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TellMeSecretRequest;
  static deserializeBinaryFromReader(message: TellMeSecretRequest, reader: jspb.BinaryReader): TellMeSecretRequest;
}

export namespace TellMeSecretRequest {
  export type AsObject = {
    message: string,
  }
}

export class TellMeSecretResponse extends jspb.Message {
  getAnswer(): string;
  setAnswer(value: string): TellMeSecretResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TellMeSecretResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TellMeSecretResponse): TellMeSecretResponse.AsObject;
  static serializeBinaryToWriter(message: TellMeSecretResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TellMeSecretResponse;
  static deserializeBinaryFromReader(message: TellMeSecretResponse, reader: jspb.BinaryReader): TellMeSecretResponse;
}

export namespace TellMeSecretResponse {
  export type AsObject = {
    answer: string,
  }
}

