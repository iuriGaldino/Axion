package com.axion.app.ui.screens.dashboard

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import org.koin.compose.koinInject

@Composable
fun DashboardScreen(viewModel: DashboardViewModel = koinInject()) {
    val uiState by viewModel.uiState.collectAsState()

    LaunchedEffect(Unit) {
        viewModel.loadData()
    }

    Scaffold(
        topBar = { TopAppBar(title = { Text("Dashboard Axion") }) }
    ) { padding ->
        if (uiState.isLoading) {
            Box(modifier = Modifier.fillMaxSize(), contentAlignment = androidx.compose.ui.Alignment.Center) {
                CircularProgressIndicator()
            }
        } else {
            LazyColumn(modifier = Modifier.padding(padding).fillMaxSize().padding(16.dp)) {
                item {
                    Card(modifier = Modifier.fillMaxWidth()) {
                        Column(modifier = Modifier.padding(16.dp)) {
                            Text("Saldo Total", style = MaterialTheme.typography.labelMedium)
                            Text("R$ ${uiState.balance}", style = MaterialTheme.typography.headlineLarge)
                            Spacer(Modifier.height(16.dp))
                            com.axion.app.ui.components.SimpleDonutChart(
                                data = listOf(60f, 30f, 10f),
                                colors = listOf(MaterialTheme.colorScheme.primary, MaterialTheme.colorScheme.secondary, MaterialTheme.colorScheme.tertiary),
                                modifier = Modifier.align(androidx.compose.ui.Alignment.CenterHorizontally)
                            )
                        }
                    }
                }
                item { Spacer(Modifier.height(16.dp)) }
                item {
                    Text("Transações Recentes", style = MaterialTheme.typography.titleMedium)
                }
            }
        }
    }
}
