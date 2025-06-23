-- Create a shallow copy of t1, where every item in the sequence is passed through callback
return function(t1, callback)
  local t2 = {}
  for _, value in ipairs(t1) do
    table.insert(t2, callback(value))
  end
  return t2
end
