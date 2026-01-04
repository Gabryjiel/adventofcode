import eu.gabryjiel.aoc2025.Day1
import eu.gabryjiel.aoclib.LoadFile

class Day1Test extends munit.FunSuite {
  lazy val input = LoadFile.loadResource("eu/gabryjiel/aoc2025/day1.txt")

  test("part 1") {
    val result = Day1.Part1(input)
    assertEquals(result, 3)
  }

  test("part 2") {
    val result = Day1.Part2(input)
    assertEquals(result, 6)
  }
}

