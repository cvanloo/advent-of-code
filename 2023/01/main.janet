#(def peg-numbers
#  (peg/compile '{:number (+ :d "one" "two" "three" "four" "five" "six" "seven" "eight" "nine" "ten") 
#                 :main (any (+ (<- (* ($) :number)) 1))}))


(def peg-numbers
  (peg/compile '{:number (+ :d "one" "two" "three" "four" "five" "six" "seven" "eight" "nine" "ten") 
                 :main (any (+ (<- :number) 1))}))

(def number-names
  {"1" 1
   "2" 2
   "3" 3
   "4" 4
   "5" 5
   "6" 6
   "7" 7
   "8" 8
   "9" 9
   "one" 1
   "two" 2
   "three" 3
   "four" 4
   "five" 5
   "six" 6
   "seven" 7
   "eight" 8
   "nine" 9})

(defn merge-numbers
  [numbers]
  (+ (* (first numbers) 10) (last numbers)))

(->> (slurp "input.txt")
     (string/split "\n")
     (filter (comp not empty?))
     (map (fn [line] (peg/match peg-numbers line)))
     (map (fn [line] (map number-names line)))
     (map merge-numbers)
     (reduce + 0))
