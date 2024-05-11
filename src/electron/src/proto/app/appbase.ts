/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "sierra_app";

export interface Sample {
  sha256: string;
  format: string;
  sourceUri: string;
  uri: string;
}

function createBaseSample(): Sample {
  return { sha256: "", format: "", sourceUri: "", uri: "" };
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
