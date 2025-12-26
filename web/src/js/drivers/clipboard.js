'use strict';

/**
 * @param {import('../clipboard/clipboard').ClipboardService} clipboard
 * @param {import('../notification/service').NotificationService} notifications
 * @returns {Driver}
 * */
export function create_clipboard_driver(clipboard, notifications) {
    /**
     * @type {Start}
     * */
    const start = (button) => {
        button.addEventListener('click', async () => {
            const output = document.getElementById('output-code');
            const code = output.innerText;

            clipboard.save(code);
            notifications.show_success('Saved to clipboard!');
        });
    };

    return {
        start,
    };
}

/**
 * @callback Start
 * @param {HTMLButtonElement} button
 * */

/**
 * @typedef Driver
 * @property {Start} start
 * */
