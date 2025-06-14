local migrator = require "@std.migrator"
local set = require "@std.set"
local map = require "@std.map"

return function(name, desired_sequence, handlers)
  return migrator(name, function(current_sequence)
    current_sequence = current_sequence or {}
    local current = set.from(current_sequence)
    local desired = set.from(desired_sequence)
    local added = set.to_sequence(set.diff(desired, current))
    local removed = set.to_sequence(set.diff(current, desired))
    return {
      next_state = desired_sequence,
      migration = table.concat({
        table.concat(map(added, function(item) return handlers.add(item) end), "\n"),
        table.concat(map(removed, function(item) return handlers.remove(item) end), "\n"),
        handlers.added(added),
        handlers.removed(removed),
      }, "\n"),
      sync = table.concat({
        table.concat(map(added, function(item) return handlers.add_sync(item) end), "\n"),
        table.concat(map(removed, function(item) return handlers.remove_sync(item) end), "\n"),
        handlers.added_sync(added),
        handlers.removed_sync(removed),
      }, "\n")
    }
  end)
end
