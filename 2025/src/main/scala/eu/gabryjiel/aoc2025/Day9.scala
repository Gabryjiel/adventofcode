package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile.loadResource
import eu.gabryjiel.aoclib.Point
import eu.gabryjiel.aoclib.Rectangle

object Day9 {
  def main(args: Array[String]): Unit = {
    lazy val input = loadResource("eu/gabryjiel/aoc2025/day9.txt")
    println(s"Day9.Part1: ${this.Part1(input)}")
    println(s"Day9.Part2: ${this.Part2(input)}")
  }

  def Part1(lines: Array[String]): Long =
    val points = lines.map({ case s"$x,$y" =>
      Point(x.toLong, y.toLong)
    })

    var max = 0L

    for x <- 0.until(points.length - 1) do
      for y <- (x + 1).until(points.length) do
        max = math.max(max, Rectangle(points(x), points(y)).area)

    max

  def Part2(lines: Array[String]): Long =
    
    0L
}
