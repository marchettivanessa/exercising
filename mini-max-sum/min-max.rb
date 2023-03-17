#!/bin/ruby

require 'json'
require 'stringio'

#
# Complete the 'miniMaxSum' function below.
#
# The function accepts INTEGER_ARRAY arr as parameter.
#

def miniMaxSum(arr)
    menor_valor = arr.sum - arr.max
    maior_valor = arr.sum - arr.min
    puts "#{menor_valor} #{maior_valor}"

end

arr = gets.rstrip.split.map(&:to_i)

miniMaxSum arr
