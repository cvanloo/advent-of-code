(ns part1
  (:require [clojure.string :as str]))

(defn find-marker
  [line]
  (loop [line line idx 4]
    (let [characters (set (take 4 line))]
      (if (= (count characters) 4)
       idx
       (recur (rest line) (inc idx))))))

(map find-marker (str/split-lines (slurp "test.txt"))) ; (7 5 6 10 11)

(find-marker (slurp "input.txt")) ; 1538
