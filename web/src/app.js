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

    register_theme_toggle_event(theme_driver)
    theme_driver.start()


    codegen_driver.start(document.querySelector("form"), document.getElementById("transform-btn"))
});


