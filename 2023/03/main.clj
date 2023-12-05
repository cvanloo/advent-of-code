(require '[clojure.string :as str])

(def $input "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..")

(def $input-lines
  (str/split-lines $input))

(def $line-length 10)
(def $line-count 9)

(defn find-number-positions
  [text]
  (filter #(Character/isDigit (first %))
          (partition 2 (interleave text (range)))))

(defn find-number-lines
  [number-pos]
  (partition 2 (interleave (range) number-pos)))

(defn combine-consecutive-numbers
  [[line number-positions]]
  (:result
    (reduce
      (fn [{result :result last-idx :last-idx} [num idx]]
        {:last-idx idx
         :result (if (= 1 (- idx last-idx))
                   (let [[num-begin line pos-start] (first result)]
                     (conj (rest result) [(str num-begin num) line pos-start idx]))
                   (conj result [num line idx]))})
      {:result '()
       :last-idx -2} 
      number-positions)))

(defn neighbour-idxs
  [line start end]
  (let [line-top (max (dec line) 0)
        line-bot (min (inc line) $line-count)
        start (max (dec start) 0)
        end (min (inc end) $line-length)]
    (map #(vector % start end)
         (range line-top (inc line-bot)))))

(defn neighbour-fields
  [l s e]
  (take (- e (dec s)) (drop s (nth $input-lines l))))

(defn has-non-empty-neighbour?
  [[_ l s e]]
  (->> (neighbour-idxs l s e)
       (map (partial apply neighbour-fields))
       flatten
       (some #(contains? #{\$ \# \+ \*} %))))

(->> $input
     str/split-lines
     (map find-number-positions)
     find-number-lines
     (map combine-consecutive-numbers)
     (apply concat)
     (filter has-non-empty-neighbour?))
