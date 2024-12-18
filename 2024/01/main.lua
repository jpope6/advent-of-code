package.path = package.path .. ";../?.lua"
local util = require("util")

local function part1(lines)
    local left = {}
    local right = {}
    for _, line in ipairs(lines) do
        local l, r = string.match(line, "(%d+)%s+(%d+)")
        table.insert(left, tonumber(l))
        table.insert(right, tonumber(r))
    end

    table.sort(left)
    table.sort(right)

    local sum = 0
    for i = 1, #left do
        sum = sum + math.abs(left[i] - right[i])
    end
    return sum
end

local function part2(lines)
    local left = {}
    local right = {}
    for _, line in ipairs(lines) do
        local l, r = string.match(line, "(%d+)%s+(%d+)")
        local num = tonumber(r) or 0
        right[num] = (right[num] or 0) + 1

        table.insert(left, tonumber(l))
    end

    local sum = 0
    for i = 1, #left do
        local product = left[i] * (right[left[i]] or 0)
        sum = sum + product
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
