local serialize = require "@std.serialize"
local sequence = require "@std.sequence"
local entries = require "@std.entries"
local dir = require "@std.dir"
local path = require "@std.path"

local function pass_paths(func)
  return function(file)
    local destination = path(file[1])
    local source = path(file[2])
    return func(source, destination)
  end
end

local function remove_file(_, destination)
  return "rm -rf " .. destination
end

local function copy_file(source, destination)
  return table.concat({
    "# Copying " .. source .. " to " .. destination,
    remove_file(nil, destination),
    "mkdir -p " .. dir(destination),
    "cp -r " .. source .. " " .. destination,
  }, "\n")
end

return function(files)
  return sequence("std.files", entries(files), {
    migration = { add = pass_paths(copy_file), remove = pass_paths(remove_file) },
    sync = { add = pass_paths(copy_file) }
  }, serialize)
end
