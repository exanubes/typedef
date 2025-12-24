'use strict'

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
 * @typedef {import("../libs/jsonrpc").JSONRPCRequest<CodegenRequest>} CodegenJSONRPCRequest
 * */


export { }

