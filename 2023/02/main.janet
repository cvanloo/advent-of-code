(def
  [nred ngreen nblue]
  [12   13     14])

(def parser
  (peg/compile '{:value (* (number :d+) :s+ (<- :w+))
                 :set (+ (* :value :s* "," :s+ :set) :value)
                 :listing (+ (* (group :set) :s* ";" :s+ :listing) (group :set))
                 :game (* :s* "Game" :s+ (number :d+) :s* ":" :s+ (group :listing))
                 :games (+ (* (group :game) "\n" :games) (group :game) (? "\n"))
                 :main (any :games)}))

(def get-game-number (partial 0))
(def get-game-listing (partial 1))

(defn get-color
  [color lst]
  (first (filter
           (fn [[n c]] (when (= c color) n))
           (partition 2 lst))))

(def get-red (partial get-color "red"))
(def get-green (partial get-color "green"))
(def get-blue (partial get-color "blue"))

(defn get-amount
  [color]
  (if (nil? color) 0 (0 color)))

(defn listing-valid?
  [lst]
  (and (<= (get-amount (get-red lst)) nred)
       (<= (get-amount (get-green lst)) ngreen)
       (<= (get-amount (get-blue lst)) nblue)))

(defn valid-game?
  [game]
  (all listing-valid? (get-game-listing game)))

(->> (slurp "input.txt")
     (peg/match parser)
     (filter valid-game?)
     (map get-game-number)
     (reduce + 0))

# => 2617
