(ns part1
  (:require [clojure.string :as str]
            [clojure.stacktrace]))

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
  (let [parts (re-find #"Operation: new = old (.) (\d+|\w+)" monkey-line)
        op (second parts)
        amount (last parts)]
    {:op op :amount amount}))

(defn parse-items
  [monkey-line]
  (letfn [(parseInts [coll] (map #(Integer/parseInt %) coll))]
    (-> monkey-line
        (str/replace "  Starting items: " "")
        (str/split #", ")
        (parseInts)
        vec)))

;; TODO: Trampoline or similar?
(defn parse-monkey
  [monkey-str]
  (let [monkey-str (str/split-lines monkey-str)
        items (parse-items (nth monkey-str 1))
        operation (parse-operation (nth monkey-str 2))
        test-num (parse-test-num (nth monkey-str 3))
        true-monkey (parse-true-monkey (nth monkey-str 4))
        false-monkey (parse-false-monkey (nth monkey-str 5))]
    {:items items :operation operation :test-num test-num :true-monkey true-monkey :false-monkey false-monkey :inspects 0}))

(map parse-monkey (str/split (slurp "test.txt") #"\n\n"))

(def operation-map {"*" *, "+" +})

(defn do-op
  [item {:keys [op amount]}]
  (let [amount (try
                 (Integer/parseInt amount)
                 (catch NumberFormatException _ item))]
    ((get operation-map op) item amount)))

(defn calculate-worry-level
  [item operation]
  (-> item
      (do-op operation)
      (/ 3)
      (Math/floor)))

(defn test-monkey
  [new-item test-num true-monkey false-monkey]
  (if (= 0.0 (mod new-item test-num))
    true-monkey
    false-monkey))

(defn act-turn
  [monkeys monkey-idx]
  (let [{:keys [items operation test-num true-monkey false-monkey]} (nth monkeys monkey-idx)
        new-item (calculate-worry-level (first items) operation)
        to-monkey (test-monkey new-item test-num true-monkey false-monkey)]
    (-> monkeys
        (update-in [to-monkey :items] #(conj % new-item))
        (update-in [monkey-idx :items] #(vec (rest %)))
        (update-in [monkey-idx :inspects] inc))))

(defn do-turns
  [monkeys monkey-idx]
  (let [monkey (get monkeys monkey-idx)]
    (loop [monkeys monkeys items (:items monkey)]
      (if (nil? (first items))
        monkeys
        ; TODO: We don't really need the items here...
        (recur (act-turn monkeys monkey-idx) (rest items))))))

(defn do-round
  [monkeys]
  (loop [monkeys monkeys monkey-idx 0]
    (if (= (count monkeys) monkey-idx)
      monkeys
      (recur (do-turns monkeys monkey-idx) (inc monkey-idx)))))

(defn play-stay-away
  [monkeys rounds]
  (loop [monkeys monkeys round 0]
    (if (= rounds round)
      monkeys
      (recur (do-round monkeys) (inc round)))))

(try
  (play-stay-away (vec (map parse-monkey (str/split (slurp "test.txt") #"\n\n"))) 20)
  (catch Exception e (clojure.stacktrace/print-stack-trace e)))

(let [final-monkeys (play-stay-away (vec (map parse-monkey (str/split (slurp "input.txt") #"\n\n"))) 20)]
  (->> final-monkeys
       (map #(:inspects %))
       (sort >)
       (take 2)
       (reduce *)))

;; Result: 119715
