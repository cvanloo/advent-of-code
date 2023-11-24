(def inc-char (comp char inc int))

(defn wrap-char [c]
  (if (> (int c) (int \z)) \a c))

(defn inc-str
  ([s] (inc-str s '()))
  ([s front]
   (let [sl (last s)
         sr (butlast s)
         nc (wrap-char (inc-char sl))]
     (if (= nc \a)
       (recur sr (conj front nc))
       (concat sr (conj front nc))))))

(defn straight? [s]
  (not (empty? (filter (fn [[a b c]]
                         (and (= 1 (- c b)) (= 1 (- b a))))
                       (partition 3 1 (map int s))))))

(defn readable? [s]
  (empty? (filter (fn [c] (some #(= c %) [\i \o \l])) s)))

(defn pairs? [s]
  (<= 2 (count (distinct (filter (fn [[a b]]
                                  (= a b))
                                (partition 2 1 s))))))

(defn inc-seq [s]
  (lazy-seq (cons s (inc-seq (inc-str s)))))

(def policy? (every-pred straight? readable? pairs?))

(time (apply str (first (filter policy? (inc-seq "hepxcrrq")))))
; => part 1: hepxxyzz "Elapsed time: 3124.9757 msecs"
(time (apply str (first (filter policy? (inc-seq (inc-str "hepxxyzz"))))))
; => part 2: heqaabcc "Elapsed time: 8045.0703 msecs"