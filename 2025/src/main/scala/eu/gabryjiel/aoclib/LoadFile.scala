package eu.gabryjiel.aoclib

import scala.io.Source

object LoadFile {
  def loadLines(pathname: os.Path): Array[String] =
    Source.fromFile(pathname.toString()).getLines().toArray

  def loadResource(resourcePath: String): Array[String] = {
    val source = Source.fromResource(resourcePath)
    val lines = source.getLines().toArray
    source.close()
    lines
  }

  def loadResourceAsString(resourcePath: String): String = {
    val source = Source.fromResource(resourcePath)
    val lines = source.getLines().toArray.mkString("\n")
    source.close()
    lines
  }
}
