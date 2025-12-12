'use strict'

import {
    create_system_media,
    create_theme_service,
    create_theme_storage,
    create_theme_ui,
    register_theme_toggle_event,
} from './js/theme'

console.log("test")
document.addEventListener('DOMContentLoaded', () => {
    const theme_service = create_theme_service(
        create_theme_storage(),
        create_system_media(),
        create_theme_ui(),
    )

    theme_service.start()
    register_theme_toggle_event(theme_service)
});
