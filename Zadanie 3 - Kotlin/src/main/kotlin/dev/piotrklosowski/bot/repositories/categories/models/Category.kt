package dev.piotrklosowski.bot.repositories.categories.models

import kotlinx.serialization.Serializable

@Serializable
data class Category(
    val id: String,
    val name: String,
    val description: String,
)
