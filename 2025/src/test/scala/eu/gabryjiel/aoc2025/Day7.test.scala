import eu.gabryjiel.aoc2025.Day7.{Part1, Part2}
import eu.gabryjiel.aoclib.LoadFile

class Day7Test extends munit.FunSuite {
  lazy val input = LoadFile.loadResource("eu/gabryjiel/aoc2025/day7.txt")

  test("part 1") {
    val result = Part1(input)
    assertEquals(result, 21)
  }

  test("part 2") {
    val result = Part2(input)
    assertEquals(result, 40L)
  }
}

