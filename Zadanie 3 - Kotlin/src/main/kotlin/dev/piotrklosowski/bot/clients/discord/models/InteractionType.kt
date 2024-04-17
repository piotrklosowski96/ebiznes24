package dev.piotrklosowski.bot.clients.discord.models

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
enum class InteractionType(i: Int) {
    @SerialName(value =  "1")
    PING(1),

    @SerialName(value =  "2")
    APPLICATION_COMMAND(2),

    @SerialName(value = "3")
    MESSAGE_COMPONENT(3),

    @SerialName(value = "4")
    APPLICATION_COMMAND_AUTOCOMPLETE(4),

    @SerialName(value = "5")
    MODAL_SUBMIT(5),
}