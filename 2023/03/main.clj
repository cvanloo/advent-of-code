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

(def $line-length 10)

(defn find-number-positions
  [text]
  (filter #(Character/isDigit (first %))
          (partition 2 (interleave text (range)))))

(defn combine-consecutive-numbers
  [number-pos]
  (:result
    (reduce
      (fn [{result :result last-idx :last-idx} [num idx]]
        {:last-idx idx
         :result (if (= 1 (- idx last-idx))
                   (let [[num-begin pos-start] (first result)]
                     (conj (rest result) [(str num-begin num) pos-start idx]))
                   (conj result [num idx]))})
      {:result '()
       :last-idx 0} 
      number-pos)))

(defn neighbour-idxs
  [i j]
  [(max (dec i) 0)
   (min (inc j) $line-length)])


(->> $input
     str/split-lines
     (map find-number-positions)
     (map combine-consecutive-numbers))
