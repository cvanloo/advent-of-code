(ns part1
  (:require (clojure [string :as str]
                     [set :as set])))

(defn parse-input
  [input]
  (map
   (fn [rucksack]
     (split-at (/ (count rucksack) 2) rucksack))
   input))

(defn find-duplicates
  [rucksacks]
  (reduce
   (fn [duplicates rucksack]
     (apply conj duplicates (set/intersection (set (first rucksack)) (set (last rucksack)))))
   [] rucksacks))

(defn prioritize
  [characters]
  (map
   (fn [int-char]
     (cond
       (<= (int \a) int-char (int \z))
       (- int-char 96)
       (<= (int \A) int-char (int \Z))
       (- int-char 38)))
   characters))

(defn main
  [filename]
  (let [input (str/split-lines (slurp filename))]
    (->> input
         parse-input
         find-duplicates
         (map #(int %))
         prioritize
         (reduce +))))

(main "test.txt") ; 157
(main "input.txt") ; 8243
