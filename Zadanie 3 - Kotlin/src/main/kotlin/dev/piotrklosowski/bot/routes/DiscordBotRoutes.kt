package dev.piotrklosowski.bot.routes

import dev.piotrklosowski.bot.clients.discord.DiscordClient
import dev.piotrklosowski.bot.commands.SendTextMessageCommand
import dev.piotrklosowski.bot.models.SendTextMessage
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.routing.*

// Route.discordBotRouting ...
fun Route.discordBotRouting(discordClient: DiscordClient) {
    post("/send_message") {
        val textMessageParams = call.receive<SendTextMessage>()
        val command = SendTextMessageCommand(discordClient, textMessageParams.content, textMessageParams.channelId)
        command.execute()
    }
}
