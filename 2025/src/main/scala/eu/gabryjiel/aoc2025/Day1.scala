package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile

object Day1 {
  def main(args: Array[String]): Unit = {
    lazy val input = LoadFile.loadResource("eu/gabryjiel/aoc2025/day1.txt")

    println(s"Day1.Part1: ${this.Part1(input)}")
    println(s"Day1.Part2: ${this.Part2(input)}")
  }

  def parseRotation(s: String): Int = s match {
    case s"L$i" => -i.toInt
    case s"R$i" => i.toInt
  }

  def Part1(input: Array[String]): Int = {
    input.map(parseRotation).scanLeft[Int](50)((acc, cur) => acc + cur).count(_ % 100 == 0)
  }

  def Part2(input: Array[String]): Int = {
    input.map(parseRotation).foldLeft((50, 0))((acc, cur) => {
      val currentPosition = acc(0)
      val nextPosition = currentPosition + cur

      val hundred1 = math.abs(currentPosition / 100)
      val hundred2 = math.abs(nextPosition / 100)
      val hunderedAddend = if currentPosition.sign != nextPosition.sign then hundred1 + 1 + hundred2 else math.abs(hundred1 - hundred2)

      val a = nextPosition % 100
      val nextPositionCycled = if a < 0 then 100 + a else a 
      val zeroAddend = if a == 0 then -1 else 0

      (nextPositionCycled, acc(1) + hunderedAddend + zeroAddend)
    })(1)
  }
}
