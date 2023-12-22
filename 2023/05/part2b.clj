(ns main-clj.core
  (:gen-class))

(require '[clojure.string :as str]
         '[instaparse.core :as insta]
         'clojure.stacktrace)

;; Parsing

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

;; Solving

(defprotocol update-and-map
  (update [thing fun])
  (map-update [thing fun]))

(deftype History [value history]
  update-and-map
  (update [h f]
    (History.
      (f (.value h))
      (conj (.history h) (.value h))))
  (map-update [h f]
    (map
      #(History. % (conj (.history h) (.value h)))
      (f (.value h)))))

(defn make-history
  ([v] (History. v []))
  ([v h] (History. v h)))

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

(defn prev-contains-next?
  "
    dst  dst-end
     v    v
  p: ■■■■■■
  n:  ■■■
      ^ ^
    src src-end
  "
  [n p]
  (and (< (dst p) (src n))
       (> (dst-end p) (src-end n))))

(defn prev-is-subset-of-next?
  "
      dst dst-end
        v v
  p:    ■■■    ■■■   ■■■       ■■■
  n:  ■■■■■■   ■■■   ■■■■■   ■■■■■
      ^    ^
     src  src-end
  "
  [n p]
  (and (>= (dst p) (src n))
       (<= (dst-end p) (src-end n))))

(defn prev-overlaps-next-end?
  "
    dst dst-end
     v  v
  p: ■■■■    ■■■■
  n: ■■    ■■■■
           ^  ^
         src  src-end
  "
  [n p]
  (and (>= (dst p) (src n))
       (< (dst p) (src-end n))
       (> (dst-end p) (src-end n))))

(defn prev-overlaps-next-begin?
  "
    dst dst-end
     v  v
  p: ■■■■  ■■■■
  n:   ■■    ■■■■
             ^  ^
            src src-end
  "
  [n p]
  (and (> (src n) (dst p))
       (< (src n) (dst-end p))
       (>= (src-end n) (dst-end p))))

(defn create-direct-mapping
  "Creates a direct mapping from (src prev-map) to (dst next-map)."
  [prev-map next-map]
  (cond
    (prev-contains-next? next-map prev-map)
    ;   src        dst
    ;    v          v
    ; p: ■■■■■■ --> ■■■■■■
    ; n:             ■■■   --> ■■■
    ;            src ^     dst ^
    ; =>  ■■■ ---------------> ■■■
    (create-mapping
      (dst next-map)
      (+ (src prev-map) (- (src next-map) (dst prev-map)))
      (len next-map))
    
    (prev-is-subset-of-next? next-map prev-map)
    ;   src     dst
    ;    v       v
    ; p: ■■■ --> ■■■
    ; n:        ■■■■■■ --> ■■■■■■
    ;       src ^      dst ^
    ; => ■■■ -------------> ■■■
    (create-mapping
      (+ (dst next-map) (- (dst prev-map) (src next-map)))
      (src prev-map)
      (len prev-map))
    
    (prev-overlaps-next-end? next-map prev-map)
    ;   src       dst
    ;    v         v
    ; p: ■■■■■ --> ■■■■■
    ; n:         ■■■■■ ----> ■■■■■
    ;        src ^       dst ^   ^ dst-end
    ; => ■■■ ----------------> ■■■
    (let [overlap-len (- (len next-map) (- (dst prev-map) (src next-map)))]
      (create-mapping
        (- (dst-end next-map) overlap-len)
        (src prev-map)
        overlap-len))
    
    (prev-overlaps-next-begin? next-map prev-map)
    ;   src       dst
    ;    v         v
    ; p: ■■■■■ --> ■■■■■
    ; n:             ■■■■ --> ■■■■
    ;            src ^    dst ^
    ; =>   ■■■ -------------> ■■■
    (let [before-len (- (src next-map) (dst prev-map))]
      (create-mapping
        (dst next-map)
        (+ (src prev-map) before-len)
        (- (len prev-map) before-len)))))

(defn create-inbetween-mapping
  "
  Given a map `prev-map` and a list of mappings `overlap-maps` that
  are all contained within `prev-map`, split `prev-map` into smaller
  maps so that the overlapping parts from `prev-map` are removed.
             p: ■■■■■■■■■■■■■■■■■
             o:  ■ ■■   ■■■ ■ ■■
             => ■ ■  ■■■   ■ ■  ■
  The returned value is the resulting maps and the overlapping maps in a
  sorted list.
  "
  [prev-map overlap-maps]
  (let [overlap-maps (sort-by src overlap-maps)]
    (loop [next (first overlap-maps)
           overlaps (rest overlap-maps)
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

(defn collapse-to-direct-mapping
  "
  Given a list of further mappings `map-data` and an initial mapping `mapping-el`,
  create direct mappings from the initial source to the respective
  destinations of the other mappings.
  `mapping-el` can be a range that spans across / overlaps with multiple of
  the ranges from `map-data`, but the ranges from `map-data` must not have any
  overlap with each other.
  "
  [map-data mapping-el]
  (->> map-data
       (map (partial create-direct-mapping mapping-el))
       (filter (comp not nil?))
       (create-inbetween-mapping mapping-el)))

(defn update-mapping
  [mappings map-data]
  (apply concat
    (map #(.map-update %
            (partial collapse-to-direct-mapping map-data))
         mappings)))

(defn collapse-mappings
  [mappings]
  (let [initial-mapping (map make-history (map (partial apply map-self) (partition 2 (:seeds mappings))))]
    (reduce
      update-mapping
      initial-mapping
      (map #(% mappings) traverse-order))))

(comment (def $input (slurp "sample.txt")))
(def $input (slurp "input.txt"))

(try
  (time
    (.value
      (->> (parse-input $input)
           collapse-mappings
           (sort-by #(.value %))
           first)))
  (catch Exception e
    (clojure.stacktrace/print-stack-trace e)))

;              Loc Seed 
; Test input: [46N 82N 10N]                     ("Elapsed time: 12.5392 msecs")
; Real input: [37806486N 1669061417N 30783317N] ("Elapsed time: 111.0806 msecs")
