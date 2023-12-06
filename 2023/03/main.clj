(require '[clojure.string :as str])

(def $input-lines (str/split-lines (slurp "input.txt")))
(def $line-length (count (first $input-lines)))
(def $line-count (count $input-lines))

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
                   (conj result [num line idx idx]))})
      {:result '()
       :last-idx -5} 
      number-positions)))

(defn neighbour-idxs
  [line start end]
  (let [line-top (max (dec line) 0)
        line-bot (min (inc line) (dec $line-count))
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
       (some #(not (or (Character/isDigit %) (= % \.))))))

(has-non-empty-neighbour? [\4 3 56])

(->> $input-lines
     (map find-number-positions)
     find-number-lines
     (map combine-consecutive-numbers)
     (apply concat)
     (filter has-non-empty-neighbour?)
     (map first)
     (map str)
     (map #(Integer/parseInt %))
     (reduce +))

; Part 1
; => 533784
