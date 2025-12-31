local visual_block_mode = "\22" --- ctrl+v
local SelectionReader = {}
SelectionReader.__index = SelectionReader

---@param with_range boolean
---@return InputReader
function SelectionReader.new(with_range)
    return setmetatable({ with_range = with_range }, SelectionReader)
end

function SelectionReader:read()
    local mode = vim.fn.mode()

    if not self.with_range and mode ~= "v" and mode ~= "V" and mode ~= visual_block_mode then
        return ""
    end

    local bufnr = 0
    local start_pos = vim.fn.getpos("'<")
    local end_pos = vim.fn.getpos("'>")

    local start_row = start_pos[2] - 1
    local start_col = start_pos[3]
    local end_row = end_pos[2]
    local end_col = end_pos[3]

    local lines = vim.api.nvim_buf_get_lines(bufnr, start_row, end_row, false)

    if #lines == 0 then
        return ""
    end

    if #lines == 1 then
        return string.sub(lines[1], start_col, end_col)
    end

    lines[1] = string.sub(lines[1], start_col)
    lines[#lines] = string.sub(lines[#lines], 1, end_col)

    return table.concat(lines, "\n")
end

return SelectionReader
