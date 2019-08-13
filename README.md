# urlcodec

## URL Encoder/Decoder CLI

URL Decode an string

    $ urlcodec https%3A%2F%2Fwww.google.com%2Fsearch%3Fsource%3Dhp%26ei%3D9TRTXfHRIYOZjLsPkZ2EyAw%26q%3Durl%2Bdecode%26oq%3Durl%2Bdecode

    https://www.google.com/search?source=hp&ei=9TRTXfHRIYOZjLsPkZ2EyAw&q=url+decode&oq=url+decode

URL Encode an string 

    $ urlcodec -e https://www.google.com/search?source=hp&ei=9TRTXfHRIYOZjLsPkZ2EyAw&q=url+decode&oq=url+decode

    https%3A%2F%2Fwww.google.com%2Fsearch%3Fsource%3Dhp%26ei%3D9TRTXfHRIYOZjLsPkZ2EyAw%26q%3Durl%2Bdecode%26oq%3Durl%2Bdecode