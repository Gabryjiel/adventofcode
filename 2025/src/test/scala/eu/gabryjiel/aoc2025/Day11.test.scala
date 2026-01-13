import eu.gabryjiel.aoc2025.Day11.{Part1, Part2}
import eu.gabryjiel.aoclib.LoadFile

class Day11Test extends munit.FunSuite {
  lazy val input = LoadFile.loadResource("eu/gabryjiel/aoc2025/day11.txt")
  lazy val input2 = LoadFile.loadResource("eu/gabryjiel/aoc2025/day11p2.txt")

  test("part 1") {
    val result = Part1(input)
    assertEquals(result, 5L)
  }

  test("part 2") {
    val result = Part2(input2)
    assertEquals(result, 2L)
  }
}

