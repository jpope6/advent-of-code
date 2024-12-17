local util = {}

function util.read_lines(filename)
    local lines = {}
    local file = io.open(filename, "r")
    if not file then
        error("Could not open file: " .. filename)
    end

    for line in file:lines() do
        table.insert(lines, line)
    end

    file:close()
    return lines
end

return util
