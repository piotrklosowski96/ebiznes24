package dev.piotrklosowski.bot.clients.discord.models

import kotlinx.serialization.KSerializer
import kotlinx.serialization.SerialInfo
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable
import kotlinx.serialization.descriptors.PrimitiveKind
import kotlinx.serialization.descriptors.PrimitiveSerialDescriptor
import kotlinx.serialization.descriptors.SerialDescriptor
import kotlinx.serialization.encoding.Decoder
import kotlinx.serialization.encoding.Encoder

@Serializable(with = InteractionResponseType.InteractionResponseTypeSerializer::class)
enum class InteractionResponseType() {
    PONG,
    CHANNEL_MESSAGE_WITH_SOURCE,
    DEFERRED_CHANNEL_MESSAGE_WITH_SOURCE,
    DEFERRED_UPDATE_MESSAGE,
    UPDATE_MESSAGE,
    APPLICATION_COMMAND_AUTOCOMPLETE_RESULT,
    MODAL,
    PREMIUM_REQUIRED;

    object InteractionResponseTypeSerializer : KSerializer<InteractionResponseType> {
        override val descriptor: SerialDescriptor = PrimitiveSerialDescriptor("InteractionResponseType", PrimitiveKind.INT)

        override fun serialize(encoder: Encoder, value: InteractionResponseType) {
            when(value) {
                PONG -> encoder.encodeInt(1)
                CHANNEL_MESSAGE_WITH_SOURCE -> encoder.encodeInt(4)
                DEFERRED_CHANNEL_MESSAGE_WITH_SOURCE -> encoder.encodeInt(5)
                DEFERRED_UPDATE_MESSAGE -> encoder.encodeInt(6)
                UPDATE_MESSAGE -> encoder.encodeInt(7)
                APPLICATION_COMMAND_AUTOCOMPLETE_RESULT -> encoder.encodeInt(8)
                MODAL -> encoder.encodeInt(9)
                PREMIUM_REQUIRED -> encoder.encodeInt(10)
            }
        }

        override fun deserialize(decoder: Decoder): InteractionResponseType {
            return when(decoder.decodeInt()) {
                1 -> PONG
                4 -> CHANNEL_MESSAGE_WITH_SOURCE
                5 -> DEFERRED_CHANNEL_MESSAGE_WITH_SOURCE
                6 -> DEFERRED_UPDATE_MESSAGE
                7 -> UPDATE_MESSAGE
                8 -> APPLICATION_COMMAND_AUTOCOMPLETE_RESULT
                9 -> MODAL
                10 -> PREMIUM_REQUIRED
                else -> error("unknown interaction")
            }
        }
    }
}