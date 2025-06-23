return function(t1)
  local entries = {}
  for key, value in pairs(t1) do
    table.insert(entries, { key, value })
  end
  return entries
end
