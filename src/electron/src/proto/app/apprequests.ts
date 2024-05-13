/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Sample, Source } from "./appbase";

export const protobufPackage = "sierra_app";

export interface GetSamplesResponse {
  samples: Sample[];
}

export interface GetSourcesResponse {
  sources: Source[];
}

export interface CreateSourceRequest {
  uri: string;
}

export interface CreateSourceResponse {
}

export interface DeleteSourceRequest {
  uri: string;
}

export interface DeleteSourceResponse {
}

function createBaseGetSamplesResponse(): GetSamplesResponse {
  return { samples: [] };
}

export const GetSamplesResponse = {
  encode(message: GetSamplesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.samples) {
      Sample.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetSamplesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetSamplesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.samples.push(Sample.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetSamplesResponse {
    return { samples: Array.isArray(object?.samples) ? object.samples.map((e: any) => Sample.fromJSON(e)) : [] };
  },

  toJSON(message: GetSamplesResponse): unknown {
    const obj: any = {};
    if (message.samples?.length) {
      obj.samples = message.samples.map((e) => Sample.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetSamplesResponse>, I>>(base?: I): GetSamplesResponse {
    return GetSamplesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetSamplesResponse>, I>>(object: I): GetSamplesResponse {
    const message = createBaseGetSamplesResponse();
    message.samples = object.samples?.map((e) => Sample.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetSourcesResponse(): GetSourcesResponse {
  return { sources: [] };
}

export const GetSourcesResponse = {
  encode(message: GetSourcesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.sources) {
      Source.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetSourcesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetSourcesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.sources.push(Source.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetSourcesResponse {
    return { sources: Array.isArray(object?.sources) ? object.sources.map((e: any) => Source.fromJSON(e)) : [] };
  },

  toJSON(message: GetSourcesResponse): unknown {
    const obj: any = {};
    if (message.sources?.length) {
      obj.sources = message.sources.map((e) => Source.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetSourcesResponse>, I>>(base?: I): GetSourcesResponse {
    return GetSourcesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetSourcesResponse>, I>>(object: I): GetSourcesResponse {
    const message = createBaseGetSourcesResponse();
    message.sources = object.sources?.map((e) => Source.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreateSourceRequest(): CreateSourceRequest {
  return { uri: "" };
}

export const CreateSourceRequest = {
  encode(message: CreateSourceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uri !== "") {
      writer.uint32(10).string(message.uri);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateSourceRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateSourceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.uri = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateSourceRequest {
    return { uri: isSet(object.uri) ? String(object.uri) : "" };
  },

  toJSON(message: CreateSourceRequest): unknown {
    const obj: any = {};
    if (message.uri !== "") {
      obj.uri = message.uri;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateSourceRequest>, I>>(base?: I): CreateSourceRequest {
    return CreateSourceRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateSourceRequest>, I>>(object: I): CreateSourceRequest {
    const message = createBaseCreateSourceRequest();
    message.uri = object.uri ?? "";
    return message;
  },
};

function createBaseCreateSourceResponse(): CreateSourceResponse {
  return {};
}

export const CreateSourceResponse = {
  encode(_: CreateSourceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateSourceResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateSourceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): CreateSourceResponse {
    return {};
  },

  toJSON(_: CreateSourceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateSourceResponse>, I>>(base?: I): CreateSourceResponse {
    return CreateSourceResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateSourceResponse>, I>>(_: I): CreateSourceResponse {
    const message = createBaseCreateSourceResponse();
    return message;
  },
};

function createBaseDeleteSourceRequest(): DeleteSourceRequest {
  return { uri: "" };
}

export const DeleteSourceRequest = {
  encode(message: DeleteSourceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uri !== "") {
      writer.uint32(10).string(message.uri);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteSourceRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteSourceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.uri = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteSourceRequest {
    return { uri: isSet(object.uri) ? String(object.uri) : "" };
  },

  toJSON(message: DeleteSourceRequest): unknown {
    const obj: any = {};
    if (message.uri !== "") {
      obj.uri = message.uri;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteSourceRequest>, I>>(base?: I): DeleteSourceRequest {
    return DeleteSourceRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteSourceRequest>, I>>(object: I): DeleteSourceRequest {
    const message = createBaseDeleteSourceRequest();
    message.uri = object.uri ?? "";
    return message;
  },
};

function createBaseDeleteSourceResponse(): DeleteSourceResponse {
  return {};
}

export const DeleteSourceResponse = {
  encode(_: DeleteSourceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteSourceResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteSourceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): DeleteSourceResponse {
    return {};
  },

  toJSON(_: DeleteSourceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteSourceResponse>, I>>(base?: I): DeleteSourceResponse {
    return DeleteSourceResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteSourceResponse>, I>>(_: I): DeleteSourceResponse {
    const message = createBaseDeleteSourceResponse();
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
