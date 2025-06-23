local sequence = require "@std.sequence"
local rhel = {}

function rhel.packages(desired_packages)
  return sequence("std.rhel.packages", desired_packages, {
    migration = {
      added = function(added)
        return "sudo dnf install -y " .. table.concat(added, " ")
      end,
      removed = function(removed)
        return "sudo dnf remove -y " .. table.concat(removed, " ")
      end
    },
    sync = "sudo dnf update -y"
  })
end

return rhel
