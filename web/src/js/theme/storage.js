'use strict'

import { ThemeSaveException } from "./errors"


/**
 * @typedef {import("./theme").Theme} Theme
 * @typedef {Object} ThemeStorage
 * @property {()=>void} load
 * @property {(theme: Theme)=>void} save
 * */

/**
 * @returns {ThemeStorage}
 * */
export function create_theme_storage() {
    const load = () => {
        try {
            const value = localStorage.getItem(storageKey);
            return isValidTheme(value) ? value : null;
        } catch {
            return null;
        }
    }

    /**
     * @param {Theme} theme
     * */
    const save = (theme) => {
        try {
            localStorage.setItem(storageKey, theme);
        } catch {
            new ThemeSaveException()
        }
    }

    return {
        load, save
    }
}
