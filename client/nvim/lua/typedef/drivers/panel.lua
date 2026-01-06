local Panel = require("typedef.infrastructure.panel.side-panel-overlay")
local PanelService = require("typedef.app.open-side-panel")
local YankReader = require("typedef.infrastructure.reader.yank")
local ChainReader = require("typedef.infrastructure.reader.chain")
local Codegen = require("typedef.infrastructure.codegen")
local InsertWriter = require("typedef.infrastructure.writer.insert")
local BufferContextWriter = require("typedef.infrastructure.writer.buffer-context")
local M = {}

---@param server JsonRpcServer
function M.register(server)
    vim.api.nvim_create_user_command("TypedefPanel", function()
        local original_buffer = vim.api.nvim_get_current_buf()
        local original_window = vim.api.nvim_get_current_win()

        local input_reader = ChainReader.new({
            YankReader.new(),
        })

        server:start()
        local codegen_repository = Codegen.new(server)
        local panel = Panel.new()

        local base_writer = InsertWriter.new()
        local output_writer = BufferContextWriter.new(original_buffer, original_window, base_writer)

        local panel_service = PanelService.new(panel, input_reader, codegen_repository, output_writer)
        panel_service:execute()
    end, {})
end

return M
