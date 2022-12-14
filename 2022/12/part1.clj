(ns part1
  (:require [clojure.string :as str]))

(defn find-start-pos
  [field]
  (let [max-y (count field)
        max-x (count (first field))]
    (loop [x 0 y 0]
      (cond
        (>= x max-x)
        (recur 0 (inc y))
        (>= y max-y)
        nil
        (= \S (nth (nth field y) x))
        [x y]
        :else (recur (inc x) y)))))

;(find-start-pos (map seq (str/split-lines (slurp "test.txt"))))
;(find-start-pos (map seq (str/split-lines (slurp "input.txt"))))

(defn neighbours
  "Returns neighbours in order up, down, left, right"
  [field [posx posy]]
  (apply merge
         (for [[x y] '([0 -1] [0 +1] [-1 0] [+1 0])]
           (let [x (+ posx x)
                 y (+ posy y)
                 val (nth (nth field y nil) x nil)]
             {val [x y]}))))

;(neighbours (map seq (str/split-lines (slurp "test.txt"))) [3 3])
;(neighbours (map seq (str/split-lines (slurp "test.txt"))) [7 5])

(defn index
  [field [x y]]
  (nth (nth field y) x))

(defn height-difference
  [field pos1 pos2]
  (let [pos1 (int (index field pos1))
        pos2 (int (index field pos2))]
    (Math/abs (- pos2 pos1))))

;(height-difference (map seq (str/split-lines (slurp "test.txt"))) [0 0] [0 1])
;(height-difference (map seq (str/split-lines (slurp "test.txt"))) [0 1] [0 0])

(defn find-path
  [field start-pos]
  (let [ns (neighbours field start-pos)]
    (->> ns
         (filter seq)
         (filter #(>= 1 (height-difference field start-pos (val %)))))))

(find-path (map seq (str/split-lines (slurp "test.txt"))) [0 0])

(first (first {\i [7 4], nil [8 5]}))
