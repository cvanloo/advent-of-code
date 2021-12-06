package main2

import java.io.File

fun main(args: Array<String>) {
    val pos = Position(depth = 0, horiz = 0, aim = 0)

    File("/home/miya/code/kotlin-intro/src/input.txt").forEachLine {
        val strings = it.split(" ")

        val dir: Direction = when (strings[0]) {
            "forward" -> Direction.Forward
            "up" -> Direction.Up
            "down" -> Direction.Down
            else -> return@forEachLine;
        }

        val amt = strings[1].toInt()

        val command = Command(direction = dir, amount = amt)
        pos.handle(command)
    }

    println("Result: " + pos.depth * pos.horiz)
}

enum class Direction {
    Forward, Up, Down
}

data class Command(val direction: Direction, val amount: Int)

data class Position(var depth: Int, var horiz: Int, var aim: Int) {
    fun handle(cmd: Command) {
        when (cmd.direction) {
            Direction.Forward -> {
                horiz += cmd.amount
                depth += aim * cmd.amount
            }
            Direction.Down -> aim += cmd.amount
            Direction.Up -> aim -= cmd.amount
        }
    }
}