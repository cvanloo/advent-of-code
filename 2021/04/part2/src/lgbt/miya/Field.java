package lgbt.miya;

/**
 * Represents a single Bingo board field.
 */
public class Field {

    // --- Member/Fields --- //

    private final int number;
    private boolean checked;

    // --- Methods/Behavior --- //

    /**
     * Constructor.
     * @param number The number the field holds.
     */
    public Field(int number) {
        this.number = number;
        this.checked = false;
    }

    /**
     * Mark this field as checked.
     */
    public void check() {
        this.checked = true;
    }

    /**
     * Get the number that this field holds.
     * @return The hold number.
     */
    public int getNumber() {
        return this.number;
    }

    /**
     * Get the checked state of this field.
     * @return `True` if the field is checked, `False` if it is not.
     */
    public boolean getChecked() {
        return this.checked;
    }
}
