`GET
/json-
import.php?
c=0&load%5B%5D=jquery
-core,jquery-
migrate&load%5B%5D=utils&ver=3.8
.2&json={%22firstName%22:%22Иван
%22,%22lastName%22:%22Иванов%22,
%22address%22:{%22postalCode%22:
101101},%22phoneNumbers%22:[%228
12123-1234%22,%22916123-
4567%22]}
`

1. urldecode `[]byte->[]byte`
https://www.urldecoder.io/golang/

c=0
load[]=jquery-core,jquery-migrate
load[]=utils
ver=3.8.2
json=`{"firstName":"Иван","lastName":"Иванов","address":{"postalCode":101101},"phoneNumbers":["812123-1234","916123-4567"]`

2. json
json.Unmarshal, separate for obj/arr

3. simple tree algo
