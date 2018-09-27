import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.List;

class ArmstrongNumbers {

    boolean isArmstrongNumber(int numberToCheck) {
        Collection<Integer> digits = getDigits(numberToCheck);
        int armstrongValue = calculateArmstrongValue(digits);
        return numberToCheck == armstrongValue;
    }

    private Collection<Integer> getDigits(int number) {
        final List<Integer> digits = new ArrayList<>();

        while (number > 0) {
            int digit = number % 10;
            number /= 10;
            digits.add(digit);
        }

        Collections.reverse(digits);

        return digits;
    }

    private int calculateArmstrongValue(Collection<Integer> digits) {
        int armstrongValue = 0;
        int numDigits = digits.size();

        for (Integer digit : digits) {
            armstrongValue += (int) Math.pow(digit, numDigits);
        }

        return armstrongValue;
    }
}
