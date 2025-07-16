import 'package:flutter_webrtc/flutter_webrtc.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class CallService {
  final int myUID;
  final int peerUID;
  final String token;
  RTCPeerConnection? _pc;
  final _localRenderer = RTCVideoRenderer();
  final _remoteRenderer = RTCVideoRenderer();
  late WebSocketChannel _ws;

  CallService(this.myUID, this.peerUID, this.token) {
    _localRenderer.initialize();
    _remoteRenderer.initialize();
    _ws = WebSocketChannel.connect(
        Uri.parse('ws://localhost:8080/ws/call?token=$token'));
    _ws.stream.listen(_onSignal);
    _createPC();
  }

  Future<void> _createPC() async {
    _pc = await createPeerConnection({
      'iceServers': [
        {'urls': 'stun:stun.l.google.com:19302'}
      ]
    });
    _pc!.onIceCandidate = (e) => _send('ice', e.toMap());
    _pc!.onTrack = (e) {
      if (e.track.kind == 'audio') {
        _remoteRenderer.srcObject = e.streams[0];
      }
    };
    _pc!.onAddTrack = (stream, track) => _remoteRenderer.srcObject = stream;
  }

  Future<void> call() async {
    final stream = await navigator.mediaDevices
        .getUserMedia({'audio': true, 'video': false});
    _localRenderer.srcObject = stream;
    stream.getTracks().forEach((track) => _pc!.addTrack(track, stream));

    final offer = await _pc!.createOffer();
    await _pc!.setLocalDescription(offer);
    _send('offer', offer.toMap());
  }

  Future<void> hangUp() async {
    _pc?.close();
    _localRenderer.srcObject?.getTracks().forEach((t) => t.stop());
    _remoteRenderer.srcObject?.getTracks().forEach((t) => t.stop());
  }

  void _onSignal(dynamic data) {
    final msg = jsonDecode(data);
    switch (msg['type']) {
      case 'offer':
        _onOffer(RTCSessionDescription(
            msg['data']['sdp'], msg['data']['type']));
        break;
      case 'answer':
        _pc?.setRemoteDescription(RTCSessionDescription(
            msg['data']['sdp'], msg['data']['type']));
        break;
      case 'ice':
        _pc?.addCandidate(RTCIceCandidate(
            msg['data']['candidate'],
            msg['data']['sdpMid'],
            msg['data']['sdpMLineIndex']));
        break;
    }
  }

  Future<void> _onOffer(RTCSessionDescription offer) async {
    await _pc!.setRemoteDescription(offer);
    final answer = await _pc!.createAnswer();
    await _pc!.setLocalDescription(answer);
    _send('answer', answer.toMap());
  }

  void _send(String type, Map<String, dynamic> data) {
    _ws.sink.add(jsonEncode(
        {'type': type, 'to': peerUID, 'data': data}));
  }
}