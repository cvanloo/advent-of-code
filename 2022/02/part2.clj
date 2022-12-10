(ns part2
  (:require [clojure.string :as str]))

(def test-inp (slurp "input.txt"))

(def opponent-mapping {"A" 1, "B" 2, "C", 3})
(def required-outcome {"X" "lose", "Y" "draw", "Z", "win"})
(def outcome-mapping {"draw" 3, "lose" 0, "win" 6})

(def parsed (map #(str/split % #" ") (str/split-lines test-inp)))

(defn move-for-strategy
  [strategy opponent-move]
  (case strategy
    ; Re-assign moves to be 0->Rock, 1->Paper, 2->Scissors
    ; Formula from https://github.com/elh/advent-2022/blob/main/src/advent_2022/day_2.clj
    "lose" (inc (mod (- opponent-move 2) 3))
    "draw" opponent-move
    "win" (inc (mod opponent-move 3))))

(reduce
 (fn [score play]
   (let [outcome (required-outcome (last play))
         op-move (opponent-mapping (first play))
         my-move (move-for-strategy outcome op-move)]
     (+ score my-move (outcome-mapping outcome))))
 0 parsed)
