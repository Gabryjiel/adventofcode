import eu.gabryjiel.aoc2025.Day10.{Part1, Part2}
import eu.gabryjiel.aoclib.LoadFile

class Day10Test extends munit.FunSuite {
  lazy val input = LoadFile.loadResource("eu/gabryjiel/aoc2025/day10.txt")

  test("part 1") {
    val result = Part1(input)
    assertEquals(result, 7L)
  }

  test("part 2") {
    val result = Part2(input)
    assertEquals(result, 33L)
  }
}

