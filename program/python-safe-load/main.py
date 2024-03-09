import yaml
import sys

path = sys.argv[1]
content = open(path, 'r')
obj = yaml.safe_load(content)
print(obj)
