(ns solutionb)

(import 'java.security.MessageDigest
        'java.math.BigInteger)

(defn md5
  [s]
  (->> (.getBytes s)
       (.digest (MessageDigest/getInstance "MD5"))
       (BigInteger. 1)
       (format "%032x")))

(defn validate-hash
  [zeroes input]
  (if (= \0 (first input))
    (->> input
         (take zeroes)
         (distinct)
         count
         (= 1))
    false))

(defn find-hash
  [zeroes input]
  (->> (iterate inc 0)
       (filter #(->> (str input %) md5 (validate-hash zeroes)))
       first))

(find-hash 5 "ckczppom") ; => 117946
(find-hash 6 "ckczppom") ; => 3938038
