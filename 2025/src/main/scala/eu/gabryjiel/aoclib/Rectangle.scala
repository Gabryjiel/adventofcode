package eu.gabryjiel.aoclib

case class Rectangle(a: Point, b: Point) {
  def area: Long = (math.abs(a.x - b.x) + 1) * (math.abs(a.y - b.y) + 1)
}
