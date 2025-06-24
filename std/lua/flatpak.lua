local sequence = require "@std.sequence"
local entries = require "@std.entries"
local serialize = require "@std.serialize"

local flatpak = {}

function flatpak.remotes(remotes)
  local remote_entries = entries(remotes)
  return sequence("std.flatpak.remotes", remote_entries, {
    migration = {
      add = function(remote)
        return "flatpak remote-add " .. remote[1] .. " " .. remote[2]
      end,
      remove = function(remote)
        return "flatpak remote-delete " .. remote[1]
      end
    }
  }, serialize)
end

return flatpak
