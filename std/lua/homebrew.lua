local sequence = require "@std.sequence"

local homebrew = {}

function homebrew.packages(packages)
  return sequence("std.homebrew.packages", packages, {
    migration = {
      added = function(added)
        return "brew install " .. table.concat(added, " ")
      end,
      removed = function(removed)
        return "brew uninstall " .. table.concat(removed, " ")
      end,
    },
    sync = "brew update"
  })
end

return homebrew
