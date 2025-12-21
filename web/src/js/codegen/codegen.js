'use strict'

import { generateCode } from "./api"
import { ExceededMaxLengthException, InvalidFormatException, InvalidInputTypeException } from "./errors"

/**
 * 
 * @param {import("../cache/request-cache").RequestCache} cache_service // TODO: replace with cache service
 * @param {import("../hasher/hasher").HasherService} hash_service
 * @returns {CodegenService}
 * */
export function create_codegen_service(cache_service, hash_service) {
    /**
     * @type {Execute}
     * */
    const execute = async (request) => {
        try {
            const cached_item = await cache_service.get(request.input())
            if (cached_item) {
                return [{
                    code: cached_item.output,
                    format: cached_item.target
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

            await cache_service.put(request.input(), request.format(), response.data.code)

            return [response.data, null]
        } catch (error) {
            const exceptions = [InvalidInputTypeException, InvalidFormatException, ExceededMaxLengthException]

            if (exceptions.some(exception => error instanceof exception)) {
                return [{ code: "", format: -1 }, new CodegenError("Invalid input", error)]
            }
            console.error("[CodegenService] " + error)

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


