package dev.piotrklosowski.bot.repositories.products

import dev.piotrklosowski.bot.repositories.IRepository
import dev.piotrklosowski.bot.repositories.products.models.Product

// ProductsRepository ...
class ProductsRepository(products: Array<Product>): IRepository {
    private val productsList: MutableList<Product> = products.toMutableList()

    fun getByCategoryName(category: String?): List<Product> {
        return productsList.filter { p -> p.category == category }
    }
}