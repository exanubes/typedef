/**
 * @param {import("../codegen/codegen").CodegenService} codegen_service
 * @param {typeof import("../codegen/request").codegen_request_factory} request_factory
 * @param {import("../notification/notification").NotificationService} notification_service
 * @returns {CodegenDriver}
 * */
export function create_codegen_driver(codegen_service, request_factory, notification_service) {
    /**
     * @type {Start}
     * */
    const start = (form, button) => {
        form.addEventListener('submit', async (event) => {
            event.preventDefault()

            button.setAttribute("disabled", "true")

            const [data, error] = await codegen_service.execute(request_factory.create(event.target))
            if (error) {
                notification_service.show_error(error.message || 'Failed to generate code')
                button.removeAttribute("disabled")
                return
            }

            const output_container = document.getElementById("output-code")
            output_container.innerText = data.code
            notification_service.show_success('Code generated successfully!')
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

