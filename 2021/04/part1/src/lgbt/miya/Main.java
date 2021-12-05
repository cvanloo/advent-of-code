package lgbt.miya;

import java.io.*;
import java.util.ArrayList;
import java.util.List;

/**
 * Advent of Code 2021, Day 4, Part 1
 */
public class Main {

    /**
     * Parse the input.
     * The file format has to obey to the following rules:
     * <ul>
     *  <li>The first line contains the called numbers, separated by
     *  commas.</li>
     *  <li>Followed by an empty line.</li>
     *  <li>Followed by all the bingo boards.</li>
     *  <li>A bingo board consists of 5 lines, each line consisting of 5
     *   numbers, separated by spaces.</li>
     *  <li>After each bingo board, an empty line must follow, including the
     *   very last board.</li>
     * </ul>
     * @param inputPath Path to the input file.
     * @param boards The List to hold the boards.
     * @return The called numbers.
     * @throws IOException An error occurred when reading the file.
     * @throws FileNotFoundException The file could not be found.
     */
    public static int[] parseInput(String inputPath, List<Board> boards)
            throws IOException, FileNotFoundException
    {
        File file = new File(inputPath);
        BufferedReader input = new BufferedReader(new FileReader(file));

        // First line contains the called numbers separated by commas.
        String numbersString = input.readLine();
        String[] numbersStringArray = numbersString.split(",");
        int[] calledNumbers = new int[numbersStringArray.length];

        for (int i = 0; i < numbersStringArray.length; i++) {
            calledNumbers[i] = Integer.parseInt(numbersStringArray[i]);
        }

        // Empty line after called numbers.
        input.readLine();

        // All other lines represent the bingo boards.
        // NOTE(miya): The input file __must__ end with an empty line. In
        // JetBrains IDEs make sure that __two__ empty lines are displayed,
        // because JetBrains is a _liar_ and one of the two empty lines is
        // a fake.
        String boardString = "";
        String line;

        while ((line = input.readLine()) != null) {
            if (!line.isEmpty()) {
                boardString = boardString.concat(line.concat("\n"));
            } else {
                boards.add(Board.FromString(boardString));
                boardString = "";
            }
        }

        return calledNumbers;
    }

    /**
     * Main entry point.
     * @param args Command line arguments.
     */
    public static void main(String[] args) {

        // NOTE(miya): Before you run this program, make sure to replace the
        // path below in the first argument to `parseInput` with your own.
        // ___Don't forget that the input file has to end in an empty line!___
        // Read the documentation for the `parseInput` function.

        List<Board> boards = new ArrayList<>();
        int[] calledNumbers;

        try {
            calledNumbers = parseInput(
                "/home/miya/code/aoc/2021/04/part1/src/lgbt/miya/input.txt", boards);
        } catch (Exception exception) {
            exception.printStackTrace();
            return;
        }

        Board winner = null;
        int lastNumber = 0;

        // Iterate over all called numbers and pass them to the boards.
break_out:
        for (int calledNumber : calledNumbers) {
            for (Board board : boards) {
                boolean isBingo = board.callNumber(calledNumber);
                if (isBingo) {
                    winner = board;
                    lastNumber = calledNumber;
                    break break_out; // Label used to break out of both loops.
                }
            }
        }

        if (null == winner) {
            System.out.println("Not a single winner found.");
            return;
        }

        // The result is calculated by multiplying the sum of all unmarked
        // numbers with the last called number.
        int[] unmarkedNumbers = winner.getUnmarkedNumbers();
        int result = 0;
        for (int number : unmarkedNumbers) {
            result += number;
        }
        result *= lastNumber;

        System.out.println("The result is: " + result);
    }
}
