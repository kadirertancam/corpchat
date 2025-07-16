class CallScreen extends StatefulWidget {
  final int peerUID;
  const CallScreen(this.peerUID, {Key? key}) : super(key: key);
  @override
  _CallScreenState createState() => _CallScreenState();
}

class _CallScreenState extends State<CallScreen> {
  late CallService _call;

  @override
  void initState() {
    super.initState();
    _call = CallService(myUID, widget.peerUID, token);
  }

  @override
  void dispose() {
    _call.hangUp();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Sesli Arama")),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(Icons.phone, size: 64),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: () => _call.call(),
              child: Text("ARA"),
            ),
            ElevatedButton(
              onPressed: _call.hangUp,
              child: Text("KAPAT"),
            ),
          ],
        ),
      ),
    );
  }
}