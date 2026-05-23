package com.axion.app.ui.screens.dashboard

import com.axion.app.domain.repository.FinanceRepository
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers

class DashboardViewModel(private val repository: FinanceRepository) {
    private val _uiState = MutableStateFlow(DashboardUiState())
    val uiState: StateFlow<DashboardUiState> = _uiState.asStateFlow()

    fun loadData() {
        // Implementation for MVVM pattern
    }
}

data class DashboardUiState(
    val balance: Double = 0.0,
    val isLoading: Boolean = false
)
