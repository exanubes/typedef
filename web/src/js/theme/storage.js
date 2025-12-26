'use strict';

import { ThemeSaveException } from './errors';
import { is_valid_theme } from './theme';

/**
 * @typedef {import("./theme").Theme} Theme
 * @typedef {object} ThemeStorage
 * @property {()=>Theme | null} load
 * @property {(theme: Theme)=>void} save
 * */

/**
 * @param {string} storage_key
 * @returns {ThemeStorage}
 * */
export function create_theme_storage(storage_key) {
    /**
     * @returns {Theme | null}
     */
    const load = () => {
        try {
            const value = localStorage.getItem(storage_key);
            return is_valid_theme(value) ? value : null;
        } catch {
            return null;
        }
    };

    /**
     * @param {Theme} theme
     * */
    const save = (theme) => {
        try {
            localStorage.setItem(storage_key, theme);
        } catch {
            new ThemeSaveException();
        }
    };

    return {
        load,
        save,
    };
}
