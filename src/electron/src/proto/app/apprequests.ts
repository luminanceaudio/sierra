/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Sample } from "./appbase";

export const protobufPackage = "sierra_app";

export interface GetSamplesResponse {
  samples: Sample[];
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

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };
