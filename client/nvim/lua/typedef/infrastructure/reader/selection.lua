local SelectionReader = {}
SelectionReader.__index = SelectionReader

---@return InputReader
function SelectionReader.new()
    return setmetatable({}, SelectionReader)
end

function SelectionReader:read()
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
