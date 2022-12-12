(ns part1
  (:require (clojure [string :as str]
                     [repl :as repl])))

(defn apply-instruction
  [stacks instruction]
  (let [amount (:amount instruction)
        from (:from instruction)
        to (:to instruction)]
    (->> (take amount (stacks from))
         (apply conj (stacks to))
         (assoc stacks to))))

(defn remove-old
  [stacks instruction]
  (let [amount (:amount instruction)
        from (:from instruction)]
    (->> (drop amount (stacks from))
         (assoc stacks from))))

(defn run
  [instructions stacks]
  (loop [ins instructions
         stacks stacks]
    (if (empty? ins)
      stacks
      (recur (rest ins)
             (remove-old (apply-instruction stacks (first ins)) (first ins))))))

;
; Parsing stuff
;

(defn get-stack-line-elements
  [line-coll]
  (reduce #(conj %1 (nth %2 1)) [] line-coll))

(defn create-stacks
  [stack-strs]
  (partition (count stack-strs)
             (apply interleave stack-strs)))

(defn remove-spaces
  [stack]
  (filter #(not (str/blank? (str %))) stack))

(defn parse-stack
  [stack-strs]
  (->> stack-strs
         (str/split-lines)
         (butlast)
         (map #(partition 4 4 '(\space) %))
         (map get-stack-line-elements)
         (create-stacks)
         (map remove-spaces)
         (zipmap (iterate inc 1))))

(defn parse-instruction
  [instruction-str]
  (zipmap '(:amount :from :to)
          (map #(Integer/parseInt %) (map last (partition 2 (str/split instruction-str #" "))))))

(defn parse-instructions
  [instruction-strs]
  (->> instruction-strs
       (str/split-lines)
       (map parse-instruction)))

(defn main
  [filename]
  (let [[stack-strs instruction-strs] (str/split (slurp filename) #"\n\n")
        stacks (parse-stack stack-strs)
        instructions (vec (parse-instructions instruction-strs))]
    (->> (run instructions stacks)
         (into (sorted-map))
         (vals)
         (map first)
         (apply str))))

(main "test.txt") ; "CMZ"
(main "input.txt") ; "CNSZFDVLJ"
