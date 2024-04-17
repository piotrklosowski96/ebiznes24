package dev.piotrklosowski.bot.clients.discord.models

import kotlinx.serialization.Serializable

@Serializable
//InteractionResponseObject ...
data class InteractionResponseObject(
    val type: InteractionResponseType,
    val data: InteractionMessageCallbackData? = null,
)
