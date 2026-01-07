package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile.loadResource
import eu.gabryjiel.aoclib.Interval
import scala.collection.mutable.HashMap
import scala.collection.mutable.HashSet

object Day5 {
  def main(args: Array[String]): Unit = {
    lazy val input = loadResource("eu/gabryjiel/aoc2025/day5.txt")
    println(s"Day5.Part1: ${this.Part1(input)}")
    println(s"Day5.Part2: ${this.Part2(input)}")
  }

  def Part1(lines: Array[String]): Long = {
    val emptyLineIndex = lines.indexOf("")
    val (rangesStr, ids) = lines.splitAt(emptyLineIndex)

    val intervals = rangesStr.map(_ match {
      case s"$l-$r" => Interval(l.toLong, r.toLong)
    })

    ids.view
      .drop(1)
      .count(idStr => intervals.exists(_.contains(idStr.toLong)))
  }

  def parseInterval(line: String): Interval = line match {
    case s"$l-$r" => Interval(l.toLong, r.toLong)
  }

  def filterHelper(p: Interval, cur: Interval): Boolean =
    p.contains(cur.start) ||
      p.contains(cur.end) ||
      (cur.start < p.start && cur.end > p.end)

  def Part2(lines: Array[String]): Long = {
    lines.view
      .dropRight(lines.length - lines.indexOf(""))
      .map(parseInterval)
      .sortBy(f => f.start)
      .foldLeft(Array.empty[Interval])((acc, cur) => {
        acc.lastOption match {
          case Some(interval) if interval.end + 1 >= cur.start =>
            acc.update(
              acc.length - 1,
              Interval(interval.start, interval.end.max(cur.end))
            )
            acc
          case _ => acc :+ cur
        }
      })
      .map(_.size)
      .sum
  }
}
