import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'services/ws_service.dart';
import 'providers/chat_provider.dart';


void main() => runApp(
  ChangeNotifierProvider(
    create: (_) => ChatProvider(WsService("JWT_HERE"), 1, 2),
    child: const MyApp(),
  ),
);
class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'CorpChat',
      home: ChatScreen(),
    );
  }
}

class ChatScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final chat = context.watch<ChatProvider>();
    final ctrl = TextEditingController();
    return Scaffold(
      appBar: AppBar(title: const Text("CorpChat")),
      body: Column(
        children: [
          Expanded(
            child: ListView.builder(
              itemCount: chat.messages.length,
              itemBuilder: (_, i) {
                final m = chat.messages[i];
                return Align(
                  alignment: m.fromUID == chat.myUID
                      ? Alignment.centerRight
                      : Alignment.centerLeft,
                  child: Container(
                    margin: const EdgeInsets.all(4),
                    padding: const EdgeInsets.all(8),
                    color: m.fromUID == chat.myUID ? Colors.blue : Colors.grey,
                    child: Text(m.body),
                  ),
                );
              },
            ),
          ),
          if (chat.peerTyping) const Text("Bob yazıyor…"),
          Row(
            children: [
              Expanded(
                child: TextField(
                  controller: ctrl,
                  onChanged: (_) => chat.setTyping(true),
                  onSubmitted: (text) {
                    chat.send(text);
                    ctrl.clear();
                    chat.setTyping(false);
                  },
                ),
              ),
              IconButton(
                icon: const Icon(Icons.send),
                onPressed: () {
                  chat.send(ctrl.text);
                  ctrl.clear();
                  chat.setTyping(false);
                },
              ),
            ],
          ),
        ],
      ),
    );
  }
}