'use strict'

import {
    create_system_media,
    create_theme_service,
    create_theme_storage,
    create_theme_ui,
    register_theme_toggle_event,
} from './js/theme'

import {
    create_codegen_service,
    codegen_request_factory,
} from './js/codegen'
import { create_codegen_driver } from './js/drivers/codegen';
import { create_clipboard_driver } from './js/drivers/clipboard';
import { create_clipboard } from './js/clipboard/clipboard';

document.addEventListener('DOMContentLoaded', () => {
    const theme_driver = create_theme_service(// TODO: refactor to fit the driver convention
        create_theme_storage('typedef-theme-preference'),
        create_system_media(),
        create_theme_ui(),
    )
    const codegen_driver = create_codegen_driver(
        create_codegen_service(),
        codegen_request_factory,
    )

    const clipboard_driver = create_clipboard_driver(
        create_clipboard(navigator.clipboard)
    )
    register_theme_toggle_event(theme_driver)
    theme_driver.start()

    clipboard_driver.start(document.getElementById("copy-btn"))



    codegen_driver.start(document.querySelector("form"), document.getElementById("transform-btn"))
});


