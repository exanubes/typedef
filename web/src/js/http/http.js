'use strict'

/**
 * @template T
 * @typedef {SuccessResponse<T> | ErrorResponse<T>} HttpResponse<T>
 * */

/**
 * @template T
 * @typedef {object} SuccessResponse<T>
 * @property {"ok"} status
 * @property {T} data
 * */

/**
 * @template T
 * @typedef {object} ErrorResponse<T>
 * @property {"error"} status
 * @property {string} message
 * */

/**
 * @param {string} endpoint
 * @param {Record<string, unknown>} body
 * @returns {HttpResponse<T>}
 * */
export async function POST(endpoint, body) {
    try {
        const response = await fetch(`https://aacg5m35g2.execute-api.eu-central-1.amazonaws.com/local/${endpoint}`, {
            method: "POST",
            body: JSON.stringify(body),
            headers: {
                "Content-Type": "application/json"
            }
        })

        return {
            status: "ok",
            data: await response.json()
        }
    } catch (error) {
        return {
            status: "error",
            message: error.message ?? error
        }
    }
}
