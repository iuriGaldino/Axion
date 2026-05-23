package com.axion.app.ui.screens.dashboard

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp

@Composable
fun DashboardScreen() {
    Scaffold(
        topBar = { TopAppBar(title = { Text("Dashboard Axion") }) }
    ) { padding ->
        LazyColumn(modifier = Modifier.padding(padding).fillMaxSize().padding(16.dp)) {
            item {
                Card(modifier = Modifier.fillMaxWidth()) {
                    Column(modifier = Modifier.padding(16.dp)) {
                        Text("Saldo Total", style = MaterialTheme.typography.labelMedium)
                        Text("R$ 12.500,00", style = MaterialTheme.typography.headlineLarge)
                    }
                }
            }
            item { Spacer(Modifier.height(16.dp)) }
            item {
                Text("Transações Recentes", style = MaterialTheme.typography.titleMedium)
            }
            // Lista de transações virá aqui
        }
    }
}
