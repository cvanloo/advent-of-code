(ns part1
  (:require [clojure.string :as str]))

(defn find-marker
  [line]
  (loop [line line idx 14]
    (let [characters (distinct (take 14 line))]
      (if (= (count characters) 14)
       idx
       (recur (rest line) (inc idx))))))

(map find-marker (str/split-lines (slurp "test.txt"))) ; (19 23 23 29 26)

(find-marker (slurp "input.txt")) ; 2315
