package dev.piotrklosowski.bot.clients.discord

import dev.piotrklosowski.bot.clients.IClient
import dev.piotrklosowski.bot.clients.discord.models.*
import io.ktor.client.*
import io.ktor.client.plugins.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.request.*
import io.ktor.http.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.json.Json

// DiscordClient ...
class DiscordClient(token: String): IClient {
    private val httpClient = HttpClient() {
        install(ContentNegotiation) {
            json(Json {
                prettyPrint = true
                isLenient = true
            })
        }
        defaultRequest {
            header(HttpHeaders.Authorization, "Bot $token")
            header(HttpHeaders.ContentType, "application/json")
            header(HttpHeaders.UserAgent, "Discord Bot v1.0.0")
        }
    }
    private val discordAPIEndpoint = "https://discord.com/api/v10"

    override suspend fun sendTextMessage(messageContents: String, channelId: String) {
        httpClient.post("$discordAPIEndpoint/channels/${channelId}/messages") {
            contentType(ContentType.Application.Json)
            setBody(CreateMessage(messageContents))
        }
    }

    suspend fun handleApplicationCommandInteraction(interactionObject: InteractionObject) {
        when(interactionObject.data?.name) {
            "message_to_bot" -> handleMessageToBotCommand(interactionObject.id, interactionObject.token, interactionObject.data)
            else -> println(interactionObject.data?.name)
        }
    }

    private suspend fun handleMessageToBotCommand(interactionId: String, interactionToken: String, commandData: ApplicationCommandData) {
        val message = commandData.options?.first { o -> o.name == "message" }?.value

        httpClient.post("$discordAPIEndpoint/interactions/${interactionId}/${interactionToken}/callback") {
            contentType(ContentType.Application.Json)
            setBody(InteractionResponseObject(
                type = InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE,
                data = InteractionMessageCallbackData(
                    content = "Message sent to bot: `$message`",
                )
            ))
        }
    }
}