/**
 * @param {import('./service').ThemeService}  service
 * */
export function register_theme_toggle_event(service) {
    const button = document.getElementById("theme-toggle")
    if (!button) return;

    button.addEventListener("click", () => {
        service.toggle()
    })
}
