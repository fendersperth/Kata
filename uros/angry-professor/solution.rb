# test "ruby solution.rb < input.txt"
def gets_da_input
  gets.to_i.times.map do
    n, k = gets.split.map(&:to_i)    
    gets.split.map(&:to_i).select do |x| 
        x <= 0
    end.length < k
  end
end

gets_da_input.each do |result|
   result == true
    unless result
        puts "NO"
    else
        puts "YES"
    end
end
