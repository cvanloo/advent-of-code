(ns part1
  (:require [clojure.string :as str]
            [clojure.stacktrace]))

(defn find-start-pos
  [fields]
  (let [max-y (count fields)
        max-x (count (first fields))]
    (loop [x 0 y 0]
      (cond
        (>= x max-x)
        (recur 0 (inc y))
        (>= y max-y)
        nil
        (= \S (nth (nth fields y) x))
        [x y]
        :else (recur (inc x) y)))))

;(find-start-pos (map seq (str/split-lines (slurp "test.txt"))))
;(find-start-pos (map seq (str/split-lines (slurp "input.txt"))))

(defn index
  [fields [x y]]
  (nth (nth fields y nil) x nil))

(defn get-neighbours
  [fields [posx posy]]
  (for [[x y] '([0 -1] [0 +1] [-1 0] [+1 0])]
    (let [x (+ posx x)
          y (+ posy y)
          val (index fields [x y])]
      {:pos [x y] :val val})))

;(neighbours (map seq (str/split-lines (slurp "test.txt"))) [3 3])
;(neighbours (map seq (str/split-lines (slurp "test.txt"))) [7 5])

(defn height-difference
  [h1 h2]
  (let [h1 (int h1)
        h2 (int h2)]
    (Math/abs (- h2 h1))))

;(height-difference (map seq (str/split-lines (slurp "test.txt"))) [0 0] [0 1])
;(height-difference (map seq (str/split-lines (slurp "test.txt"))) [0 1] [0 0])

(defn path-contains?
  [path pos]
  (some #(= (:pos pos) (:pos %)) path))

; (nil? (path-contains? '({:pos [0 1] :val \a} {:pos [0 2] :val \b}) {:pos [0 3] :val \c}))

(defn find-allowed-fields
  [fields path pos]
  (let [ns (get-neighbours fields pos)]
    (->> ns
         (filter #(not (nil? (:val %))))
         (filter #(>= 1 (height-difference \a (:val %))))
         (filter #(nil? (path-contains? path %))))))

;(find-allowed-fields (map seq (str/split-lines (slurp "test.txt"))) '() [0 0])
;(find-allowed-fields (map seq (str/split-lines (slurp "test.txt"))) '() [2 4])
;(find-allowed-fields (map seq (str/split-lines (slurp "test.txt"))) '({:pos [0 1], :val \a}) [0 0])

(defn follow-path
  "Follows down a path, taking the path walked so far as input.
   Returns the final path, either ending with \\E or nil."
  [fields path]
  (let [current (first path)
        neighbours (find-allowed-fields fields path current)]
    (if (nil? current)
      path
      (follow-path fields (conj path (first neighbours))))))

(try
 (let [fields (map seq (str/split-lines (slurp "test.txt")))
       start (find-start-pos fields)]
  (follow-path fields (list start)))
 (catch Exception e (clojure.stacktrace/print-stack-trace e)))

(defn find-path
  [fields start]
  ())

(let [fields (map seq (str/split-lines (slurp "test.txt")))
      start (find-start-pos fields)]
  (find-path fields start))
