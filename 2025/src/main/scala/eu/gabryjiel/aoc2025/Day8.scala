package eu.gabryjiel.aoc2025

import eu.gabryjiel.aoclib.LoadFile.loadResource
import eu.gabryjiel.aoclib.Interval
import scala.collection.mutable.Map
import scala.collection.mutable.PriorityQueue

case class Position(x: Int, y: Int, z: Int)
case class CalculatedDistance(start: Position, end: Position, distance: Double)

object Day8 {
  def main(args: Array[String]): Unit = {
    lazy val input = loadResource("eu/gabryjiel/aoc2025/day8.txt")
    println(s"Day8.Part1: ${this.Part1(input, 1000)}")
    println(s"Day8.Part2: ${this.Part2(input)}")
  }

  def parseCoords(line: String): Position = line match {
    case s"$s1,$s2,$s3" => Position(s1.toInt, s2.toInt, s3.toInt)
  }

  def calculatedDistance(a: Position, b: Position): Double =
    math.sqrt(
      math.pow(a.x - b.x, 2.0) + math.pow(a.y - b.y, 2.0) + math.pow(
        a.z - b.z,
        2.0
      )
    )

  def Part1(lines: Array[String], noOfConnections: Int): Long =
    val coords = lines.iterator.map(parseCoords).toArray
    val pq = PriorityQueue.empty[CalculatedDistance](
      Ordering.by[CalculatedDistance, Double](_.distance).reverse
    )

    for i <- 0.until(coords.length) do
      for j <- (i + 1).until(coords.length) do
        pq.enqueue(
          CalculatedDistance(
            coords(i),
            coords(j),
            calculatedDistance(coords(i), coords(j))
          )
        )

    val circuits = Map.from(coords.zipWithIndex.map(f => (f._1, f._2)).toMap)
    pq.dequeueAll.take(noOfConnections).foreach(f => {
      (circuits.get(f.start), circuits.get(f.end)) match {
        case (Some(i), Some(j)) => {
          circuits.filter(t => t(1) == j).foreach(t => circuits.update(t(0), i))
        }
        case _ => {
          println("FAIL")
        }
      }
    })

    circuits.values
      .groupBy(identity)
      .mapValues(_.size)
      .values
      .toArray
      .sorted
      .takeRight(3)
      .product

  def Part2(lines: Array[String]): Long =
    val coords = lines.iterator.map(parseCoords).toArray
    val pq = PriorityQueue.empty[CalculatedDistance](
      Ordering.by[CalculatedDistance, Double](_.distance).reverse
    )

    for i <- 0.until(coords.length) do
      for j <- (i + 1).until(coords.length) do
        pq.enqueue(
          CalculatedDistance(
            coords(i),
            coords(j),
            calculatedDistance(coords(i), coords(j))
          )
        )

    val circuits = Map.from(coords.zipWithIndex.map(f => (f._1, f._2)).toMap)
    var prod = 0L
    while circuits.valuesIterator.distinct.length != 1 do
      val cd = pq.dequeue() 
      prod = cd.start.x.toLong * cd.end.x.toLong

      (circuits.get(cd.start), circuits.get(cd.end)) match {
        case (Some(i), Some(j)) => {
          circuits
            .filter(t => t(1) == j)
            .foreach(t => circuits.update(t(0), i))
        }
        case _ => {
          println("FAIL")
        }
    }

    prod
}
