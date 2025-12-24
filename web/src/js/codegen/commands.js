'use strict'

import { validate_json_rpc_response } from "../libs/jsonrpc"

/**
 * @typedef {Object} CodegenResponse
 * @property {string} code
 * @property {number} format
 * */

/**
 * @typedef {Object} CodegenRequest
 * @property {import("../codegen/domain").InputType} input_type
 * @property {string} data
 * @property {import("../codegen/domain").Format} format
 * */

/**
 * @typedef {Object} CodegenCommandHandler
 * @property {(request: CodegenRequest)=>Promise<HandlerResponse>} send
 * */


/**
 * @typedef {SuccessResponse |ErrorResponse} HandlerResponse
 * */

/**
 * @typedef {Object} SuccessResponse
 * @property {"ok"} status
 * @property {CodegenResponse } data
 * */

/**
 * @typedef {Object} ErrorResponse
 * @property {"error"} status
 * @property {Error } err
 * */

/**
 * @typedef {import("../libs/jsonrpc").JSONRPCRequest<CodegenRequest>} CodegenJSONRPCRequest
 * */


/**
 * @param {object} response
 * @returns {CodegenResponse}
 * */
export function parse_codegen_response(response) {
    validate_json_rpc_response(response)
    return {
        code: response.result.code,
        format: response.result.format,
    }
}


