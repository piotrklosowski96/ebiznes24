plugins {
    application
}

repositories {
    mavenCentral()
}

dependencies {
    testImplementation(libs.junit.jupiter)
    testRuntimeOnly("org.junit.platform:junit-platform-launcher")
	implementation("org.xerial:sqlite-jdbc:3.45.2.0")
}

application {
    mainClass = "org.example.App"
}

java {
    toolchain {
        languageVersion = JavaLanguageVersion.of(8)
    }
}

tasks.withType<Jar> {
    manifest {
        attributes["Main-Class"] = "org.example.App"
    }
}

tasks.named<Test>("test") {
    useJUnitPlatform()
}
