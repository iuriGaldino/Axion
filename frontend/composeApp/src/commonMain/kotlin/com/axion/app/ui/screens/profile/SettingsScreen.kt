package com.axion.app.ui.screens.profile

import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp

@Composable
fun SettingsScreen() {
    var darkTheme by remember { mutableStateOf(false) }
    var language by remember { mutableStateOf("Português") }

    Column(modifier = Modifier.fillMaxSize().padding(16.dp)) {
        Text("Configurações", style = MaterialTheme.typography.headlineMedium)
        Spacer(Modifier.height(24.dp))
        
        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.SpaceBetween
        ) {
            Text("Tema Escuro")
            Switch(checked = darkTheme, onCheckedChange = { darkTheme = it })
        }
        
        Spacer(Modifier.height(16.dp))
        
        Text("Idioma", style = MaterialTheme.typography.titleMedium)
        Text(language, style = MaterialTheme.typography.bodyMedium)
        
        Spacer(Modifier.height(32.dp))
        
        Button(onClick = {}, modifier = Modifier.fillMaxWidth()) {
            Text("Salvar Alterações")
        }
    }
}
