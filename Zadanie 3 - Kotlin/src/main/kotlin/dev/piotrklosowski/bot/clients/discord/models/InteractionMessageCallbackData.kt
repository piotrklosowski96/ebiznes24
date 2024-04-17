package dev.piotrklosowski.bot.clients.discord.models

import kotlinx.serialization.Serializable

@Serializable
// InteractionMessageCallbackData ...
data class InteractionMessageCallbackData(
    val content: String
)
