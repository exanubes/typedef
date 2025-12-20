/**
 * @typedef {Object} Cache
 * @property {(key: string)=>CacheItem} get
 * @property {(key: string, payload: Record<string, unknown>)=>boolean} write
 * */

/**
 * @typedef {Object} Client
 * @property {(id: string, key_path: string, table: string)=>Promise<object>} find
 * @property {(payload: object, table: string)=>Promise<void>} write
 * */


/**
 * @typedef {Object} Config
 * @property {boolean} auto_increment
 * @property {string} key_path
 * @property {Map<string, {}>} indexes
 * @property {string} table_name
 * @property {string} namespace
 * */

export { }
