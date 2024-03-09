import yaml
from yaml import Loader
import sys

path = sys.argv[1]
content = open(path, 'r')
obj = yaml.load(content, Loader=Loader)
print(obj)
