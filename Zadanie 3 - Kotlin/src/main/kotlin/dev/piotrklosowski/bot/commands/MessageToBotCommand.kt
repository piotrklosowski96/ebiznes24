package dev.piotrklosowski.bot.commands

import dev.piotrklosowski.bot.clients.discord.DiscordClient
import dev.piotrklosowski.bot.clients.discord.models.InteractionMessageCallbackData
import dev.piotrklosowski.bot.clients.discord.models.InteractionObject
import dev.piotrklosowski.bot.clients.discord.models.InteractionResponseObject
import dev.piotrklosowski.bot.clients.discord.models.InteractionResponseType
import org.koin.core.component.KoinComponent
import org.koin.core.component.inject

class MessageToBotCommand(private val interactionObject: InteractionObject) : ICommand, KoinComponent {
    private val receiver: DiscordClient by inject()

    override suspend fun execute() {
        val message = interactionObject.data?.options?.first { o -> o.name == "message" }?.value
        val interactionResponseObject = InteractionResponseObject(
            type = InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
            data = InteractionMessageCallbackData(
                content = "Message sent to bot: `$message`",
            )
        )
        receiver.respondToInteraction(interactionObject.id, interactionObject.token, interactionResponseObject)
    }
}