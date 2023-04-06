#!/bin/ruby

require 'json'
require 'stringio'

#
# Complete the 'equalizeArray' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY arr as parameter.
#

def equalizeArray(arr)
    # Write your code here
    hash = Hash.new(0)
    arr.each do |value|
        hash[value] +=1
    end
    arr.size - hash.values.max
 end
    

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

n = gets.strip.to_i

arr = gets.rstrip.split.map(&:to_i)

result = equalizeArray arr

fptr.write result
fptr.write "\n"

fptr.close()
