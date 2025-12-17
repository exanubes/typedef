'use strict'

import { generateCode } from "./api"
import { ExceededMaxLengthException, InvalidFormatException, InvalidInputTypeException } from "./errors"

/**
 * @param {import('./input_resolver').CodegenRequest} input_resolver
 * @returns {CodegenService}
 * */
export function create_codegen_service() {
    /**
     * @type {Execute}
     * */
    const execute = async (request) => {
        try {
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
 * @returns {Promise<[import("../api/codegen").CodegenResponse, Error]>}
 * @throws {CodegenError}
 * */

class CodegenError extends Error {
    constructor(msg, cause) {
        super(msg, { cause })
    }
}


