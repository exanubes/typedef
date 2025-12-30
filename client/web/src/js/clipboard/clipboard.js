/**
 * @param {ClipboardClient} client
 * @returns {ClipboardService}
 * */
export function create_clipboard(client) {
    /**
     * @type {Save}
     * */
    const save = (text) => {
        // TODO:indicate success/failure of the copy to clipboard action in the UI
        return client.writeText(text);
    };

    return {
        save,
    };
}

/**
 * @typedef {object} ClipboardClient
 * @property {(string)=>Promise<void>} writeText
 * */

/**
 * @callback Save
 * @param {string} text
 * @returns {Promise<void>}
 * */

/**
 * @typedef ClipboardService
 * @property {Save} save
 * */
