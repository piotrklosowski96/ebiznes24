package dev.piotrklosowski.bot.clients.discord.models

import kotlinx.serialization.Serializable

@Serializable
data class InteractionObject(
    val id: String,
    val token: String,
    val type: InteractionType,
    val data: ApplicationCommandData? = null
)