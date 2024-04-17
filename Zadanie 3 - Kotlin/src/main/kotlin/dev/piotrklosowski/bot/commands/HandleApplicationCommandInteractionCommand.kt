package dev.piotrklosowski.bot.commands

import dev.piotrklosowski.bot.clients.discord.DiscordClient
import dev.piotrklosowski.bot.clients.discord.models.InteractionObject

// HandleApplicationCommandInteractionCommand ...
class HandleApplicationCommandInteractionCommand(
    private val receiver: DiscordClient,
    private val interactionObject: InteractionObject
): ICommand {
    override suspend fun execute() {
        receiver.handleApplicationCommandInteraction(interactionObject)
    }
}
