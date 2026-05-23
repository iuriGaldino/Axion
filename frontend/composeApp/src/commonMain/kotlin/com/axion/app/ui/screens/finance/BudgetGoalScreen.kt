package com.axion.app.ui.screens.finance

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp

@Composable
fun BudgetGoalScreen() {
    Scaffold(
        topBar = { TopAppBar(title = { Text("Orçamentos e Metas") }) }
    ) { padding ->
        LazyColumn(modifier = Modifier.padding(padding).fillMaxSize().padding(16.dp)) {
            item {
                Text("Seus Orçamentos", style = MaterialTheme.typography.titleLarge)
                Spacer(Modifier.height(8.dp))
                Card(modifier = Modifier.fillMaxWidth()) {
                    Column(modifier = Modifier.padding(16.dp)) {
                        Text("Alimentação")
                        LinearProgressIndicator(progress = 0.7f, modifier = Modifier.fillMaxWidth())
                        Text("70% do orçamento utilizado")
                    }
                }
            }
            item { Spacer(Modifier.height(24.dp)) }
            item {
                Text("Suas Metas", style = MaterialTheme.typography.titleLarge)
                Spacer(Modifier.height(8.dp))
                Card(modifier = Modifier.fillMaxWidth()) {
                    Column(modifier = Modifier.padding(16.dp)) {
                        Text("Viagem para Europa")
                        LinearProgressIndicator(progress = 0.4f, modifier = Modifier.fillMaxWidth())
                        Text("R$ 4.000,00 / R$ 10.000,00")
                    }
                }
            }
        }
    }
}
