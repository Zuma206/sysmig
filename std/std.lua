local std = {}

std.migrator = require "@std.migrator"
std.system = require "@std.system"

-- A blank migrator that does absolutely nothing
std.nothing = std.migrator("nothing", function()
  local script = "# std.nothing"
  return {
    migration = script,
    next_state = nil,
    sync = script,
  }
end)

return std
