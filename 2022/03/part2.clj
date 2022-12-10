(ns part2
  (:require (clojure [string :as str]
                     [set :as set])))

(defn find-badges
  [groups]
  (reduce
   (fn [badges group]
     (apply conj badges (apply set/intersection group)))
   [] groups))

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
         (map set)
         (partition 3)
         (find-badges)
         (map int)
         (prioritize)
         (reduce +))))

(main "test.txt") ; 70
(main "input.txt") ; 2631

