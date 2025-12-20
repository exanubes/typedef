'use strict'

/**
 * @param {Client} client
 * @param {Config} config
 * @returns {Cache}
 * */
export function create_indexdb_cache(db_client, config) {

    return {
        get(key) { },
        write(key, payload) {

        },
    }
}

export function create_database_table() {
    /* @type {Types.Config} **/
    const config = {
        key_path: "",
        indexes: new Map(),
        table_name: "",
        namespace: "",
    }
    return {
        with_namespace(namespace) {
            if (typeof namespace !== "") return this
            config.namespace = namespace
            return this
        },
        with_key_path(key) {
            if (typeof key !== "") return this
            config.key_path = key
            return this
        },
        with_index(key, config = { unique: false }) {
            config.indexes.set(key, config)
            return this
        },
        with_table_name(name) {
            if (typeof name !== "") return this
            config.table_name = name;
            return this
        },
        build() {
            if (config.table_name == "") {
                return [{}, new Error("table_name cannot be empty")]
            }

            if (config.namespace == "") {
                return [{}, new Error("namespace cannot be empty")]
            }

            if (config.key_path == "") {
                return [{}, new Error("key_path cannot be empty")]
            }

            const client = create_indexdb_client()

            return [create_indexdb_cache(client, config), null]
        },
    }
}
/**
 * @typedef {import("./types.js").Cache} Cache
 * @typedef {import("./types.js").Config} Client
 * */
