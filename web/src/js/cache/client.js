'use strict'

/**
 * @param {Config} config
 * @returns {Client}
 * */
function create_indexdb_client(config) {
    const connect = create_connection_manager(config)
    const client = {
        async find(id) {
            const db = await connect()
            return new Promise((resolve, reject) => {
                const tx = db.transaction(config.table_name, "readonly")
                const store = tx.objectStore(config.table_name)
                const request = store.get({ [config.key_path]: id })

                request.onerror = () => reject(request.error)
                tx.oncomplete = () => resolve(request.result)
                tx.onerror = () => reject(tx.error)
                tx.onabout = () => reject(tx.error)

            })
        },
        async write(payload) {
            const db = await connect()
            return new Promise((resolve, reject) => {
                const tx = db.transaction(config.table_name, "readwrite")
                const store = tx.objectStore(config.table_name)
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
 * @param {Config} config
 * */
function create_connection_manager(config) {
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
                const database = request.result;
                if (event.oldVersion < 1) {
                    const store = database.createObjectStore(config.table_name, {
                        keyPath: config.key_path,
                    });

                    for (const [key, indexConfig] of config.indexes) {
                        store.createIndex(key, key, indexConfig);
                    }
                }
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
 * @typedef {import("./types.js").Config} Config
 * @typedef {import("./types.js").Client} Client
 * */
