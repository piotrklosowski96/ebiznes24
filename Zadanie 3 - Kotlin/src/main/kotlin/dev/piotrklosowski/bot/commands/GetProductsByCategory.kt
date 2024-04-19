package dev.piotrklosowski.bot.commands

import dev.piotrklosowski.bot.clients.discord.DiscordClient
import dev.piotrklosowski.bot.clients.discord.models.InteractionMessageCallbackData
import dev.piotrklosowski.bot.clients.discord.models.InteractionObject
import dev.piotrklosowski.bot.clients.discord.models.InteractionResponseObject
import dev.piotrklosowski.bot.clients.discord.models.InteractionResponseType
import dev.piotrklosowski.bot.repositories.products.ProductsRepository
import org.koin.core.component.KoinComponent
import org.koin.core.component.inject

// GetProductsByCategory ...
class GetProductsByCategoryCommand(private val interactionObject: InteractionObject) : ICommand, KoinComponent {
    private val receiver: DiscordClient by inject()
    private val productsRepository: ProductsRepository by inject()

    override suspend fun execute() {
        val categoryName = interactionObject.data
            ?.options?.filter { p -> p.name == "get-by-category" }?.getOrNull(0)
            ?.options?.first { p -> p.name == "category-name" }?.value

        val products = productsRepository.getByCategoryName(categoryName)
        val interactionResponseObject = InteractionResponseObject(
            type = InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
            data = InteractionMessageCallbackData(
                content = products.toString(),
            )
        )
        receiver.respondToInteraction(interactionObject.id, interactionObject.token, interactionResponseObject)
    }
}