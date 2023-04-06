#!/bin/ruby

require 'json'
require 'stringio'

# Complete the compareTriplets function below.
def compareTriplets(a, b)
soma_a, soma_b = 0,0
soma_arrays = []
    if a[0] > b[0]
        soma_a +=1
    elsif a[0] < b[0]
        soma_b +=1
    else
        puts 'nothing changes'
    end

    if a[1] > b[1]
        soma_a +=1
        elsif a[1] < b[1]
            soma_b +=1
        else
            puts 'nothing changes'
    end
    
    if a[2] > b[2]
        soma_a +=1
    elsif a[2] < b[2]
        soma_b +=1
    else
        puts 'nothing changes'
    end
    soma_arrays << soma_a
    soma_arrays << soma_b
    
    soma_arrays
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

a = gets.rstrip.split.map(&:to_i)

b = gets.rstrip.split.map(&:to_i)

result = compareTriplets a, b

fptr.write result.join " "
fptr.write "\n"

fptr.close()
