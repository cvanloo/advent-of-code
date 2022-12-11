(ns part1b
  (:require (clojure [string :as str])))

(defn parse-nums
  [line]
  (map #(Integer/parseInt %) (str/split line #"(-|,)")))

(defn fully-contains?
  [[s1 e1 s2 e2]]
  (or
    (and (>= s1 s2) (<= e1 e2))
    (and (>= s2 s1) (<= e2 e1))))

(defn main
  [filename]
  (let [lines (str/split-lines (slurp filename))]
    (->> lines
         (map parse-nums)
         (filter fully-contains?)
         (count))))

(main "test.txt") ; 2
(main "input.txt") ; 496
