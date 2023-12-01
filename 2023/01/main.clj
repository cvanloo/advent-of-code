; Part 1:
(reduce
  +
  (map (fn [line]
         (->> line
              (filter #(Character/isDigit %))
              (#(str (first %) (last %)))
              (#(Integer/parseInt %))))
       (clojure.string/split-lines (slurp "input.txt"))))
; => 55386

; Part 2:
(def spelled-out-numbers
  {"one" 1
   "two" 2
   "three" 3
   "four" 4
   "five" 5
   "six" 6
   "seven" 7
   "eight" 8
   "nine" 9})

(def numbers-map
  (apply merge
    spelled-out-numbers
    (concat
      (map (fn [[name digit]]
             {(apply str (reverse name)) digit})
        spelled-out-numbers)
      (map #(hash-map (str %) %) (range 1 10)))))

(defn match-number-name
  [text]
  (first (filter (fn [[name _]]
                   (= name (apply str (take (count name) text))))
            numbers-map)))

(defn find-first-number
  [text]
  (if (empty? text)
    nil
    (if-let [[_ digit] (match-number-name text)]
       digit
       (recur (rest text)))))

(def find-last-number (comp find-first-number reverse))

(->> (slurp "input.txt")
     clojure.string/split-lines
     (map #(+ (* (find-first-number %) 10) (find-last-number %)))
     (reduce +))

; => 54824
