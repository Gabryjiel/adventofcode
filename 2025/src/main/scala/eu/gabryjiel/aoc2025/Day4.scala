package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile.loadResource
import scala.collection.mutable.HashMap
import scala.collection.mutable.HashSet

object Day4 {
  def main(args: Array[String]): Unit = {
    lazy val input = loadResource("eu/gabryjiel/aoc2025/day4.txt")
    println(s"Day4.Part1: ${this.Part1(input)}")
    println(s"Day4.Part2: ${this.Part2(input)}")
  }

  def Part1(lines: Array[String]): Int = {
    val n = lines.length
    val m = lines.head.length
    var count = 0

    for y <- 0 until n do {
      for x <- 0 until m if lines(y)(x) == '@' do {
        val surroundingCount = List(
          lines.lift(y - 1).getOrElse("").lift(x - 1).getOrElse('x'),
          lines.lift(y - 1).getOrElse("").lift(x).getOrElse('x'),
          lines.lift(y - 1).getOrElse("").lift(x + 1).getOrElse('x'),
          lines.lift(y).getOrElse("").lift(x - 1).getOrElse('x'),
          lines.lift(y).getOrElse("").lift(x + 1).getOrElse('x'),
          lines.lift(y + 1).getOrElse("").lift(x - 1).getOrElse('x'),
          lines.lift(y + 1).getOrElse("").lift(x).getOrElse('x'),
          lines.lift(y + 1).getOrElse("").lift(x + 1).getOrElse('x')
        ).count(_ == '@')

        if surroundingCount < 4 then {
          count = count + 1
        }
      }
    }

    count
  }

  def makeHashMap(lines: Array[String]): HashSet[(Int, Int)] = {
    val set = new HashSet[(Int, Int)]

    lines.zipWithIndex.foldLeft(set)((acc, line) =>
      acc ++ line(0).zipWithIndex.foldLeft(acc)((acc, char) =>
        if char(0) == '@' then acc + ((line(1), char(1))) else acc
      )
    )
  }

  def findToDelete(set: HashSet[(Int, Int)]): List[(Int, Int)] = {
    set.foldLeft(List[(Int, Int)]())((acc, cur) => {
      val surroundingCount = List(
        set.contains((cur(0) - 1, cur(1) - 1)),
        set.contains((cur(0) - 1, cur(1))),
        set.contains((cur(0) - 1, cur(1) + 1)),
        set.contains((cur(0), cur(1) - 1)),
        set.contains((cur(0), cur(1) + 1)),
        set.contains((cur(0) + 1, cur(1) - 1)),
        set.contains((cur(0) + 1, cur(1))),
        set.contains((cur(0) + 1, cur(1) + 1)),
      ).count(_ == true)

      if surroundingCount < 4 then {
        acc :+ cur
      } else {
        acc
      }
    })
  }

  def Part2(lines: Array[String]): Int = {
    val set = makeHashMap(lines)
    var flag = true
    var count = 0

    while flag do {
      val toDelete = findToDelete(set)

      if toDelete.length == 0 then
        flag = false
      else
        count += toDelete.length
        toDelete.foreach(element => set.remove(element))
    }

    count
  }
}
