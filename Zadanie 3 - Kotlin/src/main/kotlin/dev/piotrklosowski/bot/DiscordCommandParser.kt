package dev.piotrklosowski.bot

import dev.piotrklosowski.bot.clients.discord.models.InteractionObject
import dev.piotrklosowski.bot.clients.discord.models.InteractionType
import dev.piotrklosowski.bot.commands.*
import io.ktor.server.application.*
import io.ktor.server.response.*

class DiscordCommandParser(): ICommandParser<InteractionObject> {
    override fun parseInput(call: ApplicationCall, input: InteractionObject): ICommand {
        return when (input.type) {
            InteractionType.PING -> PongResponseCommand(call)
            InteractionType.APPLICATION_COMMAND -> when(input.data?.name) {
                "message_to_bot" -> MessageToBotCommand(input)
                "categories" -> GetCategoriesCommand(input)
                else -> error("unknown command type $input")
            }
            else -> error("unknown command type $input")
        }
    }
}