import sbtassembly.MergeStrategy

ThisBuild / scalaVersion := "2.13.13"

ThisBuild / version := "1.0"

ThisBuild / assemblyMergeStrategy in assembly := {
  case "application.prod.conf" => MergeStrategy.concat
  case PathList("META-INF", "versions", "9", "module-info.class") => MergeStrategy.discard
  case "module-info.class" => MergeStrategy.discard
  case "play/reference-overrides.conf" => MergeStrategy.concat
  case x =>
    val oldStrategy: String => MergeStrategy = (assemblyMergeStrategy in assembly).value
    oldStrategy(x)
}

lazy val root = (project in file("."))
  .enablePlugins(PlayScala)
  .settings(
    name := """zadanie2""",
    libraryDependencies ++= Seq(
      guice,
      "org.scalatestplus.play" %% "scalatestplus-play" % "7.0.1" % Test,
      "ch.qos.logback" % "logback-classic" % "1.5.3"
    )
  )