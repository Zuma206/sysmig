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

function flatpak.packages(packages)
  return sequence("std.flatpak.packages", packages, {
    migration = {
      added = function(packages)
        return "flatpak install -y " .. table.concat(packages, " ")
      end,
      removed = function(packages)
        return "flatpak remove -y " .. table.concat(packages, " ")
      end
    },
    sync = "sudo flatpak update"
  })
end

return flatpak
