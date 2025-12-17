/**
 * @param {import("../codegen/codegen").CodegenService} codegen_service
 * @param {typeof import("../codegen/request").codegen_request_factory} request_factory
 * @returns {CodegenDriver}
 * */
export function create_codegen_driver(codegen_service, request_factory) {
    /**
     * @type {Start}
     * */
    const start = (form, button) => {
        form.addEventListener('submit', async (event) => {
            event.preventDefault()

            button.setAttribute("disabled", "true")

            const [data, error] = await codegen_service.execute(request_factory.create(event.target))
            if (error) {
                console.error(error)
                // TODO:render error
                button.removeAttribute("disabled")
                return
            }

            const output_container = document.getElementById("output-code")
            output_container.innerText = data.code
            button.removeAttribute("disabled")
        })
    }

    return {
        start
    }
}

/**
 * @typedef {Object} CodegenDriver
 * @property {Start} start
 * */

/**
 * @callback Start
 * @param {HTMLFormElement} form
 * @param {HTMLButtonElement} button
 * @returns { void }
 * */

