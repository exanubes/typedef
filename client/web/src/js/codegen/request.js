'use strict';

import { INPUT_TYPES, TARGET_FORMATS } from './domain';
import {
    ExceededMaxLengthException,
    InvalidFormatException,
    InvalidInputTypeException,
    InvalidSyntaxException,
} from './errors';

export const codegen_request_factory = {
    /**
     * @param {HTMLFormElement} form
     * @returns {CodegenRequest}
     * */
    create(form) {
        const formValues = new FormData(form);

        const get_input_type = () => {
            const input_type = formValues.get('input_type').trim();

            if (!Object.keys(INPUT_TYPES).includes(input_type)) {
                throw new InvalidInputTypeException(input_type);
            }

            return input_type;
        };

        const get_input = () => {
            const input = formValues.get('structure-input').trim();
            if (input.length > 10_000) {
                throw new ExceededMaxLengthException(10_000);
            }

            try {
                JSON.parse(input);
            } catch (error) {
                throw new InvalidSyntaxException('json', error);
            }

            return input;
        };

        const get_format = () => {
            const format = formValues.get('format').trim();
            if (!Object.keys(TARGET_FORMATS).includes(format)) {
                throw new InvalidFormatException(format);
            }

            return format;
        };

        return {
            input_type: get_input_type,
            input: get_input,
            format: get_format,
        };
    },
};

/**
 * @typedef {object} CodegenRequest
 * @property {InputTypeGetter} input_type
 * @property {InputGetter} input
 * @property {FormatGetter} format
 * */

/**
 * @callback InputTypeGetter
 * @throws {InvalidInputTypeException}
 * @returns {import("./domain").InputType}
 * */

/**
 * @callback InputGetter
 * @returns {string}
 * */

/**
 * @callback FormatGetter
 * @throws {InvalidFormatException}
 * @returns {import("./domain").Format}
 * */
