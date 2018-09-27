class Hamming {

    companion object {
        fun compute(firstSequence: String, secondSequence: String): Int {
            require(firstSequence.length == secondSequence.length) { "left and right strands must be of equal length." }
            return firstSequence.indices.count { firstSequence[it] != secondSequence[it] }
        }
    }
}