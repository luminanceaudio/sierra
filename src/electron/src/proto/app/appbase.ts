/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "./google/protobuf/timestamp";

export const protobufPackage = "sierra_app";

export interface Sample {
  sha256: string;
  format: string;
  sourceUri: string;
  uri: string;
  duration: number;
  added: boolean;
}

export interface Source {
  uri: string;
  createTime: Date | undefined;
}

export interface SortDirection {
}

export enum SortDirection_Enum {
  Asc = 0,
  Desc = 1,
  UNRECOGNIZED = -1,
}

export function sortDirection_EnumFromJSON(object: any): SortDirection_Enum {
  switch (object) {
    case 0:
    case "Asc":
      return SortDirection_Enum.Asc;
    case 1:
    case "Desc":
      return SortDirection_Enum.Desc;
    case -1:
    case "UNRECOGNIZED":
    default:
      return SortDirection_Enum.UNRECOGNIZED;
  }
}

export function sortDirection_EnumToJSON(object: SortDirection_Enum): string {
  switch (object) {
    case SortDirection_Enum.Asc:
      return "Asc";
    case SortDirection_Enum.Desc:
      return "Desc";
    case SortDirection_Enum.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface SortColumn {
}

export enum SortColumn_Enum {
  None = 0,
  Name = 1,
  Duration = 2,
  UNRECOGNIZED = -1,
}

export function sortColumn_EnumFromJSON(object: any): SortColumn_Enum {
  switch (object) {
    case 0:
    case "None":
      return SortColumn_Enum.None;
    case 1:
    case "Name":
      return SortColumn_Enum.Name;
    case 2:
    case "Duration":
      return SortColumn_Enum.Duration;
    case -1:
    case "UNRECOGNIZED":
    default:
      return SortColumn_Enum.UNRECOGNIZED;
  }
}

export function sortColumn_EnumToJSON(object: SortColumn_Enum): string {
  switch (object) {
    case SortColumn_Enum.None:
      return "None";
    case SortColumn_Enum.Name:
      return "Name";
    case SortColumn_Enum.Duration:
      return "Duration";
    case SortColumn_Enum.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

function createBaseSample(): Sample {
  return { sha256: "", format: "", sourceUri: "", uri: "", duration: 0, added: false };
}

export const Sample = {
  encode(message: Sample, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sha256 !== "") {
      writer.uint32(10).string(message.sha256);
    }
    if (message.format !== "") {
      writer.uint32(18).string(message.format);
    }
    if (message.sourceUri !== "") {
      writer.uint32(26).string(message.sourceUri);
    }
    if (message.uri !== "") {
      writer.uint32(34).string(message.uri);
    }
    if (message.duration !== 0) {
      writer.uint32(40).int64(message.duration);
    }
    if (message.added === true) {
      writer.uint32(48).bool(message.added);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Sample {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSample();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.sha256 = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.format = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.sourceUri = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.uri = reader.string();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.duration = longToNumber(reader.int64() as Long);
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.added = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Sample {
    return {
      sha256: isSet(object.sha256) ? String(object.sha256) : "",
      format: isSet(object.format) ? String(object.format) : "",
      sourceUri: isSet(object.sourceUri) ? String(object.sourceUri) : "",
      uri: isSet(object.uri) ? String(object.uri) : "",
      duration: isSet(object.duration) ? Number(object.duration) : 0,
      added: isSet(object.added) ? Boolean(object.added) : false,
    };
  },

  toJSON(message: Sample): unknown {
    const obj: any = {};
    if (message.sha256 !== "") {
      obj.sha256 = message.sha256;
    }
    if (message.format !== "") {
      obj.format = message.format;
    }
    if (message.sourceUri !== "") {
      obj.sourceUri = message.sourceUri;
    }
    if (message.uri !== "") {
      obj.uri = message.uri;
    }
    if (message.duration !== 0) {
      obj.duration = Math.round(message.duration);
    }
    if (message.added === true) {
      obj.added = message.added;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Sample>, I>>(base?: I): Sample {
    return Sample.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Sample>, I>>(object: I): Sample {
    const message = createBaseSample();
    message.sha256 = object.sha256 ?? "";
    message.format = object.format ?? "";
    message.sourceUri = object.sourceUri ?? "";
    message.uri = object.uri ?? "";
    message.duration = object.duration ?? 0;
    message.added = object.added ?? false;
    return message;
  },
};

function createBaseSource(): Source {
  return { uri: "", createTime: undefined };
}

export const Source = {
  encode(message: Source, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uri !== "") {
      writer.uint32(10).string(message.uri);
    }
    if (message.createTime !== undefined) {
      Timestamp.encode(toTimestamp(message.createTime), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Source {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSource();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.uri = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.createTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Source {
    return {
      uri: isSet(object.uri) ? String(object.uri) : "",
      createTime: isSet(object.createTime) ? fromJsonTimestamp(object.createTime) : undefined,
    };
  },

  toJSON(message: Source): unknown {
    const obj: any = {};
    if (message.uri !== "") {
      obj.uri = message.uri;
    }
    if (message.createTime !== undefined) {
      obj.createTime = message.createTime.toISOString();
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Source>, I>>(base?: I): Source {
    return Source.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Source>, I>>(object: I): Source {
    const message = createBaseSource();
    message.uri = object.uri ?? "";
    message.createTime = object.createTime ?? undefined;
    return message;
  },
};

function createBaseSortDirection(): SortDirection {
  return {};
}

export const SortDirection = {
  encode(_: SortDirection, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SortDirection {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSortDirection();
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

  fromJSON(_: any): SortDirection {
    return {};
  },

  toJSON(_: SortDirection): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<SortDirection>, I>>(base?: I): SortDirection {
    return SortDirection.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SortDirection>, I>>(_: I): SortDirection {
    const message = createBaseSortDirection();
    return message;
  },
};

function createBaseSortColumn(): SortColumn {
  return {};
}

export const SortColumn = {
  encode(_: SortColumn, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SortColumn {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSortColumn();
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

  fromJSON(_: any): SortColumn {
    return {};
  },

  toJSON(_: SortColumn): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<SortColumn>, I>>(base?: I): SortColumn {
    return SortColumn.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SortColumn>, I>>(_: I): SortColumn {
    const message = createBaseSortColumn();
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

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
