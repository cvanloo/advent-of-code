(ns part1
  (:require (clojure [string :as str]
                     [set :as set])))

(defn inc-last
  [s]
  (conj (vec (butlast s)) (inc (last s))))

(defn parse-range
  [range-str]
  (->> (str/split range-str #"-")
       (map #(Integer/parseInt %))
       (inc-last)
       (apply range)))

(defn find-contains
  [[range1 range2]]
  (or
   (set/subset? (set range1) (set range2))
   (set/subset? (set range2) (set range1))))

(defn main
  [filename]
  (let [lines (str/split-lines (slurp filename))]
    (->> lines
         (map #(str/split % #","))
         (map #(map parse-range %))
         (map find-contains)
         (filter true?)
         (count))))

(main "test.txt") ; 2
(main "input.txt") ; 496
