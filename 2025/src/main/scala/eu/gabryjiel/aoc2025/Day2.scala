package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile

object Day2 {
  def main(args: Array[String]): Unit = {
    lazy val input = LoadFile.loadResource("eu/gabryjiel/aoc2025/day2.txt")

    println(s"Day2.Part1: ${this.Part1(input.head)}")
    println(s"Day2.Part2: ${this.Part2(input.head)}")
  }

  def Part1(line: String): Long = {
    line
      .split(",")
      .foldLeft(0L)((acc, cur) => {
        val (start, end) = parseRange(cur)
        val invalidCount = start.to(end).map(_.toString).filter(x => {
          if x.length() % 2 != 0 then
           false
          else
            val (x1, x2) = x.splitAt(x.length() / 2)
            x1 == x2
        }).map(_.toLong).sum

        acc + invalidCount
      })
  }

  def Part2(line: String): Long = {
    line
      .split(",")
      .foldLeft(0L)((acc, cur) => {
        val (start, end) = parseRange(cur)
        val invalidCount = start.to(end).map(_.toString).filter(x => {
          1.to(x.length() / 2).map(end => x.slice(0, end)).exists(needle => isSubsequenceOf(x, needle))
        }).map(_.toLong).sum

        acc + invalidCount
      })
  }

  def parseRange(rangeStr: String): (Long, Long) = rangeStr match {
      case s"$x1-$x2" => (x1.toLong, x2.toLong)
  }

  def isSubsequenceOf(haystack: String, needle: String): Boolean = {
    haystack.split(needle).length == 0
  }
}
