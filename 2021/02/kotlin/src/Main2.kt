package main2

import java.io.File

fun main(args: Array<String>) {
    val pos = Position(depth = 0, horiz = 0, aim = 0)

    File("/home/miya/.tmp/day2test/src/input.txt").forEachLine {
        val strings = it.split(" ")

        val dir = strings[0]
        val amt = strings[1].toInt()

        with (pos) {
            when (strings[0]) {
                "forward" -> {
                    horiz += amt
                    depth += aim * amt
                }
                "up" -> aim -= amt
                "down" -> aim += amt
            }
        }
    }

    println("Result: " + pos.depth * pos.horiz)
}

data class Position(var depth: Int, var horiz: Int, var aim: Int)
