import kotlin.test.Test
import kotlin.test.assertEquals

class SolutionTest {
    private val target = Solution()

    @Test
    fun testFizzBuzz_3() {
        val expected = listOf("1", "2", "Fizz")
        assertEquals(expected, target.fizzBuzz(3))
    }

    @Test
    fun testFizzBuzz_5() {
        val expected = listOf("1", "2", "Fizz", "4", "Buzz")
        assertEquals(expected, target.fizzBuzz(5))
    }

    @Test
    fun testFizzBuzz_15() {
        val expected = listOf(
            "1",
            "2",
            "Fizz",
            "4",
            "Buzz",
            "Fizz",
            "7",
            "8",
            "Fizz",
            "Buzz",
            "11",
            "Fizz",
            "13",
            "14",
            "FizzBuzz",
        )
        assertEquals(expected, target.fizzBuzz(15))
    }
}