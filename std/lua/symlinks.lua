local sequence = require "@std.sequence"
local entries = require "@std.entries"
local file = require "@std.file"
local serialize = require "@std.serialize"

local function create_symlink(source, destination)
  return table.concat({
    "# Symlinking " .. destination .. " to " .. source,
    file.remove(nil, destination),
    "ln -sT " .. source .. " " .. destination
  }, "\n")
end

return function(symlinks)
  local symlink_entries = entries(symlinks)
  return sequence("std.symlinks", symlink_entries, {
    migration = {
      add = file.pass_paths(create_symlink),
      remove = file.pass_paths(file.remove),
    }
  }, serialize)
end
