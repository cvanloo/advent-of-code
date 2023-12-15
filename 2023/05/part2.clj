(ns main-clj.core
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

(def dst (partial getter 0))
(def src (partial getter 1))
(def len (partial getter 2))

(defn resolve-mapping-backwards-line
  [m k]
  (if (<= (dst m) k (dec (+ (dst m) (len m))))
    (+ (src m) (- k (dst m)))))

(defn resolve-mapping-backwards
  [m k]
  (or (first (filter
               (comp not nil?)
               (map #(resolve-mapping-backwards-line % k) m)))
      k))

(def traverse-order [:seed-to-soil
                     :soil-to-fertilizer
                     :fertilizer-to-water
                     :water-to-light
                     :light-to-temperature
                     :temperature-to-humidity
                     :humidity-to-location])

(defn resolve-location-to-seed
  [m l]
  (reduce
    (fn [k mk]
      (resolve-mapping-backwards (get m mk) k))
    l
    (reverse traverse-order)))

(defn is-valid-seed?
  [seed-ranges s]
  (not (empty?
         (filter true?
           (map
             (fn [[start len]]
               (<= start s (dec (+ start len))))
             (partition 2 seed-ranges))))))


(def $input (slurp "input.txt"))

(defn -main []
  (let [mappings (parse-input $input)]
    (first (filter
             #(is-valid-seed? (:seeds mappings) (second %))
             (map
               #(vector % (resolve-location-to-seed mappings %))
               (range))))))


(time (-main))
; Part 2:
; "Elapsed time: 1853023.2711 msecs"
; => [37806486 1669061417N]
; Nearest location: 37806486
