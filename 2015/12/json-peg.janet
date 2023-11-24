# https://www.json.org/json-en.html
(def json (peg/compile
            '{:value (+ :object :array :str :num "true" "false" "null")
              :element (* :s* :value :s*)
              :elements (+ (* :element "," :elements) :element)
              :num (* (? "-") :d+) # @todo: fractions, scientific notation
              :str (* "\"" (any (range " !" "#~")) "\"") # @todo: escape characters, hex

              :object (* "{" (+ :members :s*) "}")
              :member (* :s* :str :s* ":" :element)
              :members (+ (* :member "," :members) :member)

              :array (* "[" (+ :elements :s*) "]")

              :main (* :element -1)}))


(defn json-valid? [j] (peg/match json j))

(defn list-same? [as bs] (all (fn [[a b]] (= a b)) (partition 2 (interleave as bs))))

(map json-valid? ["    { \"hello world\"    :    \"\"} 5    "
                  "    { \"hello world\"    :    \"\"}, 5   "
                  "    [{ \"hello world\"    :    \"\"}, 5  ]   "
                  "    { \"hello world\"    :    \"\" 5}    "
                  "    { \"hello world\"    :    \"\", 5 }   "
                  "    {}   "
                  "    []   "
                  "    [] 5  "
                  "    [], 5  "
                  "    [[], 5]  "
                  (slurp "input.txt")])
# => @[nil nil @[] nil nil @[] @[] nil nil @[] @[]]
