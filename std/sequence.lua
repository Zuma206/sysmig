local migrator = require "@std.migrator"
local set = require "@std.set"
local map = require "@std.map"

return function(name, desired_sequence, handlers)
  return migrator(name, function(current_sequence)
    local current = set.from(current_sequence)
    local desired = set.from(desired_sequence)
    local added = set.to_sequence(set.diff(desired, current))
    local removed = set.to_sequence(set.diff(current, desired))
    return {
      next_state = desired_sequence,
      migration =
          table.concat(map(added, function(item) return handlers.add(item) end), "\n") ..
          "\n" .. table.concat(map(removed, function(item) return handlers.remove(item) end), "\n") ..
          "\n" .. handlers.added(added) ..
          "\n" .. handlers.removed(removed),
      sync =
          table.concat(map(added, function(item) return handlers.add_sync(item) end), "\n") ..
          "\n" .. table.concat(map(removed, function(item) return handlers.remove_sync(item) end), "\n") ..
          "\n" .. handlers.added_sync(added) ..
          "\n" .. handlers.removed_sync(removed),
    }
  end)
end
