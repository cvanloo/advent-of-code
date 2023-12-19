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

(letfn [(test [prev-map next-map expected-result]
          [prev-map next-map expected-result])

        (run-test [tests]
          (filter (fn [[_ _ expected-result actual-result]]
                    (not (= expected-result actual-result)))
                  (map (fn [[prev-map next-map expected-result]]
                           [prev-map next-map expected-result
                            (resolve-overlap prev-map next-map)])
                       tests)))

        (print-result [tests]
          (doseq [[prev-map next-map expected-result actual-result] tests]
            (println "p" prev-map "n" next-map "expected" expected-result "actual" actual-result)))]

  (print-result
    (run-test [; Complete overlap
               (test [20 5 7] [60 20 7] [[60 5 7]])
               (test [20 5 7] [55 15 14] [[60 5 7]])
               (test [20 5 7] [60 20 14] [[60 5 7]])
               (test [20 5 7] [53 13 14] [[60 5 7]])
               ; Middle of first range overlaps with second range
               (test [45 10 7] [5 47 3] [[45 10 2] [5 12 3] [50 15 2]])
               ; Beginning overlaps
               (test [70 10 20] [60 70 5] [[60 10 5] [75 15 15]])
               (test [70 10 20] [60 63 12] [[67 10 5] [75 15 15]])
               ; End overlaps
               (test [40 10 20] [90 55 5] [[40 10 15] [90 25 5]])])))
               ;(test [40 10 20] [] [[] []])])))











(defn update-map-entry
  "mapping-el can be a range that spans across / overlaps with multiple of the
   ranges from map-data.
   The ranges from map-data must not have any overlap with each other."
  [map-data mapping-el]
  (or (->> map-data
           (map (partial resolve-overlap mapping-el))
           (apply concat)
           (#(if (empty? %) nil %)))
      [mapping-el]))

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
; => evaluates to [0N 1662378336N 37466398N] which is clearly wrong

(time (->> (parse-input $input)
           collapse-mappings
           (sort-by first)
           first))

(time (->> (parse-input $input)
           collapse-mappings
           (sort-by first)
           (take 10)))
