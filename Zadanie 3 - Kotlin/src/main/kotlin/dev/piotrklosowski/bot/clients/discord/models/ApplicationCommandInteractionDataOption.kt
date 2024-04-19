package dev.piotrklosowski.bot.clients.discord.models

import kotlinx.serialization.Serializable

@Serializable
data class ApplicationCommandInteractionDataOption(
    val name: String,
    val type: Int,
    val value: String? = null,
    val options: Array<ApplicationCommandInteractionDataOption>? = null,
) {
    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as ApplicationCommandInteractionDataOption

        if (name != other.name) return false
        if (type != other.type) return false
        if (value != other.value) return false
        if (options != null) {
            if (other.options == null) return false
            if (!options.contentEquals(other.options)) return false
        } else if (other.options != null) return false

        return true
    }

    override fun hashCode(): Int {
        var result = name.hashCode()
        result = 31 * result + type
        result = 31 * result + (value?.hashCode() ?: 0)
        result = 31 * result + (options?.contentHashCode() ?: 0)
        return result
    }
}
