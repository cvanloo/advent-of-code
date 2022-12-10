(ns part1
  (:require [clojure.string :as str]))

(def test-inp (slurp "input.txt"))

(def opponent-mapping {"A" 1, "B" 2, "C", 3})
(def my-mapping {"X" 1, "Y" 2, "Z", 3})
(def outcome-mapping {0 3, 1 0, 2 6})

(def parsed (map #(str/split % #" ") (str/split-lines test-inp)))

(reduce
 (fn [score play]
   (let [my-move (my-mapping (last play))
         op-move (opponent-mapping (first play))]
     (+ score my-move (outcome-mapping (mod (- op-move my-move) 3)))))
 0 parsed)
