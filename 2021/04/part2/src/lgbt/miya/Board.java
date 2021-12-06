package lgbt.miya;

import java.util.ArrayList;
import java.util.List;

/**
 * Board represents a Bingo Board.
 */
public class Board {

    // --- Member/Fields --- //

    private final static int WIDTH = 5;
    private final static int HEIGHT = 5;

    private final Field[] fields;

    private boolean hasBingo = false;
    private int numberAtBingo;

    // --- Methods/Behavior --- //

    /**
     * Costructor.
     */
    public Board(Field[] fields) {
        this.fields = fields;
    }

    /**
     * Create a new `Board` from a string.
     * @param boardString A string containing the board. Fields have to be
     *                    separated by _spaces_, and rows separated by
     *                    _newlines_ ("\n").
     * @return Returns a new `Board`.
     */
    public static Board FromString(String boardString) {
        Field[] fields = new Field[WIDTH*HEIGHT];
        String[] rows = boardString.split("\n");

        // x -> row
        // y -> column
        // Length: WIDTH*HEIGHT
        // Index:  WIDTH*y+x
        for (int y = 0; y < HEIGHT; y++) {

            // Strip leading and trailing whitespaces, otherwise we would have
            // empty "" strings at the start of our array, which cannot be
            // parsed to integers.
            rows[y] = rows[y].trim();

            // Split at whitespace characters.
            String[] stringFields = rows[y].split("\\s+");

            for (int x = 0; x < WIDTH; x++) {
                int number = Integer.parseInt(stringFields[x]);
                fields[WIDTH*y+x] = new Field(number);
            }
        }

        return new Board(fields);
    }

    /**
     * Tell the board what number was called.
     * @param number The called number.
     * @return `True` if the board won, `False` otherwise.
     */
    public boolean callNumber(int number) {
        for (Field field : this.fields) {
            if (field.getNumber() == number) {
                field.check();
            }
        }

        /* When each field in a _row_ or _column_ is checked, we have a bingo. */

        /* Check rows.
         * We iterate through all fields of a row, if a field is not checked,
         * we know its not possible that this row is a winner, so we skip
         * (break) to the next row.
         */
        for (int y = 0; y < HEIGHT; y++) {
            for (int x = 0; x < WIDTH; x++) {
                Field currentField = this.fields[WIDTH*y+x];

                if (!currentField.getChecked()) break;
                if (x == WIDTH-1) {
                    this.hasBingo = true;
                    this.numberAtBingo = number;
                    return true;
                }
            }
        }

        /* Check columns.
         * We iterate through all fields of a column, if a field is not
         * checked, we know its not possible that this column is a winner, so
         * we skip (break) to the next column.
         */
        for (int x = 0; x < HEIGHT; x++) {
            for (int y = 0; y < WIDTH; y++) {
                Field currentField = this.fields[WIDTH*y+x];

                if(!currentField.getChecked()) break;
                if (y == HEIGHT-1) {
                    this.hasBingo = true;
                    this.numberAtBingo = number;
                    return true;
                }
            }
        }

        return false;
    }

    /**
     * Get all of the boards unmarked numbers.
     * @return The unmarked numbers.
     */
    public int[] getUnmarkedNumbers() {

        // Get all unmarked fields.
        List<Field> fields = new ArrayList<>();

        for (Field field : this.fields) {
            if (!field.getChecked()) fields.add(field);
        }

        // Get the numbers of the unmarked fields.
        int[] unmarkedFields = new int[fields.size()];
        int iterator = 0;

        for (Field field : fields) {
            unmarkedFields[iterator++] = field.getNumber();
        }

        return unmarkedFields;
    }

    public boolean getHasBingo() {
        return this.hasBingo;
    }

    public int getNumberAtBingo() {
        return this.numberAtBingo;
    }
}
