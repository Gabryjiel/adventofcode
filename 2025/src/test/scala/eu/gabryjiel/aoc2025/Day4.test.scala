import eu.gabryjiel.aoc2025.Day4
import eu.gabryjiel.aoclib.LoadFile

class Day4Test extends munit.FunSuite {
  lazy val input = LoadFile.loadResource("eu/gabryjiel/aoc2025/day4.txt")

  test("part 1") {
    val result = Day4.Part1(input)
    assertEquals(result, 13)
  }

  test("part 2") {
    val result = Day4.Part2(input)
    assertEquals(result, 43)
  }
}

