package com.axion.app.ui.screens.profile

import androidx.compose.foundation.layout.*
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Person
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp

@Composable
fun ProfileScreen(name: String, email: String) {
    Column(
        modifier = Modifier.fillMaxSize().padding(16.dp),
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        Icon(
            Icons.Default.Person,
            contentDescription = null,
            modifier = Modifier.size(100.dp)
        )
        Spacer(Modifier.height(16.dp))
        Text(name, style = MaterialTheme.typography.headlineMedium)
        Text(email, style = MaterialTheme.typography.bodyLarge)
        Spacer(Modifier.height(32.dp))
        Button(onClick = {}, modifier = Modifier.fillMaxWidth()) {
            Text("Editar Perfil")
        }
        Spacer(Modifier.height(8.dp))
        OutlinedButton(onClick = {}, modifier = Modifier.fillMaxWidth()) {
            Text("Sair")
        }
    }
}
