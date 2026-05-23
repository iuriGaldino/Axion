package com.axion.app

import androidx.compose.runtime.*
import com.axion.app.ui.screens.auth.LoginScreen
import com.axion.app.ui.screens.auth.RegisterScreen
import com.axion.app.ui.screens.dashboard.DashboardScreen
import com.axion.app.ui.screens.profile.ProfileScreen
import com.axion.app.ui.screens.profile.SettingsScreen
import com.axion.app.ui.theme.AxionTheme

enum class Screen {
    Login, Register, Dashboard, Profile, Settings
}

@Composable
fun App() {
    var currentScreen by remember { mutableStateOf(Screen.Login) }

    AxionTheme {
        when (currentScreen) {
            Screen.Login -> LoginScreen(
                onLoginSuccess = { currentScreen = Screen.Dashboard },
                onNavigateToRegister = { currentScreen = Screen.Register }
            )
            Screen.Register -> RegisterScreen(
                onRegisterSuccess = { currentScreen = Screen.Login },
                onNavigateToLogin = { currentScreen = Screen.Login }
            )
            Screen.Dashboard -> DashboardScreen()
            Screen.Profile -> ProfileScreen("Usuário Axion", "user@axion.com")
            Screen.Settings -> SettingsScreen()
        }
    }
}
