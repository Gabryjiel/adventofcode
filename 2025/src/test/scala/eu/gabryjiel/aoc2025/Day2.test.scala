import eu.gabryjiel.aoc2025.Day2
import eu.gabryjiel.aoclib.LoadFile

class Day2Test extends munit.FunSuite {
  lazy val input = LoadFile.loadResource("eu/gabryjiel/aoc2025/day2.txt")

  test("part 1") {
    val result = Day2.Part1(input.head)
    assertEquals(result, 1227775554L)
  }

  test("part 2") {
    val result = Day2.Part2(input.head)
    assertEquals(result, 4174379265L)
  }
}

