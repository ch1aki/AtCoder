r = 0
for s in gets.chars.map(&:to_i) do
  r+=1 if s==1
end
print r
