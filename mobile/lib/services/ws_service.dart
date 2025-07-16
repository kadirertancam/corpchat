import 'dart:convert';
import 'dart:async';
import 'package:web_socket_channel/web_socket_channel.dart';

class WsService {
  final String token;
  late WebSocketChannel _channel;
  final _controller = StreamController<Map<String, dynamic>>.broadcast();
  Stream<Map<String, dynamic>> get stream => _controller.stream;

  WsService(this.token) {
    _channel = WebSocketChannel.connect(
      Uri.parse('ws://localhost:8080/ws?token=$token'),
    );
    _channel.stream.listen(
      (msg) => _controller.add(jsonDecode(msg)),
      onError: (e) => print('WS error: $e'),
    );
  }

  void sendMessage(int toUID, String body) {
    _channel.sink.add(jsonEncode({
      'to_uid': toUID,
      'body': body,
    }));
  }

  void sendTyping(int toUID, bool typing) {
    _channel.sink.add(jsonEncode({
      'to_uid': toUID,
      'type': 'typing',
      'typing': typing,
    }));
  }

  void close() {
    _channel.sink.close();
    _controller.close();
  }
}