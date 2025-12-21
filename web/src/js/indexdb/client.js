'use strict'

/**
 * @param {()=>Promise<IDBDatabase>} database
 * @returns {Client}
 * */
export function create_indexdb_client(database) {
    const client = {
        async find(id, key_path, table) {
            const db = await database()
            return new Promise((resolve, reject) => {
                const tx = db.transaction(table, "readonly")
                const store = tx.objectStore(table)
                const request = store.get(id)

                request.onerror = () => reject(request.error)
                tx.oncomplete = () => resolve(request.result)
                tx.onerror = () => reject(tx.error)
                tx.onabout = () => reject(tx.error)

            })
        },
        async write(payload, table) {
            const db = await database()
            return new Promise((resolve, reject) => {
                const tx = db.transaction(table, "readwrite")
                const store = tx.objectStore(table)
                const request = store.put(payload)
                request.onerror = () => reject(request.error)
                tx.oncomplete = () => resolve()
                tx.onerror = () => reject(tx.error)
                tx.onabout = () => reject(tx.error)
            })
        },
    }

    return client
}

/**
 * @param {ManagerConfig} config
 * */
export function create_connection_manager(config) {
    /**
     * @type {IDBDatabase}
     * */
    let db = null;
    /**
     * @type {Promise<IDBDatabase>}
     * */
    let opening = null;

    return function connect() {
        if (db) return Promise.resolve(db);
        if (opening) return opening;

        opening = new Promise((resolve, reject) => {
            const request = indexedDB.open(config.namespace, config.version);

            request.onupgradeneeded = (event) => {
                config.schema(request.result, event.oldVersion)
            };

            request.onsuccess = () => {
                db = request.result;
                opening = null;
                resolve(db);
            };

            request.onerror = () => {
                opening = null;
                reject(request.error);
            };

            request.onblocked = () => {
                console.warn("IndexedDB upgrade blocked by another tab");
            };
        });

        return opening;
    };
}

/**
 * @typedef {import("./types.js").Client} Client
 * */

/**
 * @typedef {Object} ManagerConfig
 * @property {string} namespace
 * @property {number} version
 * @property {(database: IDBDatabase, version: number)=>void} schema
 * */
