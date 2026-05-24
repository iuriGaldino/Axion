package com.axion.app

import androidx.compose.ui.window.Window
import androidx.compose.ui.window.application
import com.axion.app.di.appModule
import org.koin.core.context.startKoin

fun main() {
    startKoin {
        modules(appModule)
    }
    
    application {
        Window(onCloseRequest = ::exitApplication, title = "Axion") {
            App()
        }
    }
}
