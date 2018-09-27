import java.util.HashMap;
import java.util.Map;

class RnaTranscription {

    private static final Map<Character, Character> rnaNucleotidesByDNANucleotide = new HashMap<Character, Character>(){{
        put('G', 'C');
        put('C', 'G');
        put('T', 'A');
        put('A', 'U');
    }};

    String transcribe(String dnaStrand) {
        final StringBuilder sb = new StringBuilder();
        dnaStrand.codePoints().forEach(n -> sb.append(rnaNucleotidesByDNANucleotide.get((char) n)));
        return sb.toString();
    }

}
