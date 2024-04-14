package dev.piotrklosowski.bot.commands

// ICommand ...
interface ICommand {
    suspend fun execute()
}