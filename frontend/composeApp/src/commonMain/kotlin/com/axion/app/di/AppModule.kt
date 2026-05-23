package com.axion.app.di

import com.axion.app.data.repository.FinanceRepositoryImpl
import com.axion.app.domain.repository.FinanceRepository
import com.axion.app.ui.screens.dashboard.DashboardViewModel
import io.ktor.client.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.serialization.kotlinx.json.*
import kotlinx.serialization.json.Json
import org.koin.dsl.module

val appModule = module {
    single {
        HttpClient {
            install(ContentNegotiation) {
                json(Json {
                    ignoreUnknownKeys = true
                    prettyPrint = true
                })
            }
        }
    }
    single<FinanceRepository> { FinanceRepositoryImpl(get()) }
    factory { DashboardViewModel(get()) }
}
