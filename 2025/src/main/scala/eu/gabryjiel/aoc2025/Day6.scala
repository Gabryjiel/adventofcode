package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile.loadResource
import eu.gabryjiel.aoclib.Interval
import scala.collection.mutable.HashMap
import scala.collection.mutable.HashSet

object Day6 {
  def main(args: Array[String]): Unit = {
    lazy val input = loadResource("eu/gabryjiel/aoc2025/day6.txt")
    println(s"Day6.Part1: ${this.Part1(input)}")
    println(s"Day6.Part2: ${this.Part2(input)}")
  }

  def collectNumbers(list: List[Long], symbol: String): Long = symbol match {
    case "*" => list.product
    case "+" => list.sum
  }

  def Part1(lines: Array[String]): Long = {
    lines
      .map(x => x.split(" ").filter(_ != ""))
      .transpose
      .map(col => {
        val numbers = col.dropRight(1).map(_.toLong).toList
        collectNumbers(numbers, col.last)
      })
      .sum
  }

  def Part2(lines: Array[String]): Long = {
    val grid = lines.map(_.split(""))
    val symbols = grid.last.map(_.trim).filter(_ != "")
    var symbolIndex = 0

    grid
      .dropRight(1)
      .transpose
      .map(f => f.reduce(_ + _).trim)
      .foldLeft(List(List.empty[String])) {
        case (acc @ head :: tail, "")  => List.empty[String] :: head :: tail
        case (acc @ head :: tail, cur) => (head :+ cur) :: tail
        case (acc, cur)                => acc
      }
      .reverse
      .zipWithIndex
      .map(f => collectNumbers(f(0).map(_.toLong), symbols(f(1))))
      .sum
  }
}
