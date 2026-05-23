package com.axion.app.data.remote.dto

import kotlinx.serialization.Serializable

@Serializable
data class AuthResponse(
    val access_token: String,
    val refresh_token: String
)

@Serializable
data class LoginRequest(
    val email: String,
    val password: String
)

@Serializable
data class RegisterRequest(
    val name: String,
    val email: String,
    val password: String
)
