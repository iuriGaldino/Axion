package com.axion.app.ui.screens.finance

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.axion.app.domain.model.Category

@Composable
fun CategoryScreen(categories: List<Category>) {
    Scaffold(
        topBar = { TopAppBar(title = { Text("Categorias") }) }
    ) { padding ->
        LazyColumn(modifier = Modifier.padding(padding).fillMaxSize()) {
            items(categories) { category ->
                ListItem(
                    headlineContent = { Text(category.name) },
                    supportingContent = { Text(if (category.isIncome) "Receita" else "Despesa") },
                    trailingContent = {
                        Box(modifier = Modifier.size(24.dp)) // Placeholder para cor/ícone
                    }
                )
                HorizontalDivider()
            }
        }
    }
}
