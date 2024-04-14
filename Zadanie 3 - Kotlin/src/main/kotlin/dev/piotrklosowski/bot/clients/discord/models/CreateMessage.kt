package dev.piotrklosowski.bot.clients.discord.models

import kotlinx.serialization.Serializable

@Serializable
data class CreateMessage(
    val content: String = ""
)