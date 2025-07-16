void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: ChatScreen(),
    );
  }
}

class ChatScreen extends StatefulWidget {
  @override
  _ChatScreenState createState() => _ChatScreenState();
}

class _ChatScreenState extends State<ChatScreen> {
  final TextEditingController _ctrl = TextEditingController();
  List<String> msgs = [];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text("CorpChat MVP")),
      body: Column(
        children: [
          Expanded(
            child: ListView.builder(
              itemCount: msgs.length,
              itemBuilder: (_, i) => ListTile(title: Text(msgs[i])),
            ),
          ),
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: Row(
              children: [
                Expanded(child: TextField(controller: _ctrl)),
                IconButton(
                  icon: const Icon(Icons.send),
                  onPressed: () {
                    setState(() => msgs.add(_ctrl.text));
                    _ctrl.clear();
                  },
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}