import eu.gabryjiel.aoc2025.Day8.{Part1, Part2}
import eu.gabryjiel.aoclib.LoadFile

class Day8Test extends munit.FunSuite {
  lazy val input = LoadFile.loadResource("eu/gabryjiel/aoc2025/day8.txt")

  test("part 1") {
    val result = Part1(input, 10)
    assertEquals(result, 40L)
  }

  test("part 2") {
    val result = Part2(input)
    assertEquals(result, 25272L)
  }
}

