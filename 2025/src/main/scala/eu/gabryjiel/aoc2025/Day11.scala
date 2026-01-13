package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile.loadResource
import eu.gabryjiel.aoclib.Point
import eu.gabryjiel.aoclib.Rectangle
import scala.util.boundary
import scala.collection.concurrent.TrieMap

object Day11 {
  def main(args: Array[String]): Unit = {
    lazy val input = loadResource("eu/gabryjiel/aoc2025/day11.txt")
    println(s"Day11.Part1: ${this.Part1(input)}")
    println(s"Day11.Part2: ${this.Part2(input)}")
  }

  def parse(line: String) = line match {
    case s"$in: $out" => (in, out.split(' '))
  }

  def createMapGraph(lines: Array[String]): Map[String, Set[String]] =
    lines
      .foldLeft(Map.empty[String, Set[String]])((acc, cur) => {
        val (in, out) = parse(cur)
        acc + (in -> out.toSet)
      })

  def memoize[A, B](f: A => B): A => B =
    val cache = TrieMap.empty[A, B]
    arg => cache.getOrElseUpdate(arg, f(arg))

  def dfs(
      graph: Map[String, Set[String]],
      current: String,
      finish: String,
      visited: Array[String]
  ): Int =
    if current == finish then 1
    else
      graph.get(current) match {
        case None       => 0
        case Some(next) => {
          val updatedVisited = visited :+ current

          next.foldLeft(0)((acc, cur) => {
            if updatedVisited.contains(cur) then 0
            else
              val value = dfs(graph, cur, finish, updatedVisited)
              acc + value
          })
        }
      }

  def memoizedDfs(
      graph: Map[String, Set[String]],
      current: String,
      finish: String,
      cache: TrieMap[(String, String), Int],
      visited: Array[String]
  ): Int =
    cache.get((current, finish)) match {
      case Some(value) =>
        println(s"CACHE: $current,$finish -> $value")
        value
      case None =>
        println(s"No-CACHE: $current,$finish")
        if current == finish then
          println(s"finish: $current,$finish -> ${visited.mkString(",")}")
          if visited.contains("dac") && visited.contains("fft") then 1
          else 0
        else
          graph.get(current) match
            case None       => 0
            case Some(next) =>
              val updatedVisited = visited :+ current

              next.foldLeft(0)((acc, cur) => {
                if updatedVisited.contains(cur) then 0
                else
                  val value =
                    memoizedDfs(graph, cur, finish, cache, updatedVisited)
                  cache.update((cur, finish), value)
                  acc + value
              })

    }

  def Part1(lines: Array[String]): Long =
    val graph = createMapGraph(lines)
    dfs(graph, "you", "out", Array.empty[String])

  def Part2(lines: Array[String]): Long =
    val graph = createMapGraph(lines)
    memoizedDfs(graph, "svr", "out", TrieMap.empty, Array.empty)
    // val f2d = dfs(graph, "fft", "dac", Array.empty[String])
    // val d2f = dfs(graph, "dac", "fft", Array.empty[String])
    // val (a, b, c) =
    //   if d2f == 0 then ("fft", f2d, "dac") else ("dac", d2f, "fft")
    // dfs(graph, "svr", a, Array.empty[String]) * b * dfs(
    //   graph,
    //   c,
    //   "out",
    //   Array.empty[String]
    // )
}
