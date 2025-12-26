'use strict';

import { THEMES, opposite_theme } from './theme';

/**
 * @returns {ThemeUi}
 * */
export function create_theme_ui() {
    /**
     * @type {HTMLElement}
     * */
    const root = document.documentElement;

    /**
     * @type {Apply}
     * */
    const apply = (theme, with_transition) => {
        if (with_transition) {
            root.classList.add('theme-transition');
        }

        root.setAttribute('data-theme', theme);

        if (with_transition) {
            setTimeout(() => root.classList.remove('theme-transition'), 300);
        }
    };

    /**
     * @type {UpdateToggle}
     * */
    const update_toggle = (theme) => {
        const button = document.getElementById('theme-toggle');
        if (!button) return;

        const next = opposite_theme(theme);

        button.setAttribute('aria-label', `Switch to ${next} theme`);
        button.setAttribute('aria-pressed', String(theme === THEMES.DARK));
    };

    return {
        apply,
        update_toggle,
    };
}

/**
 * @typedef {import("./theme").Theme} Theme
 * @typedef {object} ThemeUi
 * @property {Apply} apply
 * @property {UpdateToggle} update_toggle
 * */

/**
 * @callback Apply
 * @param {Theme} theme
 * @param {boolean} with_transition
 * @returns {void}
 * */

/**
 * @callback UpdateToggle
 * @param {Theme} theme
 * @returns {void}
 * */
