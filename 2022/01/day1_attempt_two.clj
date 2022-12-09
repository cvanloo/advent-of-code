(ns day1_attempt_two
  (:require [clojure.string :as str]))

(def actual-inp (slurp "input.txt"))
(def test-inp (slurp "test.txt"))

(defn calc-elf-cal [s]
  (reduce
   (fn [a s]
     (if (str/blank? s)
       (conj a 0)
       (conj (pop a) (+ (last a) (Integer/parseInt s)))))
   [0] (str/split-lines s)))

(let [cals-per-elf (sort > (calc-elf-cal actual-inp))]
  (println "Part 1: " (first cals-per-elf))
  (println "Part 2: " (reduce + (take 3 cals-per-elf))))

; (out) Part 1:  73211
; (out) Part 2:  213958
