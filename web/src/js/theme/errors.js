'use strict';

/**
 *
 */
export class ThemeSaveException extends Error {
    /**
     *
     */
    constructor() {
        super('Failed to save theme to localstorage');
    }
}
