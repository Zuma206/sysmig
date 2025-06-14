return function(t1, callback)
  local t2 = {}
  for _, value in ipairs(t1) do
    table.insert(t2, callback(value))
  end
  return t2
end
