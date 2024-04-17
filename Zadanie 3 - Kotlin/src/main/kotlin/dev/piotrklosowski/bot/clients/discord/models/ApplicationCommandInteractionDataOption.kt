package dev.piotrklosowski.bot.clients.discord.models

import kotlinx.serialization.Serializable

@Serializable
data class ApplicationCommandInteractionDataOption(
    val name: String,
    val type: Int,
    val value: String,
)
