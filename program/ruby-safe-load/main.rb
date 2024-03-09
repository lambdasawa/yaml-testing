require 'yaml'

path = ARGV[0]
text = File.read(path)
obj = YAML.safe_load(text)
pp(obj)
