(require 'clojure.string)

(comment (def $input "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..")

(def $input (slurp "input.txt"))

(def $line-length (count (first (clojure.string/split-lines $input))))
(def $line-count (count (clojure.string/split-lines $input)))

(def digit? #(Character/isDigit %))
(def gear? #(= % \*))
(def newline? #(= % \newline))

(defn parse-characters
  [input]
  (select-keys
    (reduce
      (fn [{numbers :numbers gears :gears pos :pos line :line, :as acc}
           n]
        (cond
          (digit? n) {:numbers (conj numbers [n line pos])
                      :gears gears
                      :pos (inc pos)
                      :line line}
          (gear? n) {:numbers numbers
                     :gears (conj gears [n line pos])
                     :pos (inc pos)
                     :line line}
          (newline? n) {:numbers numbers
                        :gears gears
                        :pos 0
                        :line (inc line)}
          :else {:numbers numbers
                 :gears gears
                 :pos (inc pos)
                 :line line}))
      {:numbers []
       :gears []
       :pos 0
       :line 0}
      input)
    [:numbers :gears]))

(defn parse-numbers
  [digits-pos]
  (:numbers
    (reduce
      (fn [{numbers :numbers last-line :line last-pos :pos, :as acc}
           [digit line pos]]
        {:line line
         :pos pos
         :numbers (if (and
                           (= line last-line)
                           (= 1 (- pos last-pos)))
                    (let [[number-begin line start-pos _] (last numbers)
                          other-numbers (pop numbers)]
                      (conj other-numbers [(str number-begin digit) line start-pos pos]))
                    (conj numbers [(str digit) line pos pos]))})
      {:numbers []
       :line -1
       :pos 0}
      digits-pos)))

(defn parse
  [input]
  (update-in
    (parse-characters input)
    [:numbers]
    parse-numbers))

(defn neighbour-fields
  [[number line start end]]
  [number
   (max (dec line) 0)
   (min (inc line) (dec $line-length))
   (max (dec start) 0)
   (min (inc end) (dec $line-count))])

(defn overlaps?
  [[number line-start line-end pos-start pos-end]
   [_ line pos]]
  (if (and (<= line-start line line-end)
           (<= pos-start pos pos-end))
    number))



(def parsed-input (parse $input))
(def overlap-preds (->> (:numbers parsed-input)
                        (map neighbour-fields)
                        (map #(partial overlaps? %))))

(->> (:gears parsed-input)
     (map (fn [gear] (filter
                       (comp not nil?)
                       (map #(% gear) overlap-preds))))
     (filter #(>= (count %) 2))
     (map (fn [ns] (map #(Integer/parseInt %) ns)))
     (map (partial apply *))
     (reduce +))

; Part 2:
; => 78826761
