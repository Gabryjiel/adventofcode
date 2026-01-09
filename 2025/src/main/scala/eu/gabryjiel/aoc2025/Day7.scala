package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile.loadResource
import eu.gabryjiel.aoclib.Interval
import scala.collection.mutable.Map

object Day7 {
  def main(args: Array[String]): Unit = {
    lazy val input = loadResource("eu/gabryjiel/aoc2025/day7.txt")
    println(s"Day7.Part1: ${this.Part1(input)}")
    println(s"Day7.Part2: ${this.Part2(input)}")
  }

  def Part1(lines: Array[String]): Int =
    val grid = lines.map(line => line.toCharArray())
    val startIndex = grid.head.indexOf('S')

    var answer = 0
    var queue = List((1, startIndex))

    while !queue.isEmpty do
      queue = queue
        .flatMap(f =>
          val (y, x) = (f(0) + 1, f(1))

          if y >= grid.length then List.empty
          else if grid(y)(x) == '^' then
            answer = answer + 1
            List((y, x - 1), (y, x + 1))
          else List((y, x))
        )
        .distinct

    answer

  def Part2(lines: Array[String]): Long =
    val grid = lines.map(line => line.toCharArray())
    val startIndex = grid.head.indexOf('S')

    // Method from https://www.reddit.com/r/adventofcode/comments/1pg9w66/comment/nspy4et/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
    // val beams = scala.collection.mutable.ArraySeq.fill(grid(0).length)(0L)
    //
    // for y <- 0 until grid.length do
    //   for x <- 0 until beams.length do
    //     grid(y)(x) match {
    //       case 'S' => beams.update(x, 1)
    //       case '^' if beams(x) != 0 => {
    //         beams.update(x - 1, beams(x - 1) + beams(x))
    //         beams.update(x + 1, beams(x + 1) + beams(x))
    //         beams.update(x, 0)
    //       }
    //       case _ => None
    //     }
    //
    // println(beams.mkString(","))
    // beams.sum

    val rowsRange = 2.until(grid.length)
    val startBeamMap = Map(((1, startIndex), 1L))

    val a = rowsRange.foldLeft(startBeamMap)((prevRowBeamMap, cur) => {
      val currentGridRow = grid.view(cur)
      val prevBeams = prevRowBeamMap.view.filter(f => f(0)(0) == cur - 1)

      val nextBeamMap =
        prevBeams.foldLeft(Map.empty[(Int, Int), Long])(
          (currentRowBeamMap, prevBeam) => {
            val (y, x) = prevBeam(0)
            val potentialNewBeams =
              if currentGridRow(x) == '^' then
                List((y + 1, x - 1), (y + 1, x + 1))
              else List((y + 1, x))

            potentialNewBeams.foreach(newBeam => {
              currentRowBeamMap.get(newBeam) match {
                case None => currentRowBeamMap.update(newBeam, prevBeam(1))
                case Some(value) =>
                  currentRowBeamMap.update(newBeam, value + prevBeam(1))
              }
            })

            currentRowBeamMap
          }
        )

      nextBeamMap
    })

    a.toList.foldLeft(0L)((acc, cur) => acc + cur(1))
}
