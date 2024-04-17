package dev.piotrklosowski.bot.commands

import dev.piotrklosowski.bot.clients.discord.DiscordClient
import dev.piotrklosowski.bot.clients.discord.models.InteractionMessageCallbackData
import dev.piotrklosowski.bot.clients.discord.models.InteractionObject
import dev.piotrklosowski.bot.clients.discord.models.InteractionResponseObject
import dev.piotrklosowski.bot.clients.discord.models.InteractionResponseType
import dev.piotrklosowski.bot.repositories.categories.CategoriesRepository
import org.koin.core.component.KoinComponent
import org.koin.core.component.inject

// GetCategoriesCommand ...
class GetCategoriesCommand(private val interactionObject: InteractionObject) : ICommand, KoinComponent {
    private val receiver: DiscordClient by inject()
    private val categoriesRepository: CategoriesRepository by inject()

    override suspend fun execute() {
        val categories = categoriesRepository.getAll()

        val interactionResponseObject = InteractionResponseObject(
            type = InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
            data = InteractionMessageCallbackData(
                content = categories.toString(),
            )
        )
        receiver.respondToInteraction(interactionObject.id, interactionObject.token, interactionResponseObject)
    }
}