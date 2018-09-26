import java.lang.IllegalArgumentException

class Hamming {

    companion object {
        fun compute(firstSequence: String, secondSequence: String): Int {
            return when {
                firstSequence.length != secondSequence.length -> throw IllegalArgumentException("left and right strands must be of equal length.")
                else -> {
                    var difference = 0

                    for (i in 0 until firstSequence.length) {
                        if (firstSequence[i] != secondSequence[i]) {
                            difference++
                        }
                    }

                    difference
                }
            }
        }
    }
}