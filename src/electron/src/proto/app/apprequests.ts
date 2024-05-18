/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import {
  Sample,
  SortColumn_Enum,
  sortColumn_EnumFromJSON,
  sortColumn_EnumToJSON,
  SortDirection_Enum,
  sortDirection_EnumFromJSON,
  sortDirection_EnumToJSON,
  Source,
} from "./appbase";

export const protobufPackage = "sierra_app";

export interface QuerySamplesRequest {
  query: string;
  page: number;
  size: number;
  sortDirection: SortDirection_Enum;
  sortColumn: SortColumn_Enum;
}

export interface QuerySamplesResponse {
  samples: Sample[];
}

export interface QuerySamplesCountRequest {
  query: string;
}

export interface QuerySamplesCountResponse {
  count: number;
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

function createBaseQuerySamplesRequest(): QuerySamplesRequest {
  return { query: "", page: 0, size: 0, sortDirection: 0, sortColumn: 0 };
}

export const QuerySamplesRequest = {
  encode(message: QuerySamplesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.query !== "") {
      writer.uint32(10).string(message.query);
    }
    if (message.page !== 0) {
      writer.uint32(16).int32(message.page);
    }
    if (message.size !== 0) {
      writer.uint32(24).int32(message.size);
    }
    if (message.sortDirection !== 0) {
      writer.uint32(32).int32(message.sortDirection);
    }
    if (message.sortColumn !== 0) {
      writer.uint32(40).int32(message.sortColumn);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerySamplesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerySamplesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.query = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.page = reader.int32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.size = reader.int32();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.sortDirection = reader.int32() as any;
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.sortColumn = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QuerySamplesRequest {
    return {
      query: isSet(object.query) ? String(object.query) : "",
      page: isSet(object.page) ? Number(object.page) : 0,
      size: isSet(object.size) ? Number(object.size) : 0,
      sortDirection: isSet(object.sortDirection) ? sortDirection_EnumFromJSON(object.sortDirection) : 0,
      sortColumn: isSet(object.sortColumn) ? sortColumn_EnumFromJSON(object.sortColumn) : 0,
    };
  },

  toJSON(message: QuerySamplesRequest): unknown {
    const obj: any = {};
    if (message.query !== "") {
      obj.query = message.query;
    }
    if (message.page !== 0) {
      obj.page = Math.round(message.page);
    }
    if (message.size !== 0) {
      obj.size = Math.round(message.size);
    }
    if (message.sortDirection !== 0) {
      obj.sortDirection = sortDirection_EnumToJSON(message.sortDirection);
    }
    if (message.sortColumn !== 0) {
      obj.sortColumn = sortColumn_EnumToJSON(message.sortColumn);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QuerySamplesRequest>, I>>(base?: I): QuerySamplesRequest {
    return QuerySamplesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QuerySamplesRequest>, I>>(object: I): QuerySamplesRequest {
    const message = createBaseQuerySamplesRequest();
    message.query = object.query ?? "";
    message.page = object.page ?? 0;
    message.size = object.size ?? 0;
    message.sortDirection = object.sortDirection ?? 0;
    message.sortColumn = object.sortColumn ?? 0;
    return message;
  },
};

function createBaseQuerySamplesResponse(): QuerySamplesResponse {
  return { samples: [] };
}

export const QuerySamplesResponse = {
  encode(message: QuerySamplesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.samples) {
      Sample.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerySamplesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerySamplesResponse();
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

  fromJSON(object: any): QuerySamplesResponse {
    return { samples: Array.isArray(object?.samples) ? object.samples.map((e: any) => Sample.fromJSON(e)) : [] };
  },

  toJSON(message: QuerySamplesResponse): unknown {
    const obj: any = {};
    if (message.samples?.length) {
      obj.samples = message.samples.map((e) => Sample.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QuerySamplesResponse>, I>>(base?: I): QuerySamplesResponse {
    return QuerySamplesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QuerySamplesResponse>, I>>(object: I): QuerySamplesResponse {
    const message = createBaseQuerySamplesResponse();
    message.samples = object.samples?.map((e) => Sample.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQuerySamplesCountRequest(): QuerySamplesCountRequest {
  return { query: "" };
}

export const QuerySamplesCountRequest = {
  encode(message: QuerySamplesCountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.query !== "") {
      writer.uint32(10).string(message.query);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerySamplesCountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerySamplesCountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.query = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QuerySamplesCountRequest {
    return { query: isSet(object.query) ? String(object.query) : "" };
  },

  toJSON(message: QuerySamplesCountRequest): unknown {
    const obj: any = {};
    if (message.query !== "") {
      obj.query = message.query;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QuerySamplesCountRequest>, I>>(base?: I): QuerySamplesCountRequest {
    return QuerySamplesCountRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QuerySamplesCountRequest>, I>>(object: I): QuerySamplesCountRequest {
    const message = createBaseQuerySamplesCountRequest();
    message.query = object.query ?? "";
    return message;
  },
};

function createBaseQuerySamplesCountResponse(): QuerySamplesCountResponse {
  return { count: 0 };
}

export const QuerySamplesCountResponse = {
  encode(message: QuerySamplesCountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.count !== 0) {
      writer.uint32(8).int64(message.count);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerySamplesCountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerySamplesCountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.count = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QuerySamplesCountResponse {
    return { count: isSet(object.count) ? Number(object.count) : 0 };
  },

  toJSON(message: QuerySamplesCountResponse): unknown {
    const obj: any = {};
    if (message.count !== 0) {
      obj.count = Math.round(message.count);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QuerySamplesCountResponse>, I>>(base?: I): QuerySamplesCountResponse {
    return QuerySamplesCountResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QuerySamplesCountResponse>, I>>(object: I): QuerySamplesCountResponse {
    const message = createBaseQuerySamplesCountResponse();
    message.count = object.count ?? 0;
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

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
