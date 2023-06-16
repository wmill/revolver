"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserServiceClient = void 0;
const user_1 = require("./user");
const runtime_rpc_1 = require("@protobuf-ts/runtime-rpc");
/**
 * @generated from protobuf service user.v1.UserService
 */
class UserServiceClient {
    constructor(_transport) {
        this._transport = _transport;
        this.typeName = user_1.UserService.typeName;
        this.methods = user_1.UserService.methods;
        this.options = user_1.UserService.options;
    }
    /**
     * @generated from protobuf rpc: PasswordLogin(user.v1.PasswordLoginRequest) returns (user.v1.PasswordLoginResponse);
     */
    passwordLogin(input, options) {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return (0, runtime_rpc_1.stackIntercept)("unary", this._transport, method, opt, input);
    }
}
exports.UserServiceClient = UserServiceClient;
