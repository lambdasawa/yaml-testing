using YamlDotNet.Serialization;

var deserializer = new DeserializerBuilder().Build();
var serializer = new SerializerBuilder().Build();

var path = args[0];
var content = File.ReadAllText(path);

var obj = deserializer.Deserialize<Dictionary<object, object>>(content);
Console.WriteLine(serializer.Serialize(obj));