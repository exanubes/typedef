'use strict'

import { opposite_theme } from './theme'

/**
     * @param {import('./storage').ThemeStorage} storage
     * @param {import('./media').SystemMedia} system
     * @param {import('./ui').ThemeUi} ui
     * @returns {ThemeService}
     * */
export function create_theme_service(storage, system, ui) {
    let current_theme = ""

    const start = () => {
        current_theme = storage.load() || system.current()

        ui.apply(current_theme, false)
        ui.update_toggle(current_theme)

        system.on_change(theme => {
            if (!storage.load()) {
                current_theme = theme
                ui.apply(theme, true)
                ui.update_toggle(theme)
            }
        })
    }

    const toggle = () => {
        current_theme = opposite_theme(current_theme)
        storage.save(current_theme)
        ui.apply(current_theme, true)
        ui.update_toggle(current_theme)
    }

    return {
        start, toggle
    }
}

/**
 * @typedef  {Object} ThemeService
 * @property {()=>void} start
 * @property {()=>void} toggle
 * */

