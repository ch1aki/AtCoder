a,b = gets.split().map(&:to_i)

if a*b%2==0 then
  print "Even"
else
  print "Odd"
end

