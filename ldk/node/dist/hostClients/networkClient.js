"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.NetworkClient = void 0;
const baseClient_1 = __importDefault(require("./baseClient"));
const network_grpc_pb_1 = require("../grpc/network_grpc_pb");
const network_pb_1 = __importDefault(require("../grpc/network_pb"));
/**
 * @param values - a PB struct ListValue (list of Values)
 * @internal
 */
function parseHeaderValues(values) {
    return values.getValuesList().map((value) => {
        return value.getStringValue();
    });
}
/**
 * @param headersMap - A map of headers
 * @internal
 */
function parseHeadersMap(headersMap) {
    const record = {};
    headersMap.forEach((values, key) => {
        record[key] = parseHeaderValues(values);
    });
    return record;
}
/**
 * Adds headers to a request message
 *
 * @param message - the message that will be sent via gRPC - modified in place
 * @param headers - the headers to add to the request as a map
 * @internal
 */
function addHeadersToMessage(message, headers) {
    Object.entries(headers).forEach(([key, value]) => {
        message.getHeadersMap().set(key, value);
    });
    return message;
}
class NetworkClient extends baseClient_1.default {
    generateClient() {
        return network_grpc_pb_1.NetworkClient;
    }
    httpRequest(req) {
        return this.buildQuery((message, callback) => this.client.hTTPRequest(message, callback), () => {
            const msg = new network_pb_1.default.HTTPRequestMsg();
            msg.setUrl(req.url);
            msg.setMethod(req.method);
            msg.setBody(req.body);
            addHeadersToMessage(msg, req.headers);
            return msg;
        }, (response) => ({
            statusCode: response.getResponsecode(),
            data: response.getData(),
            headers: parseHeadersMap(response.getHeadersMap()),
        }));
    }
}
exports.NetworkClient = NetworkClient;
