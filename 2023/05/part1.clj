(ns part1
  (:gen-class))

(require '[clojure.string :as str]
         '[instaparse.core :as insta])

(def aoc-input
  (insta/parser
    "S = seeds mapping+
     seeds = <'seeds: '> numbers+
     numbers = (number <' '> numbers) | (number <#'\n'?>)
     number = #'[0-9]+'
     mapping = <ws> mapname <' map:'> <#'\n'> maplisting
     mapname = #'[a-zA-Z-]+'
     maplisting = numbers+
     ws = #'[ \n\r]+'"))

(defn parse-numbers
  [tokens]
  (loop [tokens tokens
         numbers []]
    (if (nil? tokens)
      numbers
      (recur
        (nth tokens 2 nil)
        (conj numbers (bigint (second (second tokens))))))))

(defn parse-seeds
  [tokens]
  {:seeds (parse-numbers (first (rest tokens)))})

(defn parse-mapping
  [tokens]
  (let [mapname (second (second tokens))
        maplisting (rest (first (drop 2 tokens)))]
    (hash-map
      (keyword mapname)
      (map parse-numbers maplisting))))

(defn parse-mappings
  [tokens]
  (map parse-mapping tokens))

(defn parse-input
  [input]
  (let [parsed (aoc-input input)
        seed-tokens (nth parsed 1)
        mappings-tokens (drop 2 parsed)
        seeds (parse-seeds seed-tokens)]
    (apply merge seeds (parse-mappings mappings-tokens))))





(defn getter [i m]
  (get m i))

(def dest (partial getter 0))
(def src (partial getter 1))
(def len (partial getter 2))

(defn resolve-mapping-line
  [m k]
  (if (<= (src m) k (dec (+ (src m) (len m))))
    (+ (dest m) (- k (src m)))))

(defn resolve-mapping
  [m k]
  (or (first (filter
               (comp not nil?)
               (map #(resolve-mapping-line % k) m)))
      k))

(def traverse-order [:seed-to-soil
                     :soil-to-fertilizer
                     :fertilizer-to-water
                     :water-to-light
                     :light-to-temperature
                     :temperature-to-humidity
                     :humidity-to-location])

(defn resolve-seed
  [m s]
  (reduce
    (fn [k mk] (resolve-mapping (get m mk) k))
    s
    traverse-order))

(defn resolve-seeds
  [m]
  (map
    (partial resolve-seed m)
    (:seeds m)))

(defn -main []
  (first (sort (resolve-seeds (parse-input (slurp "input.txt"))))))

; Part 1:
; => 1181555926N











