package dev.piotrklosowski.bot.models

import kotlinx.serialization.Serializable

// SendTextMessage ...
@Serializable
data class SendTextMessage(
    val content: String,
    val channelId: String,
)
