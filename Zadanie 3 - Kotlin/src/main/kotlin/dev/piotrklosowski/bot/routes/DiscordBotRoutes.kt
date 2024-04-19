package dev.piotrklosowski.bot.routes

import com.iwebpp.crypto.TweetNaclFast
import dev.piotrklosowski.bot.clients.discord.DiscordCommandParser
import dev.piotrklosowski.bot.clients.discord.DiscordClient
import dev.piotrklosowski.bot.clients.discord.models.InteractionObject
import dev.piotrklosowski.bot.commands.SendTextMessageCommand
import dev.piotrklosowski.bot.models.SendTextMessage
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*

fun String.toHexByteArray(): ByteArray {
    check(length % 2 == 0) { "Must have an even length" }

    return chunked(2)
        .map { it.toInt(16).toByte() }
        .toByteArray()
}

// Route.discordBotRouting ...
fun Route.discordBotRouting(discordClient: DiscordClient) {
    val publicKey = System.getenv("DISCORD_BOT_PUBLIC_KEY")
    val signatureVerifier = TweetNaclFast.Signature(publicKey.toHexByteArray(), null)
    val discordCommandParser = DiscordCommandParser()

    fun verifyDiscordRequest(body: String, timestamp: String?, signature: String?): Boolean {
        if (timestamp == null || signature == null) {
            return false
        }

        val message = "$timestamp$body".toByteArray()
        val requestSignature = signature.toHexByteArray()

        return signatureVerifier.detached_verify(message, requestSignature)
    }

    post("/send_message") {
        val textMessageParams = call.receive<SendTextMessage>()
        val command = SendTextMessageCommand(discordClient, textMessageParams.content, textMessageParams.channelId)
        command.execute()
    }

    post("/interactions") {
        val rawBody = call.receive<String>()
        val requestTimestamp = call.request.headers["X-Signature-Timestamp"]
        val requestSignature = call.request.headers["X-Signature-Ed25519"]
        if (!verifyDiscordRequest(rawBody, requestTimestamp, requestSignature)) {
            call.respond(HttpStatusCode.Unauthorized)
            return@post
        }

        val interactionObject = call.receive<InteractionObject>()
        discordCommandParser.parseInput(call, interactionObject).execute()
    }
}
