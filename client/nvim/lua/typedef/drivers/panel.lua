local Panel = require("typedef.infrastructure.panel.side-panel-overlay")
local PanelService = require("typedef.app.open-side-panel")
local YankReader = require("typedef.infrastructure.reader.yank")
local ChainReader = require("typedef.infrastructure.reader.chain")
local Codegen = require("typedef.infrastructure.codegen")
local InsertWriter = require("typedef.infrastructure.writer.insert")
local M = {}

---@param server JsonRpcServer
function M.register(server)
    vim.api.nvim_create_user_command("TypedefPanel", function()
        local input_reader = ChainReader.new({
            YankReader.new(),
        })

        server:start()
        local codegen_repository = Codegen.new(server)
        local panel = Panel.new()
        local output_writer = InsertWriter.new()
        local panel_service = PanelService.new(panel, input_reader, codegen_repository, output_writer)
        panel_service:execute()
    end, {})
end

return M
