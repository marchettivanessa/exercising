#!/bin/ruby

require 'json'
require 'stringio'

# Complete the pangrams function below.
def pangrams(s)
   if ('a'..'z').all? {|letra| s.downcase.include?(letra)}
   return "pangram"
else
    return "not pangram"
end
end
fptr = File.open(ENV['OUTPUT_PATH'], 'w')

s = gets.to_s.rstrip

result = pangrams s

fptr.write result
fptr.write "\n"

fptr.close()
