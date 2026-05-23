package com.axion.app.ui.screens.dashboard

import com.axion.app.domain.repository.FinanceRepository
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers

class DashboardViewModel(private val repository: FinanceRepository) {
    private val _uiState = MutableStateFlow(DashboardUiState())
    val uiState: StateFlow<DashboardUiState> = _uiState.asStateFlow()

    private val scope = CoroutineScope(Dispatchers.Main)

    fun loadData() {
        scope.launch {
            _uiState.value = _uiState.value.copy(isLoading = true)
            try {
                val transactions = repository.getTransactions()
                val balance = transactions.sumOf { if (it.type == "income") it.amount else -it.amount }
                _uiState.value = _uiState.value.copy(balance = balance, isLoading = false)
            } catch (e: Exception) {
                _uiState.value = _uiState.value.copy(isLoading = false)
            }
        }
    }
}

data class DashboardUiState(
    val balance: Double = 0.0,
    val isLoading: Boolean = false
)
