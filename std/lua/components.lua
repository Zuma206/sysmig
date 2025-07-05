local sequence = require "@std.sequence"
local map = require "@std.map"

local function get_key(component)
  return component.name
end

return function(components)
  local required_components = map(components, function(component)
    if type(component) == "string" then
      if type(components.module) == "string" then
        return require(components.module .. "." .. component)
      end
      return require(component)
    end
    return component
  end)
  return sequence("std.components", required_components, {
    migration = {
      add = function(component)
        return component.mount or ""
      end,
      remove = function(component)
        return component.unmount or ""
      end
    },
    sync = table.concat(
      map(required_components, function(component)
        return component.sync or ""
      end),
      "\n"
    )
  }, get_key)
end
