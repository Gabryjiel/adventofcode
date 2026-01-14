import eu.gabryjiel.aoc2025.Day12.Part1
import eu.gabryjiel.aoclib.LoadFile

class Day12Test extends munit.FunSuite {
  lazy val input = LoadFile.loadResourceAsString("eu/gabryjiel/aoc2025/day12.txt")

  test("part 1") {
    val result = Part1(input)
    assertEquals(result, 2L)
  }
}

