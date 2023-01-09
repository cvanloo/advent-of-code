(ns part2)

(import 'java.security.MessageDigest
        'java.math.BigInteger)

(defn md5 [^String s]
  (let [algorithm (MessageDigest/getInstance "MD5")
        raw (.digest algorithm (.getBytes s))]
    (format "%032x" (BigInteger. 1 raw))))

(def input "ckczppom")

(loop [n 0]
  (let [res (->> (md5 (str input n))
                 (take 6)
                 (distinct))]
    (if (and (= 1 (count res)) (= (first res) \0))
      n
      (recur (inc n)))))

