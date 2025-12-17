'use strict'

import { POST } from "../http/http"

/**
 * @param {CodegenRequest} request
 * @returns {Promise<import("../http/http").HttpResponse<CodegenResponse>>}
 * */
export async function generateCode(request) {
    return POST("codegen", {
        "input_type": request.input_type,
        "data": request.data,
        "format": request.format,
    })
}

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
