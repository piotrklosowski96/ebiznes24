package dev.piotrklosowski.bot.commands

import dev.piotrklosowski.bot.clients.discord.models.InteractionResponseObject
import dev.piotrklosowski.bot.clients.discord.models.InteractionResponseType
import io.ktor.server.application.*
import io.ktor.server.response.*

// PongResponseCommand ...
class PongResponseCommand(private val call: ApplicationCall): ICommand {
    override suspend fun execute() {
        val responseObject = InteractionResponseObject(type = InteractionResponseType.PONG)

        call.respond(responseObject)
    }
}