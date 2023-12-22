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









(defn make-history
  ([v]
   {:result v
    :history []})
  ([v h]
   {:result v
    :history h}))

(defn update-history
  [{result :result history :history} f]
  (make-history (f result)
                (conj history result)))

(defn map-update-history
  [{result :result history :history} f]
  (map #(make-history % (conj history result))
       (f result)))

(defmacro print-and-ret
  [v]
  `(let [res# ~v]
     (println res#)
     res#))

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



















(defn collapse-overlap
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
      (create-mapping
        (dst next-map)
        (+ (src prev-map) (- (src next-map) (dst prev-map)))
        (len next-map))
      
      (prev-is-subset-of-next? next-map prev-map)
      (create-mapping
        (+ (dst next-map) (- (dst prev-map) (src next-map)))
        (src prev-map)
        (len prev-map))
      
      (prev-overlaps-next-end? next-map prev-map)
      (let [overlap-len (- (len next-map) (- (dst prev-map) (src next-map)))]
        (create-mapping
          (- (dst-end next-map) overlap-len)
          (src prev-map)
          overlap-len))
      
      (prev-overlaps-next-begin? next-map prev-map)
      (let [before-len (- (src next-map) (dst prev-map))]
        (create-mapping
          (dst next-map)
          (+ (src prev-map) before-len)
          (- (len prev-map) before-len))))))

(defn resolve-between
  [prev-map collapsed-mapping]
  (let [collapsed-mapping (sort-by src collapsed-mapping)]
    (loop [next (first collapsed-mapping)
           overlaps (rest collapsed-mapping)
           total-len 0
           betweens []]
      (if (nil? next)
        (if (= total-len (len prev-map))
          betweens
          (let [new-dst (+ (dst prev-map) total-len)
                new-src (+ (src prev-map) total-len)
                new-len (- (len prev-map) total-len)]
            (conj betweens
              (create-mapping new-dst new-src new-len))))
        
        (let [new-dst (+ (dst prev-map) total-len)
              new-src (+ (src prev-map) total-len)
              new-len (- (src next) new-src)]
          (recur
            (first overlaps)
            (rest overlaps)
            (+ total-len new-len (len next))
            (if (zero? new-len)
              (conj betweens next)
              (conj betweens
                (create-mapping new-dst new-src new-len)
                next))))))))


      
         

(let [prev-map [50 10 100]
      mapping [[25 50 20]
               [5 70 30]
               [125 100 50]]]
  (map
    (partial collapse-overlap prev-map)
    mapping))

(let [prev-map [50 10 100]
      mapping [[25 52 16]
               [5 70 27]
               [125 101 20]]]
  (map
    (partial collapse-overlap prev-map)
    mapping))

; [50 10 2] *[25 12 16]* [68 28 2] *[5 30 27]* [97 57 4] *[125 61 20]* [121 81 29]
(resolve-between
  [50 10 100]
  [[25 12 16]
   [5 30 27]
   [125 61 20]])
; => [[50 10 2] [25 12 16] [68 28 2] [5 30 27] [97 57 4] [125 61 20] [121 81 29]]

(resolve-between
  [50 10 100]
  [[25 10 20]
   [5 30 30]
   [125 60 50]])
; => [[25 10 20] [5 30 30] [125 60 50]]

(resolve-between
  [50 10 100]
  [])
; => [[50 10 100]]



(defn update-map-entry
  "mapping-el can be a range that spans across / overlaps with multiple of the
  ranges from map-data.
  The ranges from map-data must not have any overlap with each other."
  [map-data mapping-el]
  (->> map-data
       (map (partial collapse-overlap mapping-el))
       (filter (comp not nil?))
       (resolve-between mapping-el)))
         

(update-map-entry
  [[25 50 20]
   [5 70 30]
   [125 100 50]]
  [50 10 100])

(update-map-entry
  [[25 52 16]
   [5 70 27]
   [125 101 20]]
  [50 10 100])

(update-map-entry
  [[70 205 5]
   [80 357 3]]
  [50 10 100])




(defn update-mapping
  [mappings map-data]
  (apply concat
    (map (fn [m]
           (map-update-history
             m
             (partial update-map-entry map-data)))
         mappings)))








(defn collapse-mappings
  [mappings]
  (let [collapsed-mapping (map make-history (map (partial apply map-self) (partition 2 (:seeds mappings))))]
    (reduce update-mapping
            collapsed-mapping
            (map #(% mappings) traverse-order))))

(def $input (slurp "sample.txt"))
; => [46N 82N 10N] (the correct result)

(comment (def $input (slurp "input.txt")))
; => {:result [37806486N 1669061417N 30783317N], :history [[1499552892N 1499552892N 200291842N] [3654918290N 1499552892N 200291842N] [3855550220N 1669061417N 30783317N] [3855550220N 1669061417N 30783317N] [296168274N 1669061417N 30783317N] [860738727N 1669061417N 30783317N] [37806486N 1669061417N 30783317N]]}
; => 37806486N is the correct result!

(time (->> (parse-input $input)
           collapse-mappings
           (sort-by first)
           first))



(defn -main []
  ())


