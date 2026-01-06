local Panel = require("typedef.infrastructure.panel.side-panel-overlay")
local PanelService = require("typedef.app.open-side-panel")
local M = {}

function M.register(server)
    vim.api.nvim_create_user_command("TypedefPanel", function()
        local panel = Panel.new()
        local panel_service = PanelService.new(panel)
        panel_service:execute()
    end, {})
end

return M
