import urllib
import urllib2

url = 'http://10.14.26.102:8080/view/c'
# values = {'name' : 'Michael Foord',
#           'location' : 'Northampton',
#           'language' : 'Python' }

# data = urllib.urlencode(values)
req = urllib2.Request(url)
response = urllib2.urlopen(req)
print response.read()