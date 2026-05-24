package com.axion.app.ui.screens.finance

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.axion.app.domain.model.Transaction

@Composable
fun TransactionListScreen(transactions: List<Transaction>) {
    LazyColumn(modifier = Modifier.fillMaxSize()) {
        items(transactions) { transaction ->
            ListItem(
                headlineContent = { Text(transaction.description) },
                supportingContent = { Text(transaction.date) },
                trailingContent = { 
                    Text(
                        "R$ ${transaction.amount}",
                        color = if (transaction.type == "income") MaterialTheme.colorScheme.primary else MaterialTheme.colorScheme.error
                    ) 
                }
            )
            HorizontalDivider()
        }
    }
}
