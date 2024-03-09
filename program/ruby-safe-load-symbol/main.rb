require 'yaml'

path = ARGV[0]
text = File.read(path)
obj = YAML.load(text, permitted_classes: [Symbol])
pp(obj)
