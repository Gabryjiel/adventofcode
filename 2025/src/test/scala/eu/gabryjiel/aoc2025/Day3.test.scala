import eu.gabryjiel.aoc2025.Day3
import eu.gabryjiel.aoclib.LoadFile

class Day3Test extends munit.FunSuite {
  lazy val input = LoadFile.loadResource("eu/gabryjiel/aoc2025/day3.txt")

  test("part 1") {
    val result = Day3.Part1(input)
    assertEquals(result, 357)
  }

  test("part 2") {
    val result = Day3.Part2(input)
    assertEquals(result, 3121910778619L)
  }
}

