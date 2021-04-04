n = gets.to_i
input = gets.split().map(&:to_i)

flag = true
r = 0

while flag do
  if (input.select { |i| i%2==0 }).length == n then
    input.map! {|i| i/2}
    r+=1
  else
    flag = false
  end
end

print r
