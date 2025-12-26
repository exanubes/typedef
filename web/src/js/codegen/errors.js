import { INPUT_TYPES, TARGET_FORMATS } from './domain';

/**
 *
 */
export class InvalidFormatException extends Error {
    /**
     *
     * @param format
     */
    constructor(format) {
        const formats = Object.keys(TARGET_FORMATS).join(', ');
        super(`Expected one of: ${formats}, received: '${format}'`);
    }
}

/**
 *
 */
export class InvalidInputTypeException extends Error {
    /**
     *
     * @param input_type
     */
    constructor(input_type) {
        const types = Object.keys(INPUT_TYPES).join('|');
        super(`Expected one of: ${types}, received: '${input_type}'`);
    }
}
/**
 *
 */
export class ExceededMaxLengthException extends Error {
    /**
     *
     * @param max_length
     */
    constructor(max_length) {
        super(`Input value cannot exceed ${max_length}`);
    }
}

/**
 *
 */
export class InvalidSyntaxException extends Error {
    /**
     *
     * @param format
     * @param cause
     */
    constructor(format, cause) {
        super(`Invalid ${format} syntax`, { cause });
    }
}
