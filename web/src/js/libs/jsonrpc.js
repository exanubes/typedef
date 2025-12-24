'use strict'

let id = 1

/**
 * @template T
 * @param {string} method
 * @param {T} body
 * @returns {JSONRPCRequest<T>}
 * */
export function create_json_rpc_request(method, body) {
    return {
        method,
        jsonrpc: "2.0",
        params: body,
        id: id++,
    }
}


export function validate_json_rpc_response(response) {
    if (!('jsonrpc' in response)) {
        throw new Error("Missing required param \"jsonrpc\"")
    }

    if (typeof response.jsonrpc !== "string") {
        throw new Error(`Invalid jsonrpc type. Expected string, received ${typeof response.jsonrpc}`)
    }

    if (response.jsonrpc !== "2.0") {
        throw new Error(`Invalid rpc version. Expected 2.0, received ${response.jsonrpc}`)
    }

    if (typeof response.id !== "number") {
        throw new Error(`Invalid id type. Expected number, received ${typeof response.jsonrpc}`)
    }

    if (response.error) {
        if (typeof response.error.code !== "number") {
            throw new Error(`Invalid error code type. Expected number, received ${typeof response.jsonrpc}`)
        }
    }
}





/**
 * @template T
 * @typedef {Object} JSONRPCRequest
 * @property {"2.0"} jsonrpc
 * @property {string} method
 * @property {T} params
 * @property {number} id
 * */

/**
 * @template T
 * @typedef {Object} JSONRPCResponse
 * @property {"2.0"} jsonrpc
 * @property {T}  [result]
 * @property {number} [id]
 * @property {JSONRPCError} [error]
 * */

/**
 * @typedef {Object} JSONRPCError
 * @property {number} code
 * @property {string} [message]
 * @property {object} [data]
 * */

