#!/bin/ruby

#
# Complete the simpleArraySum function below.
#
def simpleArraySum(ar)
    #
    # Write your code here.
    #puts(ar.reduce{|soma, numero| soma+numero})    
    
    return ar.inject(:+)

end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

ar_count = gets.to_i

ar = gets.rstrip.split(' ').map(&:to_i)

result = simpleArraySum ar

fptr.write result
fptr.write "\n"

fptr.close()
