package dev.piotrklosowski.bot.commands

import dev.piotrklosowski.bot.clients.IClient

// SendTextMessageCommand ...
class SendTextMessageCommand(
    private val receiver: IClient,
    private val content: String = "",
    private val channelId: String = "",
): ICommand {
    override suspend fun execute() {
        receiver.sendTextMessage(content, channelId)
    }
}