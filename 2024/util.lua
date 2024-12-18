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

function util.run_function_with_timer(func, ...)
    local startTime = os.clock()
    local result = func(...)
    local endTime = os.clock()
    local elapsedTime = endTime - startTime
    return result, elapsedTime
end

return util
