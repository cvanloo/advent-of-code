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

(defn pow
  [a n]
  ; note that (apply * '()) results in 1
  (apply * (repeat n a)))

(defn scratchcard-value
  [scratchcard]
  (let [freqs (map frequencies scratchcard)
        common-numbers (apply set/intersection (map (comp set keys) freqs))
        number-of-matches (reduce + (map #((second freqs) %) common-numbers))]
    (if (= 0 number-of-matches)
      0
      (pow 2 (dec number-of-matches)))))

(def $input (slurp "input.txt"))

(->> $input
     clojure.string/split-lines
     (map parse-line)
     (map scratchcard-value)
     (reduce +))

; Part 1:
; => 24542
