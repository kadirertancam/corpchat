Future<List<Channel>> fetchChannels() async =>
    (await dio.get('/channels')).data.map<Channel>((c) => Channel(c['id'], c['name'])).toList();

Future<String> uploadFile(File file) async {
  final form = FormData();
  form.files.add(MapEntry('file', await MultipartFile.fromFile(file.path)));
  final res = await dio.post('/upload', data: form);
  return res.data['url'];
}