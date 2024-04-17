package dev.piotrklosowski.bot.clients.discord.models

import kotlinx.serialization.Serializable

@Serializable
data class ApplicationCommandData(
    val id: String,
    val name: String,
    val type: String,
    val options: Array<ApplicationCommandInteractionDataOption>? = null,
) {
    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as ApplicationCommandData

        if (id != other.id) return false
        if (name != other.name) return false
        if (type != other.type) return false
        if (options != null) {
            if (other.options == null) return false
            if (!options.contentEquals(other.options)) return false
        } else if (other.options != null) return false

        return true
    }

    override fun hashCode(): Int {
        var result = id.hashCode()
        result = 31 * result + name.hashCode()
        result = 31 * result + type.hashCode()
        result = 31 * result + (options?.contentHashCode() ?: 0)
        return result
    }
}
