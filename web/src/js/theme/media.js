'use strict'

import { THEMES } from "./theme"


/**
 * @returns {SystemMedia}
 * */
export function create_system_media() {
    const media = window.matchMedia('(prefers-color-scheme: dark)')

    /**
     * @type {Current}
     * */
    const current = () => {
        return media.matches ? THEMES.DARK : THEMES.LIGHT
    }

    /**
     * @type {OnChange}
     * */
    const on_change = (handler) => {
        media.addEventListener("change", event => {
            handler(event.matches ? THEMES.DARK : THEMES.LIGHT)
        })
    }

    return {
        current, on_change
    }
}


/**
 * @typedef {import("./theme").Theme} Theme
 * @typedef {Object} SystemMedia
 * @property {Current} current
 * @property {OnChange} on_change
 * */

/**
 * @callback ThemeChangeHandler
 * @param {Theme} theme
 * @returns {void}
 * */

/**
 * @callback Current
 * @returns {Theme}
 * */

/**
 * @callback OnChange
 * @param {ThemeChangeHandler} handler
 * @returns {void}
 * */
