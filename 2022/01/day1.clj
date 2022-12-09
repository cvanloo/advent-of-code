(ns day1
  (:require [clojure.string :as str]))

(def test-inp "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000")

(def actual-inp (slurp "input.txt"))

(defn calc-elf-cal [s]
  (reduce #(if (str/blank? %2)
             (conj %1 0)
             (conj (pop %1) (+ (last %1) (Integer/parseInt %2))))
          [0] (str/split-lines s)))

(defn sort-cal [c]
  (sort-by last > (map-indexed (fn [i v] (vector (inc i) v)) c)))

(defn max-cal [c]
  (first (sort-cal c)))

(defn print-cal [[n cal]]
  (println (str "The largest amount of calories is carried by the " n "th Elf with " cal " calories.")))

(defn top-three-n [c]
  (reduce (fn [a [n x]]
            (conj [] (conj (first a) n) (+ (second a) x)))
          [[] 0] (take 3 c)))

(defn print-3-cal [[n x]]
  (println (str "The top 3 elfs are " n " and carry together " x " calories.")))

(let [cals-per-elf (calc-elf-cal actual-inp)]
  (print-cal (max-cal cals-per-elf))
  (print-3-cal (top-three-n (sort-cal cals-per-elf))))

; (out) The largest amount of calories is carried by the 36th Elf with 73211 calories.
; (out) The top 3 elfs are [36 65 165] and carry together 213958 calories.
