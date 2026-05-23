package com.axion.app.domain.repository

import com.axion.app.domain.model.Transaction
import com.axion.app.domain.model.Category

interface FinanceRepository {
    suspend fun getTransactions(): List<Transaction>
    suspend fun getCategories(): List<Category>
}
