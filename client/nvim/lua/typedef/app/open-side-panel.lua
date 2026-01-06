local Map = require("typedef.helpers.map")
local M = {
    --- @type Panel
    panel = nil,
}
M.__index = M

---@param panel Panel
function M.new(panel)
    return setmetatable({
        panel = panel,
    }, M)
end

function M:execute()
    local keymap = Map.new()
    keymap:set("1", 1)
    local lines = {
        "Select input format",
        "(*) json",
        "",
        "Select output format",
        "(*) Go",
        "( ) Typescript",
        "( ) Zod",
        "( ) JSDoc",
        "",
        "[Insert]",
        "[Copy to clipboard]",
        "[Exit]",
    }

    self.panel:add_keymap("q", function()
        self.panel:close()
    end)

    self.panel:add_keymap("<CR>", function(event)
        vim.notify("Row selected: " .. lines[event.current_line])
    end)

    self.panel:render(lines)
end

return M
