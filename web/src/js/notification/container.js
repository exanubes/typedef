'use strict';
const CONTAINER_ID = 'notification-container';

/**
 * @returns {NotificationContainer}
 * */
export function create_notification_container() {
    const container = get_container(CONTAINER_ID);
    /**
     * @type {Set<import("./notification").Notification>}
     * */
    const notifications = new Set();

    // TODO: Display a "clear" button when 3 or more notifications are in the container

    return Object.freeze({
        add(notification) {
            const toast = notification.html();
            container.appendChild(toast);

            // Trigger reflow to ensure animation plays
            void toast.offsetHeight;

            toast.classList.add('toast--visible');
            notifications.add(notification);
        },
        remove(notification) {
            if (notifications.has(notification)) {
                notification.close();
                notifications.delete(notification);
            }
        },
        clear() {
            for (const notification of notifications) {
                notification.close();
                notifications.delete(notification);
            }
        },
    });
}

/**
 * @param id
 * @returns {HTMLElement}
 */
const get_container = (id) => {
    let container = document.getElementById(id);
    if (!container) {
        container = document.createElement('div');
        container.id = CONTAINER_ID;
        container.className = 'toast-container';
        container.setAttribute('aria-live', 'polite');
        container.setAttribute('aria-atomic', 'false');
        document.body.appendChild(container);
    }

    return container;
};

/**
 * @typedef {object} NotificationContainer
 * @property {(notification: import("./notification").Notification)=>void} add
 * @property {(notification: import("./notification").Notification)=>void} remove
 * @property {()=>void} clear
 * */
