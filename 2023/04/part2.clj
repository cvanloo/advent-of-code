(require '(clojure [string :as str]
                   [set :as set]))

(comment (def $input "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11")

(defn number-try-parse
  [number-str]
  (try (Integer/parseInt number-str)
    (catch NumberFormatException e)))

(defn parse-numbers
  [numbers-str]
  (->> (str/split numbers-str #" ")
       (map number-try-parse)
       (filter (comp not nil?))))

(defn parse-line
  [line]
  (let [[winners numbers] (str/split line #"\|")]
    (map parse-numbers [winners numbers])))

(defn scratchcard-matches
  [scratchcard]
  (let [freqs (map frequencies scratchcard)
        common-numbers (apply set/intersection (map (comp set keys) freqs))
        number-of-matches (reduce + (map #((second freqs) %) common-numbers))]
    number-of-matches))

(defn count-cards
  [matches]
  (reduce
    (fn [counts [amount index]]
      (let [amount-of-cur-card (nth counts (dec index))
            begin (take index counts)
            middle (take amount (drop index counts))
            end (drop (+ amount index) counts)]
        (concat begin (map #(+ % amount-of-cur-card) middle) end)))
    (vec (repeat (count matches) 1))
    (partition 2 (interleave matches (rest (range))))))
 
(->> $input
     clojure.string/split-lines
     (map parse-line)
     (map scratchcard-matches)
     count-cards
     (reduce +))

; Part 2:
; => 8736438
