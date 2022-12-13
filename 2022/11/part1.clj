(ns part1
  (:require [clojure.string :as str]
            [clojure.repl :as repl]))

(defn parse-false-monkey
  [monkey-line]
  (Integer/parseInt (str/replace monkey-line "    If false: throw to monkey " "")))

(defn parse-true-monkey
  [monkey-line]
  (Integer/parseInt (str/replace monkey-line "    If true: throw to monkey " "")))

(defn parse-test-num
  [monkey-line]
  (Integer/parseInt (str/replace monkey-line "  Test: divisible by " "")))

(defn parse-operation
  [monkey-line]
  (let [parts (re-find #"Operation: new = old (.) (\d|\w)" monkey-line)
        op (second parts)
        amount (last parts)]
    {:op op :amount amount}))

(defn parse-items
  [monkey-line]
  (letfn [(parseInts [coll] (map #(Integer/parseInt %) coll))]
    (-> monkey-line
      (str/replace "  Starting items: " "")
      (str/split #", ")
      (parseInts))))

;; TODO: Trampoline or similar?
(defn parse-monkey
  [monkey-str]
  (let [monkey-str (str/split-lines monkey-str)
        items (parse-items (nth monkey-str 1))
        operation (parse-operation (nth monkey-str 2))
        test-num (parse-test-num (nth monkey-str 3))
        true-monkey (parse-true-monkey (nth monkey-str 4))
        false-monkey (parse-false-monkey (nth monkey-str 5))]
    {:items items :operation operation :test-num test-num :true-monkey true-monkey :false-monkey false-monkey}))

(def monkeys (map parse-monkey (str/split (slurp "test.txt") #"\n\n")))

(def operation-map {"*" *, "+" +})

(defn move-to-monkey
  [worry-level divisor true-m false-m]
  (if (mod worry-level divisor)
    true
    false))

(defn do-op
  [item {:keys [op num]}]
  (let [num (try
             (Integer/parseInt num)
             (catch Exception _ item))]
    ((get operation-map op) item num)))

(defn calculate-worry-level
  [item operation]
  (-> item
    (do-op operation)
    (/ 3)
    (Math/floor)))

(defn act-turn
  [monkey]
  (let [item (first (:items monkey))
        operation (:operation monkey)
        divisor (:test-num monkey)
        true-m (:true-monkey monkey)
        false-m (:false-monkey monkey)]
    (calculate-worry-level item operation)))

(for [i (range 20)]
  (map act-turn monkeys))
