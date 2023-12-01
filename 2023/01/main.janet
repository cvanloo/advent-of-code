(def peg-numbers
  (peg/compile '{:number (+ (* (> 0 "one")   (constant 1) 1)
                            (* (> 0 "two")   (constant 2) 1)
                            (* (> 0 "three") (constant 3) 1)
                            (* (> 0 "four")  (constant 4) 1)
                            (* (> 0 "five")  (constant 5) 1)
                            (* (> 0 "six")   (constant 6) 1)
                            (* (> 0 "seven") (constant 7) 1)
                            (* (> 0 "eight") (constant 8) 1)
                            (* (> 0 "nine")  (constant 9) 1))
                 :main (any (+ :number (number :d) 1))}))

(defn merge-numbers
  [numbers]
  (+ (* (first numbers) 10) (last numbers)))

(->> (slurp "input.txt")
     (string/split "\n")
     (filter (comp not empty?))
     (map (fn [line] (peg/match peg-numbers line)))
     (map merge-numbers)
     (reduce + 0))

# => 54824
