package com.axion.app.data.repository

import com.axion.app.domain.model.Category
import com.axion.app.domain.model.Transaction
import com.axion.app.domain.repository.FinanceRepository
import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.request.*

class FinanceRepositoryImpl(private val client: HttpClient) : FinanceRepository {
    private val baseUrl = "http://localhost:8080/api/v1"

    override suspend fun getTransactions(): List<Transaction> {
        return client.get("$baseUrl/transactions").body()
    }

    override suspend fun getCategories(): List<Category> {
        return client.get("$baseUrl/categories").body()
    }
}
