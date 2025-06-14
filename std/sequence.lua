local migrator = require "@std.migrator"
local set = require "@std.set"
local map = require "@std.map"

local function get_group_script(list, handle, handle_no)
  if #list == 0 or handle == nil then
    return handle_no .. "\n"
  elseif handle then
    return handle(list) .. "\n"
  end
  return ""
end

local function get_iterative_script(list, handle)
  if handle then
    return table.concat(map(list, handle), "\n") .. "\n"
  end
  return ""
end

local function get_sequence_script(added, removed, handlers)
  if type(handlers) == "string" then
    return handlers
  else
    handlers = handlers or {}
    return get_group_script(added, handlers.added, handlers.no_added or "# Nothing to add") ..
        get_group_script(removed, handlers.removed, handlers.no_removed or "# Nothing to remove") ..
        get_iterative_script(added, handlers.add) ..
        get_iterative_script(removed, handlers.remove)
  end
end

-- Creates a migrator that acts upon a set, adding various commands for added/removed items
return function(name, desired_sequence, handlers)
  return migrator(name, function(current_sequence)
    current_sequence = current_sequence or {}
    local current = set.from(current_sequence)
    local desired = set.from(desired_sequence)
    local added = set.to_sequence(set.diff(desired, current))
    local removed = set.to_sequence(set.diff(current, desired))
    return {
      next_state = desired_sequence,
      migration = get_sequence_script(added, removed, handlers.migration),
      sync = get_sequence_script(added, removed, handlers.sync)
    }
  end)
end
