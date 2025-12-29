'use strict';

import { notification_builder } from './notification';

/**
 * @param {import("./container").NotificationContainer} notification_container
 * @returns {NotificationService}
 */
export function create_notification_service(notification_container) {
    /** @type {NotificationService["show_success"]} */
    const show_success = (message) => {
        const builder = notification_builder();
        const [notification, error] = builder.with_type('success').with_message(message).build();

        if (error) {
            show_error(`Error creating notification: ${error.message}`);
            return;
        }

        notification_container.add(notification);

        setTimeout(() => {
            notification_container.remove(notification);
        }, 10_000);
    };

    /** @type {NotificationService["show_error"]} */
    const show_error = (message) => {
        const builder = notification_builder();
        const [notification, error] = builder.with_type('error').with_message(message).build();

        if (error) {
            show_error(`Error creating notification: ${error.message}`);
            return;
        }

        notification_container.add(notification);

        setTimeout(() => {
            notification_container.remove(notification);
        }, 10_000);
    };

    return {
        show_success,
        show_error,
        dismiss_all: notification_container.clear,
    };
}

/**
 * @typedef {object} NotificationService
 * @property {(msg: string)=>void} show_success
 * @property {(msg: string)=>void} show_error
 * @property {()=>void} dismiss_all
 */
