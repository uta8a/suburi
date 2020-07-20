/**
 * @fileoverview gRPC-Web generated client stub for check
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as check_pb from './check_pb';


export class RoutesClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoHealthCheck = new grpcWeb.AbstractClientBase.MethodInfo(
    check_pb.Response,
    (request: check_pb.Request) => {
      return request.serializeBinary();
    },
    check_pb.Response.deserializeBinary
  );

  healthCheck(
    request: check_pb.Request,
    metadata: grpcWeb.Metadata | null): Promise<check_pb.Response>;

  healthCheck(
    request: check_pb.Request,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: check_pb.Response) => void): grpcWeb.ClientReadableStream<check_pb.Response>;

  healthCheck(
    request: check_pb.Request,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: check_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/check.Routes/HealthCheck',
        request,
        metadata || {},
        this.methodInfoHealthCheck,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/check.Routes/HealthCheck',
    request,
    metadata || {},
    this.methodInfoHealthCheck);
  }

  methodInfoSecretCheck = new grpcWeb.AbstractClientBase.MethodInfo(
    check_pb.Response,
    (request: check_pb.Request) => {
      return request.serializeBinary();
    },
    check_pb.Response.deserializeBinary
  );

  secretCheck(
    request: check_pb.Request,
    metadata: grpcWeb.Metadata | null): Promise<check_pb.Response>;

  secretCheck(
    request: check_pb.Request,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: check_pb.Response) => void): grpcWeb.ClientReadableStream<check_pb.Response>;

  secretCheck(
    request: check_pb.Request,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: check_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/check.Routes/SecretCheck',
        request,
        metadata || {},
        this.methodInfoSecretCheck,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/check.Routes/SecretCheck',
    request,
    metadata || {},
    this.methodInfoSecretCheck);
  }

}

