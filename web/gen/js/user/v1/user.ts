// @generated by protobuf-ts 2.8.3 with parameter long_type_string
// @generated from protobuf file "user/v1/user.proto" (package "user.v1", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message user.v1.PasswordLoginRequest
 */
export interface PasswordLoginRequest {
    /**
     * @generated from protobuf field: string email = 1;
     */
    email: string;
    /**
     * @generated from protobuf field: string password = 2;
     */
    password: string;
}
/**
 * @generated from protobuf message user.v1.PasswordLoginResponse
 */
export interface PasswordLoginResponse {
    /**
     * @generated from protobuf field: string email = 1;
     */
    email: string;
    /**
     * @generated from protobuf field: string user_id = 2;
     */
    userId: string;
    /**
     * @generated from protobuf field: bool admin = 3;
     */
    admin: boolean;
}
// @generated message type with reflection information, may provide speed optimized methods
class PasswordLoginRequest$Type extends MessageType<PasswordLoginRequest> {
    constructor() {
        super("user.v1.PasswordLoginRequest", [
            { no: 1, name: "email", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "password", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<PasswordLoginRequest>): PasswordLoginRequest {
        const message = { email: "", password: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<PasswordLoginRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PasswordLoginRequest): PasswordLoginRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string email */ 1:
                    message.email = reader.string();
                    break;
                case /* string password */ 2:
                    message.password = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: PasswordLoginRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string email = 1; */
        if (message.email !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.email);
        /* string password = 2; */
        if (message.password !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.password);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message user.v1.PasswordLoginRequest
 */
export const PasswordLoginRequest = new PasswordLoginRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PasswordLoginResponse$Type extends MessageType<PasswordLoginResponse> {
    constructor() {
        super("user.v1.PasswordLoginResponse", [
            { no: 1, name: "email", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "user_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "admin", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<PasswordLoginResponse>): PasswordLoginResponse {
        const message = { email: "", userId: "", admin: false };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<PasswordLoginResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PasswordLoginResponse): PasswordLoginResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string email */ 1:
                    message.email = reader.string();
                    break;
                case /* string user_id */ 2:
                    message.userId = reader.string();
                    break;
                case /* bool admin */ 3:
                    message.admin = reader.bool();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: PasswordLoginResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string email = 1; */
        if (message.email !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.email);
        /* string user_id = 2; */
        if (message.userId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.userId);
        /* bool admin = 3; */
        if (message.admin !== false)
            writer.tag(3, WireType.Varint).bool(message.admin);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message user.v1.PasswordLoginResponse
 */
export const PasswordLoginResponse = new PasswordLoginResponse$Type();
/**
 * @generated ServiceType for protobuf service user.v1.UserService
 */
export const UserService = new ServiceType("user.v1.UserService", [
    { name: "PasswordLogin", options: {}, I: PasswordLoginRequest, O: PasswordLoginResponse }
]);
