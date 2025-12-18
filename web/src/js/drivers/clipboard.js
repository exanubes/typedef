'use strict'

/**
 * @param {import('../clipboard/clipboard').ClipboardService} clipboard
 * @returns {Driver}
 * */
export function create_clipboard_driver(clipboard) {

    /**
     * @type {Start}
     * */
    const start = (button) => {
        button.addEventListener('click', async () => {
            const output = document.getElementById("output-code")
            const code = output.innerText

            clipboard.save(code)
        })

    }

    return {
        start
    }
}

/**
 * @callback Start
 * @param {HTMLButtonElement} button
 * */

/**
 * @typedef Driver
 * @property {Start} start
 * */
