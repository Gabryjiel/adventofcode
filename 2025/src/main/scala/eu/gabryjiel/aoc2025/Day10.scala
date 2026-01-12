package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile.loadResource
import eu.gabryjiel.aoclib.Point
import eu.gabryjiel.aoclib.Rectangle
import scala.util.boundary

object Day10 {
  def main(args: Array[String]): Unit = {
    lazy val input = loadResource("eu/gabryjiel/aoc2025/day10.txt")
    println(s"Day10.Part1: ${this.Part1(input)}")
    println(s"Day10.Part2: ${this.Part2(input)}")
  }

  def parse(line: String) = line match {
    case s"[$lights] $buttons {$joltage}" =>
      (parseLights(lights), parseButtons(buttons), parseJoltage(joltage))
  }

  def parseLights(lights: String) =
    lights.toCharArray.map(c => c == '#').toList

  def parseButtons(buttons: String) =
    buttons.trim
      .split(' ')
      .map({ case s"($values)" =>
        values.split(',').map(_.toInt).toList
      })
      .toList

  def parseJoltage(joltage: String) =
    joltage.split(',').map(_.toInt).toList

  def getAllCombinations(list: List[List[Int]]) =
    for {
      len <- 1.to(list.length)
      combinations <- list.combinations(len)
    } yield combinations

  def flickButtons(buttons: List[Int], targetSize: Int): List[Boolean] =
    buttons
      .foldLeft(Array.fill(targetSize)(false))((acc, cur) => {
        acc.update(cur, !acc(cur))
        acc
      })
      .toList

  def Part1(lines: Array[String]): Long =
    lines
      .map(line => {
        val (target, buttons, _) = parse(line)
        val validCombinations = getAllCombinations(buttons)
          .map(t => (t.length, t.flatten))
          .filter(t => flickButtons(t._2, target.length).sameElements(target))

        val bestCombination = validCombinations.sortBy(t => t._1).head
        bestCombination._1.toLong
      })
      .sum

  def Part2(lines: Array[String]): Long =
    lines
      .take(1)
      .map(line => {
        val (_, buttons, joltage) = parse(line)
        val maxCombination =
          buttons.flatMap(t => List.fill(t.map(b => joltage(b)).min)(t))
        var result = maxCombination.length

        boundary {
          for i <- 1.to(maxCombination.length) do
            val a = maxCombination
              .combinations(i)
              .exists(t => {
                t.flatten
                  .foldLeft(Array.fill(joltage.length)(0))((acc, cur) => {
                    acc.update(cur, acc(cur) + 1)
                    acc
                  })
                  .toList
                  .sameElements(joltage)
              })

            if a == true then 
              result = i
              boundary.break()
        }

        result
      })
      .sum
}
