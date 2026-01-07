import eu.gabryjiel.aoc2025.Day5.{Part1, Part2}
import eu.gabryjiel.aoclib.LoadFile

class Day5Test extends munit.FunSuite {
  lazy val input = LoadFile.loadResource("eu/gabryjiel/aoc2025/day5.txt")

  test("part 1") {
    val result = Part1(input)
    assertEquals(result, 3L)
  }

  test("part 2") {
    val result = Part2(input)
    assertEquals(result, 14L)
  }
}

