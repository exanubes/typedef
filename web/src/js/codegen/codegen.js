'use strict'

import { generateCode } from "./api"
import { ExceededMaxLengthException, InvalidFormatException, InvalidInputTypeException } from "./errors"

/**
 * @param {import("../cache/repositories/codegen").CodegenInputRepository} cache_service
 * @returns {CodegenService}
 * */
export function create_codegen_service(cache_service, hash_service) {
    /**
     * @type {Execute}
     * */
    const execute = async (request) => {
        try {
            const input_hash = hash_service.hash(request.input())
            const result = await cache_service.find(input_hash)
            if (result) {
                return [{
                    // TODO: replace with output
                    code: result.output_hash,
                    format: result.target
                }, null]
            }
            const response = await generateCode({
                input_type: request.input_type(),
                data: request.input(),
                format: request.format(),
            })

            if (response.status === "error") {
                return [{ code: "", format: -1 }, new CodegenError(response.message)]
            }
            // TODO::
            // - Deduplication
            // - Caching

            return [response.data, null]
        } catch (error) {
            const exceptions = [InvalidInputTypeException, InvalidFormatException, ExceededMaxLengthException]

            if (exceptions.some(exception => error instanceof exception)) {
                return [{ code: "", format: -1 }, new CodegenError("Invalid input", error)]
            }

            return [{ code: "", format: -1 }, new CodegenError("Unhandled exception", error)]
        }
    }

    return {
        execute
    }
}

/**
 * @typedef {Object} CodegenService
 * @property {Execute} execute
 * */

/**
 * @callback Execute
 * @param {import("./request").CodegenRequest} request
 * @returns {Promise<[import("./api").CodegenResponse, Error]>}
 * @throws {CodegenError}
 * */

class CodegenError extends Error {
    constructor(msg, cause) {
        super(msg, { cause })
    }
}


