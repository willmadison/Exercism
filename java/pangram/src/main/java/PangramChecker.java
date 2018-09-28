import java.util.HashSet;
import java.util.Set;

public class PangramChecker {

    public static final int ALPHABET_LENGTH = 26;

    public boolean isPangram(String input) {
        Set<Character> seen = new HashSet<>();

        input.toLowerCase().codePoints()
                .filter(c -> Character.isLetter((char) c))
                .forEach(c -> seen.add((char) c));

        return seen.size() == ALPHABET_LENGTH;
    }

}
