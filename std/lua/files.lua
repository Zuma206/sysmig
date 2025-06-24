local serialize = require "@std.serialize"
local sequence = require "@std.sequence"
local entries = require "@std.entries"
local dir = require "@std.dir"
local file = require "@std.file"
local map = require "@std.map"

local function copy_file(source, destination)
  return table.concat({
    "# Copying " .. source .. " to " .. destination,
    file.remove(nil, destination),
    "mkdir -p " .. dir(destination),
    "cp -r " .. source .. " " .. destination,
  }, "\n")
end

return function(files)
  local file_entries = entries(files)
  local pass_paths_copy_file = file.pass_paths(copy_file)
  return sequence("std.files", file_entries, {
    migration = { add = pass_paths_copy_file, remove = file.pass_paths(file.remove) },
    sync = table.concat(map(file_entries, pass_paths_copy_file), "\n")
  }, serialize)
end
