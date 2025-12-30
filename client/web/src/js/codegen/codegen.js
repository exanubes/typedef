'use strict';

import {
    ExceededMaxLengthException,
    InvalidFormatException,
    InvalidInputTypeException,
    InvalidSyntaxException,
} from './errors';

/**
 *
 * @param {import("./commands").CodegenCommandHandler} codegen_handler
 * @returns {CodegenService}
 * */
export function create_codegen_service(codegen_handler) {
    /**
     * @type {Execute}
     * */
    const execute = async (request) => {
        try {
            const response = await codegen_handler.send({
                input_type: request.input_type(),
                input: request.input(),
                format: request.format(),
            });

            if (response.status === 'error') {
                return [{ code: '', format: -1 }, new CodegenError(response.message)];
            }

            return [response.data, null];
        } catch (error) {
            const exceptions = [
                InvalidInputTypeException,
                InvalidFormatException,
                ExceededMaxLengthException,
                InvalidSyntaxException,
            ];

            if (exceptions.some((exception) => error instanceof exception)) {
                return [{ code: '', format: -1 }, error];
            }

            return [{ code: '', format: -1 }, new CodegenError('Unhandled exception', error)];
        }
    };

    return {
        execute,
    };
}

/**
 * @typedef {object} CodegenService
 * @property {Execute} execute
 * */

/**
 * @callback Execute
 * @param {import("./request").CodegenRequest} request
 * @returns {Promise<[import("./api").CodegenResponse, Error]>}
 * @throws {CodegenError}
 * */

/**
 *
 */
class CodegenError extends Error {
    /**
     *
     * @param msg
     * @param cause
     */
    constructor(msg, cause) {
        super(msg, { cause });
    }
}
