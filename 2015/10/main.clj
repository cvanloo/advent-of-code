(def input "3113322113")

(defn parse-string
  [s]
  (map #(Integer/parseInt (str %)) s))

(defn partition-same
  [xs]
  (reduce (fn [acc n]
            (if-let [prev (last (last acc))]
              (if (= prev n)
                (conj (pop acc) (conj (last acc) n))
                (conj acc [n]))
              (conj acc [n])))
          []
          xs))

(defn count-same
  [xxs]
  (reduce (fn [acc n]
            (conj acc (count n) (first n)))
          []
          xxs))

(def look-and-say (comp count-same partition-same parse-string))

(look-and-say "211333") ; => [1 2 2 1 3 3]

(time (count (loop [i 40
                    input input]
              (if (zero? i)
                input
                (recur (dec i) (look-and-say input))))))

; part 1 (40x): 329356
; part 2 (50x): 
