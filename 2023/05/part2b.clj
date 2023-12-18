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
(def dst-end #(+ (dst %) (len %)))
(def src-end #(+ (src %) (len %)))

(defn create-mapping
  [to from len]
  [to from len])

(defn map-self
  [from len]
  (create-mapping from from len)) 

(def traverse-order [:seed-to-soil
                     :soil-to-fertilizer
                     :fertilizer-to-water
                     :water-to-light
                     :light-to-temperature
                     :temperature-to-humidity
                     :humidity-to-location])

(defn resolve-overlap
  "Dear future me: I do not expect you to understand this. Sorry."
  [prev-map next-map]
  (letfn [(prev-contains-next? [n p]
            (and (< (dst p) (src n))
                 (> (dst-end p) (src-end n))))
          (prev-is-subset-of-next? [n p]
            (and (>= (dst p) (src n))
                 (<= (dst-end p) (src-end n))))
          (prev-overlaps-next-end? [n p]
            (and (>= (dst p) (src n))
                 (< (dst p) (src-end n))
                 (> (dst-end p) (src-end n))))
          (prev-overlaps-next-begin? [n p]
            (and (> (src n) (dst p))
                 (< (src n) (dst-end p))
                 (>= (src-end n) (dst-end p))))]
    (cond
      (prev-contains-next? next-map prev-map)
      (let [before-len (- (src next-map) (dst prev-map))
            overlap-len (len next-map)
            after-len (- (len prev-map) (+ before-len overlap-len))]
        [(create-mapping (dst prev-map)
                         (src prev-map)
                         before-len)
         (create-mapping (dst next-map)
                         (+ (src prev-map) before-len)
                         overlap-len)
         (create-mapping (- (dst-end prev-map) after-len)
                         (- (src-end prev-map) after-len)
                         after-len)])

      (prev-is-subset-of-next? next-map prev-map)
      [(create-mapping (+ (dst next-map) (- (dst prev-map) (src next-map)))
                       (src prev-map)
                       (len prev-map))]

      (prev-overlaps-next-end? next-map prev-map)
      (let [overlap-len (- (len next-map) (- (dst prev-map) (src next-map)))
            after-len (- (len prev-map) overlap-len)]
        [(create-mapping (- (dst-end next-map) overlap-len)
                         (src prev-map)
                         overlap-len)
         (create-mapping (+ (dst prev-map) overlap-len)
                         (+ (src prev-map) overlap-len)
                         after-len)])

      (prev-overlaps-next-begin? next-map prev-map)
      (let [before-len (- (src next-map) (dst prev-map))]
        [(create-mapping (dst prev-map)
                         (src prev-map)
                         before-len)
         (create-mapping (dst next-map)
                         (+ (src prev-map) before-len)
                         (- (len prev-map) before-len))]))))

(defn update-map-entry
  [mapping-el map-data]
  (or (->> map-data
          (map (partial resolve-overlap mapping-el))
          (apply concat)
          (#(if (empty? %) nil %)))
      [mapping-el]))

(defn update-mapping
  [mapping map-data]
  (loop [mapping-el (first mapping)
         mapping (rest mapping)
         collapsed-mapping []]
    (if (nil? mapping-el)
      (apply concat collapsed-mapping)
      (recur (first mapping)
             (rest mapping)
             (conj collapsed-mapping
                   (update-map-entry mapping-el map-data))))))

(defn collapse-mappings
  [mappings]
  (let [collapsed-mapping (map (partial apply map-self) (partition 2 (:seeds mappings)))]
    (reduce update-mapping
            collapsed-mapping
            (map #(% mappings) traverse-order))))

(def $input (slurp "sample.txt"))
; Expected result => [46 82N]

(comment (def $input (slurp "input.txt")))

(time (let [mappings (collapse-mappings (parse-input $input))]
       (->> mappings
            (sort-by first)
            first)))
