'use strict';

export const THEMES = Object.freeze({
    LIGHT: 'light',
    DARK: 'dark',
});

/**
 * @param value
 * @returns {boolean}
 * */
export function is_valid_theme(value) {
    return value === THEMES.LIGHT || value === THEMES.DARK;
}

/**
 * @param theme
 * @returns {THEMES[keyof typeof THEMES]}
 * */
export function opposite_theme(theme) {
    return theme === THEMES.DARK ? THEMES.LIGHT : THEMES.DARK;
}

/**
 * @typedef {THEMES[keyof typeof THEMES]} Theme
 * */
