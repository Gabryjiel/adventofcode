package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile.loadResource

object Day3 {
  def main(args: Array[String]): Unit = {
    lazy val input = loadResource("eu/gabryjiel/aoc2025/day3.txt")
    println(s"Day3.Part1: ${this.Part1(input)}")
    println(s"Day3.Part2: ${this.Part2(input)}")
  }

  def Part1(lines: Array[String]): Int = {
    lines
      .map(line => {
        val segment = line.slice(0, line.length() - 1)
        val max = segment.max
        val index = segment.indexOf(max)

        val secondMax = line.slice(index + 1, line.length()).max

        s"$max$secondMax".toInt
      })
      .sum
  }

  def Part2(lines: Array[String]): Long = {
    lines
      .map(line => {
        var startIndex = 0;

        val digits = 11.to(0, -1).map(end => {
          val segment = line.slice(startIndex, line.length() - end)
          val max = segment.max
          val index = segment.indexOf(max)

          startIndex += index + 1
          max
        }).mkString

        digits.toLong
      })
      .sum
  }
}
