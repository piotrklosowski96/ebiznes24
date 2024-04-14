package dev.piotrklosowski.bot.clients

import dev.piotrklosowski.bot.commands.ICommand

// IClient ...
interface IClient {
    suspend fun sendTextMessage(messageContents: String, channelId: String)
}