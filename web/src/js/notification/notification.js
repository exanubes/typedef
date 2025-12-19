'use strict'

/**
 * @returns {NotificationBuilder}
 * */
export function notification_builder() {
    const config = {
        type: '',
        msg: '',
    }

    /**
     * @type {NotificationBuilder}
     * */
    return {
        with_type(type) {
            config.type = type

            return this
        },

        with_message(msg) {
            config.msg = msg

            return this
        },

        build() {
            if (config.type !== "error" && config.type !== "success") {
                return [null_notification, new Error(`Expected type to be "success" or "error", received: ${config.type}`)]
            }

            if (config.msg === "") {
                return [null_notification, new Error(`Notification cannot have an empty message`)]
            }

            return [create_notification(config), null]
        }
    }
}

/**
 * @param {{msg: string; type: string}} config
 * */
function create_notification(config) {
    let is_active = true
    const toast = document.createElement('div')
    toast.className = `toast toast--${config.type}`
    toast.setAttribute('role', 'alert')

    const icon = config.type === 'success' ? create_success_icon() : create_error_icon()

    toast.innerHTML = `
            ${icon}
            <span class="toast__message">${config.msg}</span>
            <button class="toast__close" aria-label="Close notification">
                ${create_close_icon()}
            </button>
        `


    const close_button = toast.querySelector('.toast__close')

    const handle_close = () => {
        if (!is_active) {
            console.warn("Notification is already closed!: '%s'", config.msg)
            return;
        }
        toast.classList.remove('toast--visible')
        toast.classList.add('toast--removing')
        toast.addEventListener('animationend', () => {
            toast.remove()
            is_active = false
        }, { once: true })
    }

    if (close_button) {
        close_button.addEventListener('click', () => {
            close_button.setAttribute("disabled", "true")
            handle_close()
        })
    }

    return {
        html() {
            return toast
        },
        value() {
            return toast.innerHTML
        },
        active() {
            return is_active
        },
        close: handle_close
    }
}

/**
 * @returns {string}
 */
const create_success_icon = () => `
        <svg class="toast__icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
            <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
            <polyline points="22 4 12 14.01 9 11.01"></polyline>
        </svg>
    `

/**
 * @returns {string}
 */
const create_error_icon = () => `
        <svg class="toast__icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
            <circle cx="12" cy="12" r="10"></circle>
            <line x1="15" y1="9" x2="9" y2="15"></line>
            <line x1="9" y1="9" x2="15" y2="15"></line>
        </svg>
    `

/**
 * @returns {string}
 */
const create_close_icon = () => `
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
        </svg>
    `

/**
 * @type {Notification}
 * */
const null_notification = { html: () => document.createElement('div'), value: () => "", is_active: () => false, close() { } }

/**
 * @typedef {Object} Notification
 * @property {()=>HTMLElement} html
 * @property {()=>string} value
 * @property {()=>boolean} active
 * @property {()=>void} close
 * */

/**
 * @typedef {Object} NotificationBuilder
 * @property {(type: "error" | "success")=>NotificationBuilder} with_type
 * @property {(msg: string)=>NotificationBuilder} with_message
 * @property {()=>[ Notification, Error ]} build
 * */
