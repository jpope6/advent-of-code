package.path = package.path .. ";../?.lua"
local util = require("util")

local function part1(lines)
    local line = lines[1]
    local sum = 0
    for num1, num2 in string.gmatch(line, "mul%((%d+),(%d+)%)") do
        sum = sum + (tonumber(num1) * tonumber(num2))
    end
    return sum
end

local function part2(lines)
    local line = lines[1]
    local mul_pattern = "mul%((%d+),(%d+)%)"
    local do_pattern = "do%(%s*%)"
    local dont_pattern = "don't%(%s*%)"

    local index = 1
    local sum = 0
    local mul_enabled = true -- Start with `mul` enabled

    while index <= #line do
        local start_mul, end_mul, num1, num2 = string.find(line, mul_pattern, index)
        local start_do, end_do = string.find(line, do_pattern, index)
        local start_dont, end_dont = string.find(line, dont_pattern, index)

        local start_values = {}
        if start_mul then table.insert(start_values, start_mul) end
        if start_do then table.insert(start_values, start_do) end
        if start_dont then table.insert(start_values, start_dont) end
        if #start_values == 0 then
            break
        end
        table.sort(start_values)

        local smallest_start = start_values[1]

        -- Case when the smallest start is mul and mul_enabled is true
        if smallest_start == start_mul and mul_enabled then
            sum = sum + (tonumber(num1) * tonumber(num2))
            index = end_mul + 1
            -- Case when the smallest start is do
        elseif smallest_start == start_do then
            mul_enabled = true
            index = end_do + 1
            -- Case when the smallest start is dont
        elseif smallest_start == start_dont then
            mul_enabled = false
            index = end_dont + 1
        else
            index = index + 1
        end
    end

    return sum
end

local function main()
    local lines = util.read_lines("input.txt")

    local result, elapsedTime = util.run_function_with_timer(part2, lines)
    print(string.format("\nResult: %d", result))
    print(string.format("\nElapsed Time: %.6f seconds\n", elapsedTime))
end

main()
