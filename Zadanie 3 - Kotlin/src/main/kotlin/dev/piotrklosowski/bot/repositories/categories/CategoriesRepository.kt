package dev.piotrklosowski.bot.repositories.categories

import dev.piotrklosowski.bot.repositories.IRepository
import dev.piotrklosowski.bot.repositories.categories.models.Category

// CategoriesRepository ...
class CategoriesRepository(categories: Array<Category>): IRepository {
    private val categoriesList: MutableList<Category> = categories.toMutableList();

    fun getAll(): List<Category> {
        return categoriesList
    }
}