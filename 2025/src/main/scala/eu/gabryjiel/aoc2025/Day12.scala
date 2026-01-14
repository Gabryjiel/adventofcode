package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile.loadResourceAsString

object Day12 {
  def main(args: Array[String]): Unit = {
    lazy val input = loadResourceAsString("eu/gabryjiel/aoc2025/day12.txt")
    println(s"Day12.Part1: ${this.Part1(input)}")
  }

  def parseShape(str: String): Int =
    str.toCharArray.count(_ == '#')

  def parseRule(str: String) = str match {
    case s"${w}x${y}: $nums" => (w.toInt, y.toInt, nums.split(' ').map(_.toInt))
  }

  // Taken from https://github.com/jackysee/aoc/blob/main/2025/day12.js
  def Part1(input: String): Long =
    val splits = input.split("\n\n")
    val (shapes, rules) = splits.splitAt(splits.length - 1)
    val areas = shapes.map(parseShape)

    rules
      .mkString("")
      .split("\n")
      .filter(t =>
        val (width, height, counts) = parseRule(t)
        val ruleArea = width * height
        val need = counts.zipWithIndex
          .foldLeft(0)((acc, cur) => acc + cur(0) * areas(cur(1)))

        need.toDouble / ruleArea.toDouble <= 0.85
      )
      .length
}
