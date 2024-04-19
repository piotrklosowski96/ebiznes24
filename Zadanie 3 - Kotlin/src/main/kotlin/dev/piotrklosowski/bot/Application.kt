package dev.piotrklosowski.bot

import dev.piotrklosowski.bot.clients.discord.DiscordClient
import dev.piotrklosowski.bot.repositories.categories.CategoriesRepository
import dev.piotrklosowski.bot.repositories.categories.models.Category
import dev.piotrklosowski.bot.repositories.products.ProductsRepository
import dev.piotrklosowski.bot.repositories.products.models.Product
import dev.piotrklosowski.bot.routes.discordBotRouting
import io.ktor.serialization.kotlinx.json.*
import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.plugins.contentnegotiation.*
import io.ktor.server.plugins.doublereceive.*
import io.ktor.server.routing.*
import kotlinx.serialization.json.Json
import org.koin.dsl.module
import org.koin.ktor.plugin.Koin

fun main() {
    embeddedServer(Netty, port = 8080, host = "0.0.0.0", module = Application::main).start(wait = true)
}

fun Application.main() {
    val discordClient = DiscordClient(System.getenv("DISCORD_BOT_TOKEN"))

    install(Koin) {
        modules(module {
            single<CategoriesRepository> { CategoriesRepository(
                Array(5) {
                i -> Category("$i", "category_name_$i", "category_description_$i")
            }) }
            single<ProductsRepository> { ProductsRepository(arrayOf(
                Product("0", "product_name_0", "product_description_0", "category_name_0"),
                Product("1", "product_name_1", "product_description_1", "category_name_0"),
                Product("2", "product_name_2", "product_description_2", "category_name_1"),
                Product("3", "product_name_3", "product_description_3", "category_name_1"),
            )) }
            single<DiscordClient> { discordClient }
        })
    }

    install(DoubleReceive)

    install(ContentNegotiation) {
        json(Json {
            prettyPrint = true
            isLenient = true
            ignoreUnknownKeys = true
        })
    }

    install(Routing) {
        route("/api/v1") {
            route("/discord") {
                discordBotRouting(discordClient)
            }
        }
    }
}
