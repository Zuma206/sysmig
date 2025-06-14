local migrator = require "@std.migrator"

return migrator("nothing", function()
  local script = "# std.nothing"
  return {
    migration = script,
    next_state = nil,
    sync = script,
  }
end)
