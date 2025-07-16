import 'package:flutter/foundation.dart';
import '../models/message.dart';
import '../services/ws_service.dart';

class ChatProvider with ChangeNotifier {
  final WsService _ws;
  final int myUID;
  final int peerUID;
  final List<Message> _messages = [];
  List<Message> get messages => _messages;
  bool peerTyping = false;

  ChatProvider(this._ws, this.myUID, this.peerUID) {
    _ws.stream.listen((json) {
      if (json['type'] == 'typing') {
        peerTyping = json['typing'];
      } else {
        _messages.add(Message(json['from_uid'], json['to_uid'], json['body']));
      }
      notifyListeners();
    });
  }

  void send(String text) {
    _ws.sendMessage(peerUID, text);
    _messages.add(Message(myUID, peerUID, text));
    notifyListeners();
  }

  void setTyping(bool typing) => _ws.sendTyping(peerUID, typing);
}