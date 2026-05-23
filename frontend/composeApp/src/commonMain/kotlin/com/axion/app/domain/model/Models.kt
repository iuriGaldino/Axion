package com.axion.app.domain.model

import kotlinx.serialization.Serializable

@Serializable
data class User(
    val id: String,
    val name: String,
    val email: String,
    val avatar: String? = null
)

@Serializable
data class Transaction(
    val id: String,
    val amount: Double,
    val description: String,
    val date: String,
    val type: String,
    val categoryId: String? = null
)

@Serializable
data class Category(
    val id: String,
    val name: String,
    val icon: String? = null,
    val color: String? = null,
    val isIncome: Boolean
)
