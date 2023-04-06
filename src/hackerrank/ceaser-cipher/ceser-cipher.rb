#!/bin/ruby

require 'json'
require 'stringio'

# Complete the caesarCipher function below.
def caesarCipher(s, k)
alphabet_downcase = Array('a'..'z')
alphabet_upcase = Array('A'..'Z')
  
  result = s.chars.map do |letter|
    if alphabet_downcase.include?(letter)
        r = (alphabet_downcase.index(letter)+ k) % 26
        alphabet_downcase[r]
    elsif alphabet_upcase.include?(letter)
        r = (alphabet_upcase.index(letter) + k ) % 26
        alphabet_upcase[r]
    else
        letter
    end
  end

  result.join
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

n = gets.to_i

s = gets.to_s.rstrip

k = gets.to_i

result = caesarCipher s, k

fptr.write result
fptr.write "\n"

fptr.close()
