(->> (slurp "input.txt")
  (peg/match '(any (+ (number (* (? "-") :d+) nil) 1)))
  (reduce + 0)
  print)
; => part 1: 191164

; the input does not contain any formatting or whitespace, so we won't handle it here
(def json (peg/compile
            '{:object (* "{" (group (any :keyvals)) "}")
              :keyvals (* :str ":" (+ :object :array (<- :str) :num) (? ","))
              :array (* "[" (any (* (+ :object :array :str :num) (? ","))) "]")
              :num (number (* (? "-") :d+) nil)
              :str (* "\"" (any (range " !" "#~")) "\"")
              :main (any (+ :object :array 1))}))

; (peg/match json "{\"key\":\"red\",\"key2\":{\"key3\":\"hello\",\"keyarray\":[1,2,\"red\",{\"k1\":\"red\",\"k2\":777}]}}")
; => @[@["\"red\"" @["\"hello\"" "1" "2" @["\"red\"" "777"]]]]

(defn has-red? [xs]
  (some (fn [n] (= n "\"red\"")) xs))

(defn solve [xs]
  (reduce (fn [a n]
            (cond
              (number? n) (+ a n)
              (indexed? n) (if (has-red? n)
                           a
                           (+ a (solve n)))
              :else a))
          0
          xs))

; (solve (peg/match json (slurp "input.txt")))
; => part 2: 87842
