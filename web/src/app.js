'use strict'

import {
    create_system_media,
    create_theme_service,
    create_theme_storage,
    create_theme_ui, register_theme_toggle_event,
} from './js/theme'
import {
    create_codegen_service,
    codegen_request_factory,
} from './js/codegen'
import { create_codegen_driver } from './js/drivers/codegen';
import { create_clipboard_driver } from './js/drivers/clipboard';
import { create_clipboard } from './js/clipboard/clipboard';
import { create_notification_container, create_notification_service } from './js/notification';
import { create_codegen_input_repository } from './js/indexdb/repositories/codegen-input';
import { create_database } from './js/indexdb/factory';
import { create_hasher } from './js/hasher/hasher';
import { serialize } from './js/libs/canonicalize';
import { create_request_cache } from './js/cache/request-cache';
import { create_rpc_client } from './js/rpc/client';
import { create_rpc_codegen_command_handler } from './js/codegen/command-handler';

document.addEventListener('DOMContentLoaded', () => {
    const database = create_database()
    const codegen_repository = create_codegen_input_repository(database)
    const hashing_service = create_hasher(serialize)
    const cache_service = create_request_cache(codegen_repository, hashing_service)

    const theme_driver = create_theme_service(// TODO: refactor to fit the driver convention
        create_theme_storage('typedef-theme-preference'),
        create_system_media(),
        create_theme_ui(),
    )
    const notification_service = create_notification_service(
        create_notification_container()
    )
    const rpc_client = create_rpc_client()
    const codegen_handler = create_rpc_codegen_command_handler(rpc_client)

    const codegen_driver = create_codegen_driver(
        create_codegen_service(cache_service, hashing_service, codegen_handler),
        codegen_request_factory,
        notification_service,
    )

    const clipboard_driver = create_clipboard_driver(
        create_clipboard(navigator.clipboard),
        notification_service,
    )

    register_theme_toggle_event(theme_driver)
    theme_driver.start()

    clipboard_driver.start(document.getElementById("copy-btn"))

    codegen_driver.start(document.querySelector("form"), document.getElementById("transform-btn"))
});


