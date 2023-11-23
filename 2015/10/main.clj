(def input "3113322113")

(defn parse-string
  [s]
  (map #(Integer/parseInt (str %)) s))

(defn next-cypher-step
  [{last :last count :count say :say}]
  (conj say count last))

(defn look-and-say
  [input]
  (next-cypher-step
    (reduce
      (fn [{last :last count :count say :say :as acc} n]
        {:last n
         :count (if (= last n) (inc count) 1)
         :say (if (= last n) say (next-cypher-step acc))})
      {:last (first input) :count 1 :say []}
      (rest input))))

(defmacro repeat-fun
  [n f a]
  `(loop [i# ~n
          a# ~a]
     (if (zero? i#)
       a#
       (recur (dec i#) (~f a#)))))

(macroexpand-1 '(repeat-fun 40 look-and-say (parse-string input)))

(time (count (repeat-fun 40 look-and-say (parse-string input))))
; part 1 (40x): 329356  "Elapsed time: 210.1865 msecs"
(time (count (repeat-fun 50 look-and-say (parse-string input))))
; part 2 (50x): 4666278 "Elapsed time: 2979.234 msecs"
