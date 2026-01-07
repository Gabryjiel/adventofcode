package eu.gabryjiel.aoclib

case class Interval(start: Long, end: Long) {
  def contains(value: Long): Boolean = value >= start && value <=end
  def size: Long = end - start + 1
}
