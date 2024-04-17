package dev.piotrklosowski.bot

import dev.piotrklosowski.bot.commands.ICommand
import io.ktor.server.application.*

interface ICommandParser<T> {
    fun parseInput(call: ApplicationCall, input: T): ICommand
}