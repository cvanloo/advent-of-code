(ns part2b
  (:require (clojure [string :as str])))

(defn parse-nums
  [line]
  (map #(Integer/parseInt %) (str/split line #"(-|,)")))

(defn overlaps?
  [[s1 e1 s2 e2]]
  (or
    (and (>= s1 s2) (<= s1 e2))
    (and (>= s2 s1) (<= s2 e1))))

(defn main
  [filename]
  (let [lines (str/split-lines (slurp filename))]
    (->> lines
         (map parse-nums)
         (filter overlaps?)
         (count))))

(main "test.txt") ; 4
(main "input.txt") ; 847
